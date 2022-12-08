package typed

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Client is a typed client that allows for querying objects directly through their GVK (group, version, and kind)
// without having to pre-initialize objects of the corresponding type.
type Client interface {
	// Get retrieves an object for the given object key.
	Get(ctx context.Context, key types.NamespacedName, opts ...client.GetOption) (client.Object, error)
	// List retrieves list of objects for a given namespace and list options.
	List(ctx context.Context, namespace string, opts ...client.ListOption) (client.ObjectList, error)
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
