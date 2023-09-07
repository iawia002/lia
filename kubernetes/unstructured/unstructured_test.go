package unstructured

import (
	"testing"
)

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
