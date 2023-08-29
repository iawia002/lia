package typed

import (
	"context"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"

	genericclient "github.com/iawia002/lia/kubernetes/client/generic"
)

type unstructuredTypedClient struct {
	client  client.Reader
	gvk     schema.GroupVersionKind
	obj     runtime.Object
	listObj runtime.Object
}

// NewUnstructuredTypedClient returns a new Client implementation that returns all objects as Unstructured objects.
func NewUnstructuredTypedClient(gvk schema.GroupVersionKind, opts ...func(*options)) (Client, error) {
	o := &options{}
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
		cache, err := genericclient.NewCache(o.config)
		if err != nil {
			return nil, err
		}
		o.cache = cache
	}

	obj := &unstructured.Unstructured{}
	obj.SetGroupVersionKind(gvk)

	listObj := &unstructured.UnstructuredList{}
	listObj.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   gvk.Group,
		Version: gvk.Version,
		Kind:    gvk.Kind + "List",
	})

	return &unstructuredTypedClient{
		client:  o.cache,
		gvk:     gvk,
		obj:     obj,
		listObj: listObj,
	}, nil
}

// Get retrieves an object for the given object key.
func (t *unstructuredTypedClient) Get(ctx context.Context, key types.NamespacedName, opts ...client.GetOption) (client.Object, error) {
	obj := t.obj.(client.Object)
	if err := t.client.Get(ctx, key, obj, opts...); err != nil {
		return nil, err
	}
	return obj, nil
}

// List retrieves list of objects for a given namespace and list options.
func (t *unstructuredTypedClient) List(ctx context.Context, namespace string, opts ...client.ListOption) (client.ObjectList, error) {
	listObj := t.listObj.(client.ObjectList)
	if err := t.client.List(ctx, listObj, append(opts, client.InNamespace(namespace))...); err != nil {
		return nil, err
	}
	return listObj, nil
}
