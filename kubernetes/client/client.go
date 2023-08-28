package client

import (
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

const (
	// DefaultQPS is the default QPS value.
	DefaultQPS = 50
	// DefaultBurst is the default Burst value.
	DefaultBurst = 100
)

// SetQPS sets the QPS and Burst.
func SetQPS(qps float32, burst int) func(c *rest.Config) {
	return func(c *rest.Config) {
		c.QPS = qps
		c.Burst = burst
	}
}

// BuildConfigFromFlags builds rest configs from a master url or a kube config filepath.
func BuildConfigFromFlags(masterURL, kubeConfigPath string, options ...func(c *rest.Config)) (*rest.Config, error) {
	restConfig, err := clientcmd.BuildConfigFromFlags(masterURL, kubeConfigPath)
	if err != nil {
		return nil, err
	}

	for _, opt := range options {
		opt(restConfig)
	}
	if restConfig.QPS == 0 {
		restConfig.QPS = DefaultQPS
	}
	if restConfig.Burst == 0 {
		restConfig.Burst = DefaultBurst
	}
	return restConfig, nil
}

// BuildConfigFromKubeConfig builds rest configs from kube config data.
func BuildConfigFromKubeConfig(kubeconfig []byte, options ...func(c *rest.Config)) (*rest.Config, error) {
	restConfig, err := clientcmd.RESTConfigFromKubeConfig(kubeconfig)
	if err != nil {
		return nil, err
	}

	for _, opt := range options {
		opt(restConfig)
	}
	if restConfig.QPS == 0 {
		restConfig.QPS = DefaultQPS
	}
	if restConfig.Burst == 0 {
		restConfig.Burst = DefaultBurst
	}
	return restConfig, nil
}
