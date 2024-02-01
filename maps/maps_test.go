package maps

import (
	"reflect"
	"testing"
)

func TestSetNestedField(t *testing.T) {
	tests := []struct {
		name   string
		src    map[string]interface{}
		key    string
		value  interface{}
		wanted map[string]interface{}
	}{
		{
			name: "normal test",
			src: map[string]interface{}{
				"aa": 1,
			},
			key:   "bb.cc",
			value: 1,
			wanted: map[string]interface{}{
				"aa": 1,
				"bb": map[string]interface{}{
					"cc": 1,
				},
			},
		},
		{
			name: "normal test 2",
			src: map[string]interface{}{
				"aa": 1,
				"bb": map[string]interface{}{},
			},
			key:   "bb.cc",
			value: 1,
			wanted: map[string]interface{}{
				"aa": 1,
				"bb": map[string]interface{}{
					"cc": 1,
				},
			},
		},
		{
			name: "override test",
			src: map[string]interface{}{
				"aa": 1,
				"bb": 1,
			},
			key:   "bb.cc",
			value: 1,
			wanted: map[string]interface{}{
				"aa": 1,
				"bb": map[string]interface{}{
					"cc": 1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetNestedField(tt.src, tt.key, tt.value)
			if !reflect.DeepEqual(tt.src, tt.wanted) {
				t.Errorf("SetNestedField() = %v, want %v", tt.src, tt.wanted)
			}
		})
	}
}
