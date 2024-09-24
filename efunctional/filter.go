package efunctional

// Filter returns a new slice by applying the predicate function on each element:
// the predicate need to return true in order to keep the element.
// Note that the new slice has the same capacity (but not necessarily the same length) as the original one.
func Filter[T any](data []T, predicate func(T) bool) []T {
	res := make([]T, 0, len(data))
	for _, v := range data {
		if predicate(v) {
			res = res[:len(res)+1]
			res[len(res)-1] = v
		}
	}

	return res
}
