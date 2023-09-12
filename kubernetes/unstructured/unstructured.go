package unstructured

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/yaml"
)

// ConvertToUnstructured converts a typed object to an unstructured object.
func ConvertToUnstructured(obj interface{}) (*unstructured.Unstructured, error) {
	unstructuredObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	return &unstructured.Unstructured{Object: unstructuredObj}, nil
}

// ConvertToTyped converts an unstructured object to a typed object.
// Usage:
//
//	node := &corev1.Node{}
//	ConvertToTyped(object, node)
//
//nolint:gofmt,goimports
func ConvertToTyped(obj runtime.Unstructured, typedObj interface{}) error {
	return runtime.DefaultUnstructuredConverter.FromUnstructured(obj.UnstructuredContent(), typedObj)
}

// YAMLToUnstructured converts the object's YAML content into an unstructured object.
func YAMLToUnstructured(content []byte) (*unstructured.Unstructured, error) {
	obj := make(map[string]interface{})
	if err := yaml.Unmarshal(content, &obj); err != nil {
		return nil, err
	}
	return &unstructured.Unstructured{
		Object: obj,
	}, nil
}
