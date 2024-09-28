package efunctional

import (
	"iter"

	"github.com/bylexus/go-stdlib/eslices"
)

// Map applies the function mapFn to each element of the slice data and returns the result in a new slice.
//
// Example:
//
//	var doubled []int = efunctional.Map([]int{1, 2, 3, 4}, func(el int) int { return el * 2 })
func Map[T any, U any](data []T, mapFn func(T) U) []U {
	res := make([]U, len(data))
	idx := 0
	for u := range MapIter(eslices.ToIter(data), func(el T) U { return mapFn(el) }) {
		res[idx] = u
		idx++
	}

	return res
}

// MapIter iterates over an iterator, and yields a new iterator that yields
// the result of applying the function mapFn to each element.
//
// Example:
//
//	// Convert a sequence of integers to strings:
//	var intSeq = eslices.ToIter([]int{1, 2, 3, 4})
//	for v := range efunctional.MapIter(intSeq, func(el int) string { return strconv.Itoa(el) }) {
//		fmt.Printf("%s ", v)
//	}
//	// Outputs: 1 2 3 4
func MapIter[T any, U any](data iter.Seq[T], mapFn func(T) U) iter.Seq[U] {
	return func(yield func(U) bool) {
		for el := range data {
			if !yield(mapFn(el)) {
				return
			}
		}
	}
}

// Same as MapIter, but takes and yields two values at a time. Useful if one of the values is an error.
func MapIter2[T, U, V, W any](data iter.Seq2[T, U], mapFn func(T, U) (V, W)) iter.Seq2[V, W] {
	return func(yield func(V, W) bool) {
		for t, u := range data {
			if !yield(mapFn(t, u)) {
				return
			}
		}
	}
}

// Converts a iter.Seq to a iter.Seq2, by mapping each element to a tuple of two elements.
func MapIterToIter2[T any, V any, U any](data iter.Seq[T], mapFn func(T) (V, U)) iter.Seq2[V, U] {
	return func(yield func(V, U) bool) {
		for el := range data {
			v, u := mapFn(el)
			if !yield(v, u) {
				return
			}
		}
	}
}

// Converts a iter.Seq222 to a iter.Seq, by mapping two values to a single value while yielding the result.
func MapIter2ToIter[T any, V any, U any](data iter.Seq2[T, V], mapFn func(T, V) U) iter.Seq[U] {
	return func(yield func(U) bool) {
		for t, v := range data {
			u := mapFn(t, v)
			if !yield(u) {
				return
			}
		}
	}
}
