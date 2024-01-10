package emaps

// Returns a slice of keys of the given map.
func GetMapKeys[T comparable, U any](m *map[T]U) []T {
	keys := make([]T, 0)
	for k := range *m {
		keys = append(keys, k)
	}
	return keys
}
