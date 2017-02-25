## Writes/Reader Interfaces - Standard Library

### Writer

	type Writer interface {
	        Write(p []byte) (n int, err error)
	}

* The implementation of the Write method should attempt to write the entire length
  of the byte slice that is passed in.

* If the entire slice could not be written, the number of bytes successfully written
  along with a non-nil error must be returned.

* Otherwise, when completely successful, the returned `n` must equal `len(p)` and a
  nil error should be returned.

* The Writer implementation must never modify the provided byte slice at any time.

### Reader

	type Reader interface {
	        Read(p []byte) (n int, err error)
	}

* The implementation should fill the provided byte slice with all the data that is
  immediately available, up to `len(p)` bytes.

* If no data is immediately available, Read should block.

* `io.EOF` is a special `error` value used to signal an end-of-file condition: no more
  data will follow.

* When an error or EOF condition occurs, Reader implementations can choose either
  to return the error value immediately from that same Read call, or may wait and return
  `n == 0` along with the error on the next Read call. Code using Readers should be able
  to handle both situations.

* Anytime the Read method returns bytes, those bytes should be processed first before
  checking the error value for an EOF or other error value.

* The implementation must never return a 0 byte read count with an error value of nil.
  Reads that result in no bytes read should always return an error.

___
[![Ardan Labs](../../../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../../../00-slides/images/ardan_logo.png)](http://www.ardanlabs.com)
[![GoingGo Blog](../../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
