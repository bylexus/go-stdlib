package eslices_test

import (
	"testing"

	"github.com/bylexus/go-stdlib/eslices"
)

type findFirstTestData struct {
	arr      []string
	searchFn func(el *string) bool
	exp      *string
}

func TestFindFirst(t *testing.T) {
	var data = []string{
		"alex", "blex", "clex",
	}
	var tests = []findFirstTestData{
		{
			arr:      []string{},
			searchFn: func(el *string) bool { return true },
			exp:      nil,
		},
		{
			arr:      []string{},
			searchFn: func(el *string) bool { return true },
			exp:      nil,
		},
		{
			arr:      data,
			searchFn: func(el *string) bool { return *el == "" },
			exp:      nil,
		},
		{
			arr:      data,
			searchFn: func(el *string) bool { return *el == "blex" },
			exp:      &data[1],
		},
		{
			arr:      data,
			searchFn: func(el *string) bool { return *el == "dlex" },
			exp:      nil,
		},
		{
			arr:      data,
			searchFn: func(el *string) bool { return true },
			exp:      &data[0],
		},
	}

	for _, test := range tests {
		if eslices.FindFirst(&test.arr, test.searchFn) != test.exp {
			t.Errorf("FindFirst(%#v, %p) != %s\n", test.arr, test.searchFn, *test.exp)
		}
	}
}
