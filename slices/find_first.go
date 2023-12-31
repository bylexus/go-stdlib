package slices

// Returns the first element from the given slice that matches the predicate.
func FindFirst[T any](slice *[]T, predicate func(*T) bool) *T {
	for i := range *slice {
		if predicate(&(*slice)[i]) {
			return &(*slice)[i]
		}
	}
	return nil
}
