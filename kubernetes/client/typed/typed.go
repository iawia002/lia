package typed

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"

	genericclient "github.com/iawia002/lia/kubernetes/client/generic"
)

// Client is a typed client that allows for querying objects directly through their GVK (group, version, and kind)
// without having to pre-initialize objects of the corresponding type.
type Client interface {
	// Get retrieves an object for the given object key.
	Get(ctx context.Context, key types.NamespacedName, opts ...client.GetOption) (client.Object, error)
	// List retrieves list of objects for a given namespace and list options.
	List(ctx context.Context, namespace string, opts ...client.ListOption) (client.ObjectList, error)
}

type typedClient struct {
	client  client.Reader
	gvk     schema.GroupVersionKind
	obj     runtime.Object
	listObj runtime.Object
}

type options struct {
	config *rest.Config
	cache  client.Reader
	scheme *runtime.Scheme
}

// WithRestConfig is used to pass in the rest config parameter, which is used to generate a client.
// The default value is InClusterConfig.
func WithRestConfig(config *rest.Config) func(opts *options) {
	return func(opts *options) {
		opts.config = config
	}
}

// WithClientReader is used to pass in a client parameter, and you can choose to pass in an existing client or
// use WithRestConfig to pass in the rest config to generate one.
func WithClientReader(cache client.Reader) func(opts *options) {
	return func(opts *options) {
		opts.cache = cache
	}
}

// WithScheme sets the custom scheme for the client.
// The default value is Kubernetes scheme.
func WithScheme(scheme *runtime.Scheme) func(opts *options) {
	return func(opts *options) {
		opts.scheme = scheme
	}
}

// NewTypedClient returns a new Client implementation.
func NewTypedClient(gvk schema.GroupVersionKind, opts ...func(*options)) (Client, error) {
	o := &options{
		scheme: clientgoscheme.Scheme,
	}
	for _, f := range opts {
		f(o)
	}

	if o.cache == nil {
		if o.config == nil {
			inClusterConfig, err := rest.InClusterConfig()
			if err != nil {
				return nil, err
			}
			o.config = inClusterConfig
		}
		cache, err := genericclient.NewClient(o.config, genericclient.WithScheme(o.scheme))
		if err != nil {
			return nil, err
		}
		o.cache = cache
	}

	var (
		obj     runtime.Object
		listObj runtime.Object
	)
	if o.scheme.Recognizes(gvk) {
		obj, _ = o.scheme.New(gvk)
	}
	listGVK := schema.GroupVersionKind{
		Group:   gvk.Group,
		Version: gvk.Version,
		Kind:    gvk.Kind + "List",
	}
	if o.scheme.Recognizes(listGVK) {
		listObj, _ = o.scheme.New(listGVK)
	}

	return &typedClient{
		client:  o.cache,
		gvk:     gvk,
		obj:     obj,
		listObj: listObj,
	}, nil
}

var resourceNotRegisteredError = "kind %s is not registered in scheme"

// Get retrieves an object for the given object key.
func (t *typedClient) Get(ctx context.Context, key types.NamespacedName, opts ...client.GetOption) (client.Object, error) {
	if t.obj == nil {
		return nil, fmt.Errorf(resourceNotRegisteredError, t.gvk.String())
	}

	obj := t.obj.(client.Object)
	if err := t.client.Get(ctx, key, obj, opts...); err != nil {
		return nil, err
	}
	return obj, nil
}

// List retrieves list of objects for a given namespace and list options.
func (t *typedClient) List(ctx context.Context, namespace string, opts ...client.ListOption) (client.ObjectList, error) {
	if t.listObj == nil {
		return nil, fmt.Errorf(resourceNotRegisteredError, t.gvk.String())
	}

	listObj := t.listObj.(client.ObjectList)
	if err := t.client.List(ctx, listObj, append(opts, client.InNamespace(namespace))...); err != nil {
		return nil, err
	}
	return listObj, nil
}
