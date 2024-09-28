package efunctional

import (
	"iter"

	"github.com/bylexus/go-stdlib/eslices"
)

// Filter returns a new slice by applying the predicate function on each element:
// the predicate need to return true in order to keep the element.
// Note that the new slice has the same capacity (but not necessarily the same length) as the original one.
//
// Example:
//
// To keep only even numbers:
//
//	res := efunctional.Filter([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, func(el int) bool { return el%2 == 0 })
func Filter[T any](data []T, predicate func(T) bool) []T {
	res := make([]T, 0, len(data))

	for v := range FilterIter(eslices.ToIter(data), predicate) {
		res = res[:len(res)+1]
		res[len(res)-1] = v
	}

	return res
}

// FilterIter iterates over an iterator, and yields a new iterator that yields
// only the elements for which the predicate function returns true.
//
// Example:
//
//		 iter := eslices.ToIter([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
//		 for v := range efunctional.FilterIter(iter, func(el int) bool { return el%2 == 0 }) {
//		    fmt.Printf("%d ", v)
//	     }
//		 // Outputs: 2 4 6 8 10
func FilterIter[T any](data iter.Seq[T], predicate func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for el := range data {
			if predicate(el) {
				if !yield(el) {
					return
				}
			}
		}
	}
}

// Same as FilterIter, but with an iterator that yields two values at a time. (iter.Seq2)
func FilterIter2[T, U any](data iter.Seq2[T, U], predicate func(T, U) bool) iter.Seq2[T, U] {
	return func(yield func(T, U) bool) {
		for t, u := range data {
			if predicate(t, u) {
				if !yield(t, u) {
					return
				}
			}
		}
	}
}
