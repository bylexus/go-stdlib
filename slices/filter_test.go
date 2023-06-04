package slices_test

import (
	"testing"

	"github.com/bylexus/go-stdlib/slices"
)

type filterTestData struct {
	arr      []string
	filterFn func(el *string) bool
	exp      []string
}

func TestFilter(t *testing.T) {
	var tests = []filterTestData{
		{
			arr:      []string{},
			filterFn: func(el *string) bool { return true },
			exp:      []string{},
		},
		{
			arr:      []string{"a", "b", "c"},
			filterFn: func(el *string) bool { return true },
			exp:      []string{"a", "b", "c"},
		},
		{
			arr:      []string{"a", "b", "c"},
			filterFn: func(el *string) bool { return false },
			exp:      []string{},
		},
		{
			arr:      []string{"a", "b", "c"},
			filterFn: func(el *string) bool { return *el != "b" },
			exp:      []string{"a", "c"},
		},
	}

	for _, data := range tests {
		res := slices.Filter(&data.arr, data.filterFn)
		if len(res) != len(data.exp) {
			t.Errorf("Filter(%v,%p) != %v\n", data.arr, data.filterFn, data.exp)
		}
		if &res == &data.arr {
			t.Errorf("Filter(%v,%p) returned same slice pointer!\n", data.arr, data.filterFn)
		}

		for i := range data.exp {
			if data.exp[i] != res[i] {
				t.Errorf("Filter: element #%d of %v does not match\n", i, res)

			}
		}
	}
}
