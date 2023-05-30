package slices

import "testing"

type inSliceTestData struct {
	arr    []string
	search string
	exp    bool
}

func TestInSlice(t *testing.T) {
	var tests = []inSliceTestData{
		{
			arr:    []string{},
			search: "foo",
			exp:    false,
		},
		{
			arr:    []string{},
			search: "",
			exp:    false,
		},
		{
			arr:    []string{""},
			search: "",
			exp:    true,
		},
		{
			arr:    []string{"", "a", "b"},
			search: "c",
			exp:    false,
		},
		{
			arr:    []string{"", "a", "b"},
			search: "b",
			exp:    true,
		},
		{
			arr:    []string{"a", "b", ""},
			search: "",
			exp:    true,
		},
	}

	for _, data := range tests {
		if InSlice(&data.arr, data.search) != data.exp {
			t.Errorf("InSlice(%#v, %s) != %t\n", data.arr, data.search, data.exp)
		}
	}
}
