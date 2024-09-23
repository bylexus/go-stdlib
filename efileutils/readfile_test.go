package efileutils_test

import (
	"os"
	"path"
	"runtime"
	"strings"
	"testing"

	"github.com/bylexus/go-stdlib/efileutils"
)

func TestReadFileIter_TestFullRead(t *testing.T) {
	// file name of this test:
	_, file, _, _ := runtime.Caller(0)
	inputfile := path.Join(path.Dir(file), "..", "testdata", "plaintext.txt")
	readLines := make([]string, 0)

	for line, err := range efileutils.ReadFileIter(inputfile) {
		if err != nil {
			t.Fatal(err)
		}
		readLines = append(readLines, string(line))
	}
	if len(readLines) != 4 {
		t.Fatalf("expected 4 lines, got %d", len(readLines))
	}

	inputBytes, _ := os.ReadFile(inputfile)
	inputLines := strings.Split(string(inputBytes), "\n")
	for i, line := range inputLines {
		if readLines[i] != line {
			t.Fatalf("expected %q at index %d, got %q", line, i, readLines[i])
		}
	}
}

func TestReadFileIter_TestPartialRead(t *testing.T) {
	// file name of this test:
	_, file, _, _ := runtime.Caller(0)
	inputfile := path.Join(path.Dir(file), "..", "testdata", "plaintext.txt")
	readLines := make([]string, 0)
	counter := 0

	for line, err := range efileutils.ReadFileIter(inputfile) {
		counter++
		if err != nil {
			t.Fatal(err)
		}
		readLines = append(readLines, string(line))
		if counter == 2 {
			break
		}
	}
	if len(readLines) != 2 {
		t.Fatalf("expected 2 lines, got %d", len(readLines))
	}

	inputBytes, _ := os.ReadFile(inputfile)
	inputLines := strings.Split(string(inputBytes), "\n")
	for i, line := range readLines {
		if readLines[i] != line {
			t.Fatalf("expected %q at index %d, got %q", line, i, inputLines[i])
		}
	}
}
