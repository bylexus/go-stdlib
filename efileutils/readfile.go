package efileutils

import (
	"bufio"
	"iter"
	"os"

	"github.com/bylexus/go-stdlib/efunctional"
)

// Returns an iterator that reads from filename, and
// yields each line separately as a byte slice. It reads and yields
// line-by-line, so it does not read the whole file into memory at once.
//
// Example usage:
//
//		for line, err := range efileutils.ReadByteLinesIter(filename) {
//			if err != nil {
//				// handle error
//			} else {
//	            fmt.Println(string(line))
//	        }
//		}
func ReadByteLinesIter(filename string) iter.Seq2[[]byte, error] {
	return func(yield func([]byte, error) bool) {
		f, err := os.Open(filename)
		if err != nil {
			yield(nil, err)
			return
		}
		defer f.Close()

		fs := bufio.NewScanner(f)
		fs.Split(bufio.ScanLines)
		for fs.Scan() {
			err := fs.Err()
			again := true
			if err != nil {
				again = yield(nil, err)
			} else {
				again = yield(fs.Bytes(), nil)
			}
			if !again {
				return
			}
		}
	}
}

// Same as efileutils.ReadByteLinesIter, but yields each line as string instead a byte array.
func ReadStringLinesIter(filename string) iter.Seq2[string, error] {
	return efunctional.MapIter2(
		ReadByteLinesIter(filename),
		func(line []byte, err error) (string, error) { return string(line), err },
	)
}
