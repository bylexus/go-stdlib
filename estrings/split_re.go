package estrings

import "regexp"

func SplitRe(str string, pattern string) ([]string, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	return re.Split(str, -1), nil
}
