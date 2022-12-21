package design

// Interface say what we need, not what we provide.
// Rule of thumb: Accept interfaces, return types.

// io

func Copy(dst Writer, src Reader) (written int64, err error) {
	// Reducated

	return 0, nil
}

// os

func Open(name string) (*File, error) {
	// Redacted
	return &File{}, nil
}

// Make compiler happy

type File struct {
	// Redacted
}

type Writer interface {
	Write(p []byte) (n int, err error)
}
