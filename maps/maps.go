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
	setValueRecursively(src, strings.Split(key, "."), value)
}

func setValueRecursively(src map[string]interface{}, keys []string, value interface{}) {
	if len(keys) == 1 {
		src[keys[0]] = value
		return
	}

	currentKey := keys[0]
	if _, ok := src[currentKey].(map[string]interface{}); !ok {
		src[currentKey] = make(map[string]interface{})
	}
	setValueRecursively(src[currentKey].(map[string]interface{}), keys[1:], value)
}
