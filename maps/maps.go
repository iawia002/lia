package maps

import (
	"strings"
)

// SetNestedField sets the value of a nested field (eg: a.b.c) for the source map.
//
// eg:
//
//	src := make(map[string]interface{})
//	SetNestedField(src, "a.b", 1)
//
// src: map[a:map[b:1]]
func SetNestedField(src map[string]interface{}, key string, value interface{}) {
	keys := strings.Split(key, ".")
	l := len(keys) - 1
	currentMap := src
	// Move to the target level
	for _, k := range keys[:l] {
		if _, ok := currentMap[k].(map[string]interface{}); !ok {
			currentMap[k] = make(map[string]interface{})
		}
		currentMap = currentMap[k].(map[string]interface{})
	}
	currentMap[keys[l]] = value
}

// GetNestedField returns the value of a nested field (eg: a.b.c) from the source map.
//
// eg:
//
//	src := map[string]interface{}{
//		"aa": map[string]interface{}{
//			"bb": 1,
//		},
//	}
//	GetNestedField(src, "aa.bb")
func GetNestedField(src map[string]interface{}, key string) interface{} {
	keys := strings.Split(key, ".")
	l := len(keys) - 1
	currentMap := src
	for _, k := range keys[:l] {
		nestedMap, ok := currentMap[k].(map[string]interface{})
		if !ok {
			// Not a map, cannot enter the next level
			return nil
		}
		currentMap = nestedMap
	}
	return currentMap[keys[l]]
}
