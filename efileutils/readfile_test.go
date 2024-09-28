package efileutils_test

import (
	"os"
	"path"
	"reflect"
	"regexp"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/bylexus/go-stdlib/efileutils"
	"github.com/bylexus/go-stdlib/efunctional"
	"github.com/bylexus/go-stdlib/estrings"
)

func TestReadFileIter_TestFullRead(t *testing.T) {
	// file name of this test:
	_, file, _, _ := runtime.Caller(0)
	inputfile := path.Join(path.Dir(file), "..", "testdata", "plaintext.txt")
	readLines := make([]string, 0)

	for line, err := range efileutils.ReadByteLinesIter(inputfile) {
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

	for line, err := range efileutils.ReadByteLinesIter(inputfile) {
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

func TestReadFileSplitReIter_TestFullRead(t *testing.T) {
	// file name of this test:
	_, file, _, _ := runtime.Caller(0)
	inputfile := path.Join(path.Dir(file), "..", "testdata", "plaintext.txt")
	readWords := make([][]string, 0)

	fileIter := estrings.SplitReErrIter(efileutils.ReadStringLinesIter(inputfile), "\\s+")

	for words, err := range fileIter {
		if err != nil {
			t.Fatal(err)
		}
		readWords = append(readWords, words)
	}

	// check if the number of lines is correct
	if len(readWords) != 4 {
		t.Fatalf("expected 4 lines, got %d", len(readWords))
	}

	// check if we have all the words split:
	expectedData, _ := os.ReadFile(inputfile)
	expectedLines, _ := estrings.SplitRe(string(expectedData), "\\n")
	expectedWords := efunctional.Map(expectedLines[:], func(s string) []string { r, _ := estrings.SplitRe(s, "\\s+"); return r })
	if !reflect.DeepEqual(readWords, expectedWords) {
		t.Fatalf("expected %v, got %v", expectedWords, readWords)
	}
}

func TestLogfileParsing(t *testing.T) {
	// file name of this test:
	_, file, _, _ := runtime.Caller(0)
	inputfile := path.Join(path.Dir(file), "..", "testdata", "logfile.txt")
	lineMatch := regexp.MustCompile("DELETE FROM")
	counter := 0

	// read lines, filter by lines with "DELETE FROM", extract timestamp from line
	for parts, err := range estrings.FindStringSubmatchErrIter(
		efunctional.FilterIter2(
			efileutils.ReadStringLinesIter(inputfile),
			func(line string, err error) bool { return lineMatch.MatchString(line) }),
		`^\[(.*?)\]\s+main\.DEBUG:.*"type":"([^"]+)"`) {

		counter++

		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v", parts)
		if len(parts) != 3 {
			t.Fatalf("expected 3 parts, got %d", len(parts))
		}

		_, err := time.Parse(time.RFC3339, parts[1])
		if err != nil {
			t.Fatal(err)
		}

		if parts[2] != "DATABASE" {
			t.Fatalf("expected DATABASE, got %s", parts[2])
		}
	}
	if counter != 5 {
		t.Fatalf("expected 6 lines, got %d", counter)
	}
}
