package eslices

import "iter"

func ToIter[T any](data []T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, el := range data {
			if !yield(el) {
				return
			}
		}
	}

}

func ToIter2[T any](data []T) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i, el := range data {
			if !yield(i, el) {
				return
			}
		}
	}
}
