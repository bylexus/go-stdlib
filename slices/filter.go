package slices

/*
Takes a slice pointer and a predicate function, and returns
a new slice filtered by the predicate function:
Each element is passed to the predicate fn, which must return
true to keep the element, false to discard it.

Example usage:

	data := []string{"a", "b", "c"}
	filtered := slices.Filter(&data, func(el *string) bool { return *el != "b"}) // returns []string {"a", "c"}
*/
func Filter[T any](slice *[]T, predicate func(*T) bool) []T {
	ret := make([]T, 0)
	for _, el := range *slice {
		if predicate(&el) {
			ret = append(ret, el)
		}
	}
	return ret
}
