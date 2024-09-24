package efunctional

// Map applies the function f to each element of the slice s and returns the result in a new slice.
func Map[T any, U any](s []T, f func(T) U) []U {
	res := make([]U, len(s))
	for i, v := range s {
		res[i] = f(v)
	}

	return res
}
