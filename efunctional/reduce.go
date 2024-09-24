package efunctional

// reduce accumulates the result of applying the function f to each element of the slice s,
// reducing the slice to a single value in the process.
//
// the initial value is the value taken as initial accumulation value.
//
// Example:
//
// To sum the elements of a slice:
//
//	Reduce([]int{1, 2, 3, 4}, func(acc, el int) int { return acc + el }, 0)
func Reduce[T any, U any](s []T, f func(U, T) U, initial U) U {
	acc := initial
	for _, v := range s {
		acc = f(acc, v)
	}

	return acc
}
