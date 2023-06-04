package strings_test

import (
	"testing"

	"github.com/bylexus/go-stdlib/strings"
)

func TestSplitRe(t *testing.T) {
	testStr := "foo, bar baz ; too ;,"
	testRe := `[,\s;]+`
	exp := []string{"foo", "bar", "baz", "too", ""}

	res, err := strings.SplitRe(testStr, testRe)
	if err != nil {
		t.Errorf("split unsuccessful: %s\n", err)
	}
	if len(res) != len(exp) {
		t.Errorf("split unsuccessful: %s by %s ==> %#v\n", testStr, testRe, res)
	}

	for i, expStr := range exp {
		if res[i] != expStr {
			t.Errorf("split wrong: %d: %s != %s\n", i, expStr, res[i])
		}
	}
}
