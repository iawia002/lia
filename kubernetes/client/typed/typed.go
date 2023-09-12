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

type typedClient struct {
	client  client.Reader
	scheme  *runtime.Scheme
	gvk     schema.GroupVersionKind
	listGVK schema.GroupVersionKind
}

// NewTypedClient returns a new Client implementation.
func NewTypedClient(gvk schema.GroupVersionKind, opts ...func(*options)) (Client, error) {
	o := &options{}
	for _, f := range opts {
		f(o)
	}

	if o.scheme == nil {
		o.scheme = clientgoscheme.Scheme
	}
	if o.cache == nil {
		if o.config == nil {
			inClusterConfig, err := rest.InClusterConfig()
			if err != nil {
				return nil, err
			}
			o.config = inClusterConfig
		}
		cache, err := genericclient.NewCache(o.config, genericclient.WithScheme(o.scheme))
		if err != nil {
			return nil, err
		}
		o.cache = cache
	}

	return &typedClient{
		client: o.cache,
		scheme: o.scheme,
		gvk:    gvk,
		listGVK: schema.GroupVersionKind{
			Group:   gvk.Group,
			Version: gvk.Version,
			Kind:    gvk.Kind + "List",
		},
	}, nil
}

var resourceNotRegisteredError = "kind %s is not registered in scheme"

// Get retrieves an object for the given object key.
func (t *typedClient) Get(ctx context.Context, key types.NamespacedName, opts ...client.GetOption) (client.Object, error) {
	if !t.scheme.Recognizes(t.gvk) {
		return nil, fmt.Errorf(resourceNotRegisteredError, t.gvk.String())
	}

	obj, err := t.scheme.New(t.gvk)
	if err != nil {
		return nil, err
	}
	clientObj := obj.(client.Object)
	if err = t.client.Get(ctx, key, clientObj, opts...); err != nil {
		return nil, err
	}
	return clientObj, nil
}

// List retrieves list of objects for a given namespace and list options.
func (t *typedClient) List(ctx context.Context, namespace string, opts ...client.ListOption) (client.ObjectList, error) {
	if !t.scheme.Recognizes(t.listGVK) {
		return nil, fmt.Errorf(resourceNotRegisteredError, t.gvk.String())
	}

	listObj, err := t.scheme.New(t.listGVK)
	if err != nil {
		return nil, err
	}
	clientObj := listObj.(client.ObjectList)
	if err = t.client.List(ctx, clientObj, append(opts, client.InNamespace(namespace))...); err != nil {
		return nil, err
	}
	return clientObj, nil
}
