package maps

// DeepCopy returns a copy of the original map object.
// Deprecated: Use the built-in maps.Clone function instead
func DeepCopy[K comparable, V any](src map[K]V) map[K]V {
	re := make(map[K]V, len(src))
	for k, v := range src {
		re[k] = v
	}
	return re
}
