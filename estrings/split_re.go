package estrings

import (
	"iter"
	"regexp"

	"github.com/bylexus/go-stdlib/efunctional"
)

func SplitRe(str string, pattern string) ([]string, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	return re.Split(str, -1), nil
}

func SplitReIter(strIter iter.Seq[string], pattern string) iter.Seq2[[]string, error] {
	return efunctional.MapIterToIter2(strIter, func(str string) ([]string, error) {
		return SplitRe(str, pattern)
	})
}

func SplitReErrIter(strIter iter.Seq2[string, error], pattern string) iter.Seq2[[]string, error] {
	return func(yield func([]string, error) bool) {
		for str, err := range strIter {
			if err != nil {
				yield(nil, err)
				return
			}
			splitts, err := SplitRe(str, pattern)
			if err != nil {
				yield(nil, err)
				return
			}
			if !yield(splitts, nil) {
				return
			}
		}
	}
}

func FindStringSubmatchErrIter(strIter iter.Seq2[string, error], pattern string) iter.Seq2[[]string, error] {
	return func(yield func([]string, error) bool) {
		re, err := regexp.Compile(pattern)
		if err != nil {
			yield(nil, err)
			return
		}
		for str, err := range strIter {
			if err != nil {
				yield(nil, err)
				return
			}
			parts := re.FindStringSubmatch(str)
			if !yield(parts, nil) {
				return
			}
		}
	}
}
