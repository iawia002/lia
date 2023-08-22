package generic

import (
	"context"
	"errors"
	"time"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"
)

// Options defines options needed to generate a client.
type Options struct {
	syncPeriod  *time.Duration
	scheme      *runtime.Scheme
	cacheReader bool
	ctx         context.Context
}

// WithSyncPeriod sets the SyncPeriod time option.
// The default value is nil.
func WithSyncPeriod(syncPeriod *time.Duration) func(opts *Options) {
	return func(opts *Options) {
		opts.syncPeriod = syncPeriod
	}
}

// WithScheme sets the custom scheme for the client.
// The default value is Kubernetes scheme.
func WithScheme(scheme *runtime.Scheme) func(opts *Options) {
	return func(opts *Options) {
		opts.scheme = scheme
	}
}

// WithCacheReader sets whether to use the cache reader.
// The default value is true.
func WithCacheReader(cacheReader bool) func(opts *Options) {
	return func(opts *Options) {
		opts.cacheReader = cacheReader
	}
}

// WithContext sets the context for the client.
// The default value is context.Background().
func WithContext(ctx context.Context) func(opts *Options) {
	return func(opts *Options) {
		opts.ctx = ctx
	}
}

// NewClient returns a controller-runtime generic Client implementation.
func NewClient(config *rest.Config, options ...func(*Options)) (client.Client, error) {
	opts := &Options{
		scheme:      scheme.Scheme,
		cacheReader: true,
		ctx:         context.Background(),
	}
	for _, f := range options {
		f(opts)
	}

	httpClient, err := rest.HTTPClientFor(config)
	if err != nil {
		return nil, err
	}
	mapper, err := apiutil.NewDynamicRESTMapper(config, httpClient)
	if err != nil {
		return nil, err
	}

	clientOptions := client.Options{
		HTTPClient: httpClient,
		Scheme:     opts.scheme,
		Mapper:     mapper,
	}

	if opts.cacheReader {
		cacheClient, err := cache.New(config, cache.Options{
			HTTPClient: httpClient,
			Scheme:     opts.scheme,
			Mapper:     mapper,
			SyncPeriod: opts.syncPeriod,
		})
		if err != nil {
			return nil, err
		}
		go cacheClient.Start(opts.ctx) // nolint
		if !cacheClient.WaitForCacheSync(opts.ctx) {
			return nil, errors.New("WaitForCacheSync failed")
		}
		clientOptions.Cache = &client.CacheOptions{
			Reader:       cacheClient,
			Unstructured: true,
		}
	}

	genericClient, err := client.New(config, clientOptions)
	if err != nil {
		return nil, err
	}
	return genericClient, nil
}
