package efileutils

import (
	"bufio"
	"iter"
	"os"
)

// ReadFileIter returns an iterator that reads from filename.
// Each yield is a byte slice that represents one line of the file.
func ReadFileIter(filename string) iter.Seq2[[]byte, error] {
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
			if err != nil {
				yield(nil, err)
				return
			}
			if !yield(fs.Bytes(), nil) {
				return
			}
		}
	}
}
