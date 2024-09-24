package efunctional_test

import (
	"testing"

	"github.com/bylexus/go-stdlib/efunctional"
)

func TestFilter(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	exp := []int{2, 4, 6, 8, 10}

	res := efunctional.Filter(data, func(el int) bool { return el%2 == 0 })

	if len(res) != len(exp) || res[0] != exp[0] || res[1] != exp[1] || res[2] != exp[2] || res[3] != exp[3] || res[4] != exp[4] {
		t.Errorf("Filter: wrong filtering: %v != %v\n", res, exp)
	}
	if cap(res) != cap(data) {
		t.Errorf("Filter: wrong capacity: %d != %d\n", cap(res), cap(data))
	}
}
