package eslices_test

import (
	"testing"

	"github.com/bylexus/go-stdlib/eslices"
)

func TestToIter(t *testing.T) {
	testSl := []string{"a", "b", "c", "d", "e"}
	outputSl := make([]string, 0, len(testSl))

	for s := range eslices.ToIter(testSl) {
		outputSl = append(outputSl, s)
	}

	if len(outputSl) != len(testSl) {
		t.Errorf("expected %d, got %d", len(testSl), len(outputSl))
	}

	for i := 0; i < len(testSl); i++ {
		if testSl[i] != outputSl[i] {
			t.Errorf("expected %s, got %s", testSl[i], outputSl[i])
		}
	}

	// stop the iterator after 2 elements:
	outputSl = make([]string, 0, len(testSl))
	for s := range eslices.ToIter(testSl) {
		outputSl = append(outputSl, s)
		if len(outputSl) == 2 {
			break
		}
	}
	if len(outputSl) != 2 {
		t.Errorf("expected %d, got %d", 2, len(outputSl))
	}
	if outputSl[0] != "a" || outputSl[1] != "b" {
		t.Errorf("expected %s, %s, got %s, %s", "a", "b", outputSl[0], outputSl[1])
	}
}

func TestToIter2(t *testing.T) {
	testSl := []string{"a", "b", "c", "d", "e"}
	outputSl := make([]string, 0, len(testSl))
	outputIdx := make([]int, 0, len(testSl))

	for i, s := range eslices.ToIter2(testSl) {
		outputSl = append(outputSl, s)
		outputIdx = append(outputIdx, i)
	}

	if len(outputSl) != len(testSl) {
		t.Errorf("expected strings: %d, got %d", len(testSl), len(outputSl))
	}
	if len(outputIdx) != len(testSl) {
		t.Errorf("expected indexes: %d, got %d", len(testSl), len(outputIdx))
	}

	for i := 0; i < len(testSl); i++ {
		if testSl[i] != outputSl[i] {
			t.Errorf("expected string %s, got %s", testSl[i], outputSl[i])
		}
		if i != outputIdx[i] {
			t.Errorf("expected index %d, got %d", i, outputIdx[i])
		}
	}

	// stop the iterator after 2 elements:
	outputSl = make([]string, 0, len(testSl))
	outputIdx = make([]int, 0, len(testSl))
	for i, s := range eslices.ToIter2(testSl) {
		outputSl = append(outputSl, s)
		outputIdx = append(outputIdx, i)
		if len(outputSl) == 2 {
			break
		}
	}
	if len(outputSl) != 2 {
		t.Errorf("expected strings %d, got %d", 2, len(outputSl))
	}
	if len(outputIdx) != 2 {
		t.Errorf("expected indexes %d, got %d", 2, len(outputSl))
	}
	if outputSl[0] != "a" || outputSl[1] != "b" {
		t.Errorf("expected string %s, %s, got %s, %s", "a", "b", outputSl[0], outputSl[1])
	}
	if outputIdx[0] != 0 || outputIdx[1] != 1 {
		t.Errorf("expected %d, %d, got %d, %d", 0, 1, outputIdx[0], outputIdx[1])
	}
}
