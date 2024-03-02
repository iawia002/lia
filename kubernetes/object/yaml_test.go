package object

import (
	"testing"
)

func TestParseYAMLDocuments(t *testing.T) {
	tests := []struct {
		name   string
		obj    []byte
		isErr  bool
		wanted int
	}{
		{
			name: "single doc test",
			obj: []byte(`
apiVersion: v1
kind: Namespace
metadata:
  name: test
---
`),
			isErr:  false,
			wanted: 1,
		},
		{
			name: "multiple docs test",
			obj: []byte(`---
apiVersion: v1
kind: Namespace
metadata:
  name: test

---
apiVersion: v1
kind: Namespace
metadata:
  name: test2
`),
			isErr:  false,
			wanted: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			docs, err := ParseYAMLDocuments(tt.obj)
			if tt.isErr != (err != nil) {
				t.Errorf("%s ParseYAMLDocuments() unexpected error: %v", tt.name, err)
			}
			if len(docs) != tt.wanted {
				t.Errorf("%s ParseYAMLDocuments() = %v, want %v", tt.name, len(docs), tt.wanted)
			}
		})
	}
}
