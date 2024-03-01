package generic

import (
	"context"
	"errors"
	"net/http"
	"time"

	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	toolscache "k8s.io/client-go/tools/cache"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"
)

// Options defines options needed to generate a client.
type Options struct {
	syncPeriod        *time.Duration
	scheme            *runtime.Scheme
	cacheReader       bool
	ctx               context.Context
	httpClient        *http.Client
	mapper            meta.RESTMapper
	defaultNamespaces map[string]cache.Config
	defaultTransform  toolscache.TransformFunc
	byObject          map[client.Object]cache.ByObject
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

// WithHTTPClient sets the HTTPClient for the client.
func WithHTTPClient(httpClient *http.Client) func(opts *Options) {
	return func(opts *Options) {
		opts.httpClient = httpClient
	}
}

// WithMapper sets the Mapper for the client.
func WithMapper(mapper meta.RESTMapper) func(opts *Options) {
	return func(opts *Options) {
		opts.mapper = mapper
	}
}

// WithDefaultNamespaces sets the DefaultNamespaces for the cache client.
func WithDefaultNamespaces(defaultNamespaces map[string]cache.Config) func(opts *Options) {
	return func(opts *Options) {
		opts.defaultNamespaces = defaultNamespaces
	}
}

// WithDefaultTransform sets the DefaultTransform for the cache client.
func WithDefaultTransform(defaultTransform toolscache.TransformFunc) func(opts *Options) {
	return func(opts *Options) {
		opts.defaultTransform = defaultTransform
	}
}

// WithByObject sets the ByObject for the cache client.
func WithByObject(byObject map[client.Object]cache.ByObject) func(opts *Options) {
	return func(opts *Options) {
		opts.byObject = byObject
	}
}

// NewCache returns a controller-runtime cache client implementation.
func NewCache(config *rest.Config, options ...func(*Options)) (cache.Cache, error) {
	opts := &Options{
		scheme: scheme.Scheme,
		ctx:    context.Background(),
	}
	for _, f := range options {
		f(opts)
	}

	if opts.httpClient == nil {
		httpClient, err := rest.HTTPClientFor(config)
		if err != nil {
			return nil, err
		}
		opts.httpClient = httpClient
	}
	if opts.mapper == nil {
		mapper, err := apiutil.NewDynamicRESTMapper(config, opts.httpClient)
		if err != nil {
			return nil, err
		}
		opts.mapper = mapper
	}

	cacheClient, err := cache.New(config, cache.Options{
		HTTPClient:        opts.httpClient,
		Scheme:            opts.scheme,
		Mapper:            opts.mapper,
		SyncPeriod:        opts.syncPeriod,
		DefaultNamespaces: opts.defaultNamespaces,
		DefaultTransform:  opts.defaultTransform,
		ByObject:          opts.byObject,
	})
	if err != nil {
		return nil, err
	}
	go cacheClient.Start(opts.ctx) // nolint
	if !cacheClient.WaitForCacheSync(opts.ctx) {
		return nil, errors.New("WaitForCacheSync failed")
	}
	return cacheClient, nil
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

	if opts.httpClient == nil {
		httpClient, err := rest.HTTPClientFor(config)
		if err != nil {
			return nil, err
		}
		opts.httpClient = httpClient
	}
	if opts.mapper == nil {
		mapper, err := apiutil.NewDynamicRESTMapper(config, opts.httpClient)
		if err != nil {
			return nil, err
		}
		opts.mapper = mapper
	}

	clientOptions := client.Options{
		HTTPClient: opts.httpClient,
		Scheme:     opts.scheme,
		Mapper:     opts.mapper,
	}

	if opts.cacheReader {
		cacheClient, err := NewCache(config,
			WithHTTPClient(opts.httpClient),
			WithScheme(opts.scheme),
			WithMapper(opts.mapper),
			WithSyncPeriod(opts.syncPeriod),
			WithContext(opts.ctx),
			WithDefaultNamespaces(opts.defaultNamespaces),
			WithDefaultTransform(opts.defaultTransform),
			WithByObject(opts.byObject),
		)
		if err != nil {
			return nil, err
		}
		clientOptions.Cache = &client.CacheOptions{
			Reader:       cacheClient,
			Unstructured: true,
		}
	}
	return client.New(config, clientOptions)
}
