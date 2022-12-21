package design

// Go's io.Reader
type Reader interface {
	Read(p []byte) (n int, err error)
}

// Python's "read" method
type PyReader interface {
	Read(n int) (p []byte, err error)
}
