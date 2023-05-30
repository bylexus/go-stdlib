package strings

import (
	"testing"
)

func TestSplitRe(t *testing.T) {
	testStr := "foo, bar baz ; too ;,"
	testRe := `[,\s;]+`
	exp := []string{"foo", "bar", "baz", "too", ""}

	res, err := SplitRe(testStr, testRe)
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
