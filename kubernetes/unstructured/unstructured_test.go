package unstructured

import (
	"testing"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

func TestConvertToUnstructured(t *testing.T) {
	tests := []struct {
		name  string
		obj   interface{}
		isErr bool
	}{
		{
			name: "obj test",
			obj: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test",
				},
			},
			isErr: false,
		},
		{
			name: "list test",
			obj: &corev1.NamespaceList{
				Items: []corev1.Namespace{
					{
						ObjectMeta: metav1.ObjectMeta{
							Name: "test",
						},
					},
				},
			},
			isErr: false,
		},
		{
			name: "error test",
			obj: []byte(`
apiVersion: v1
kind: Namespace
metadata:
  name: test
`),
			isErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := ConvertToUnstructured(tt.obj); tt.isErr != (err != nil) {
				t.Errorf("%s ConvertToUnstructured() unexpected error: %v", tt.name, err)
			}
		})
	}
}

func TestConvertToTyped(t *testing.T) {
	tests := []struct {
		name     string
		obj      runtime.Unstructured
		typedObj interface{}
		isErr    bool
	}{
		{
			name: "obj test",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind": "Namespace",
					"metadata": map[string]interface{}{
						"name": "test",
					},
				},
			},
			typedObj: &corev1.Namespace{},
			isErr:    false,
		},
		{
			name: "list test",
			obj: &unstructured.UnstructuredList{
				Object: map[string]interface{}{
					"kind": "List",
				},
				Items: []unstructured.Unstructured{
					{
						Object: map[string]interface{}{
							"kind": "Namespace",
							"metadata": map[string]interface{}{
								"name": "test",
							},
						},
					},
				},
			},
			typedObj: &corev1.NamespaceList{},
			isErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ConvertToTyped(tt.obj, tt.typedObj); tt.isErr != (err != nil) {
				t.Errorf("%s ConvertToTyped() unexpected error: %v", tt.name, err)
			}
		})
	}
}

func TestYAMLToUnstructured(t *testing.T) {
	tests := []struct {
		name  string
		obj   []byte
		isErr bool
	}{
		{
			name: "normal test",
			obj: []byte(`
apiVersion: v1
kind: Namespace
metadata:
  name: test
`),
			isErr: false,
		},
		{
			name:  "error test",
			obj:   []byte(`aaa`),
			isErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := YAMLToUnstructured(tt.obj); tt.isErr != (err != nil) {
				t.Errorf("%s YAMLToUnstructured() unexpected error: %v", tt.name, err)
			}
		})
	}
}
