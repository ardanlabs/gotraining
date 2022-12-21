package wc

import (
	"compress/gzip"
	"io"
	"os"
	"strings"
)

func LineCount(fileName string) (int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	var r io.Reader = file

	if strings.HasSuffix(fileName, ".gz") {
		r, err = gzip.NewReader(file)
		if err != nil {
			return 0, err
		}
	}

	var lc LineCounter
	if _, err := io.Copy(&lc, r); err != nil {
		return 0, err
	}

	return int(lc), nil
}

type LineCounter int

func (lc *LineCounter) Write(data []byte) (int, error) {
	for _, c := range data {
		if c == '\n' {
			*lc++
		}
	}

	return len(data), nil
}
