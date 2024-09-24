package efunctional_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/bylexus/go-stdlib/efunctional"
)

func TestMap(t *testing.T) {
	data := []int{1, 2, 3}
	exp := []string{"11", "22", "33"}

	res := efunctional.Map(data, func(el int) string { return fmt.Sprintf("%s%s", strconv.Itoa(el), strconv.Itoa(el)) })

	if (len(res) != len(exp)) || (res[0] != exp[0]) || (res[1] != exp[1]) || (res[2] != exp[2]) {
		t.Errorf("map wrong: %v != %v\n", res, exp)
	}
}
