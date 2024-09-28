package efunctional

// reduce accumulates the result of applying the function accFn to each element of the slice data,
// reducing the slice to a single value in the process.
//
// the initial value is the value taken as initial accumulation value.
//
// Example:
//
// To sum the elements of a slice:
//
//	var sum int = efunctional.Reduce([]int{1, 2, 3, 4}, func(acc, el int) int { return acc + el }, 0)
func Reduce[T any, U any](data []T, accFn func(U, T) U, initial U) U {
	acc := initial
	for _, v := range data {
		acc = accFn(acc, v)
	}

	return acc
}
