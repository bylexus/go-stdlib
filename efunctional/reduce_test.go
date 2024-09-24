package efunctional_test

import (
	"strconv"
	"testing"

	"github.com/bylexus/go-stdlib/efunctional"
)

func TestReduce(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// test with []int -> int
	exp := 55
	res := efunctional.Reduce(data, func(acc, el int) int { return acc + el }, 0)
	if res != exp {
		t.Errorf("Reduce: int wrong result: %d != %d\n", res, exp)
	}

	// test with []int -> string
	expStr := "12345678910"
	resStr := efunctional.Reduce(data, func(acc string, el int) string { return acc + strconv.Itoa(el) }, "")
	if resStr != expStr {
		t.Errorf("Reduce: string wrong result: %s != %s\n", resStr, expStr)
	}
}
