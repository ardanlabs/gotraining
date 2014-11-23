## Writes/Reader Interfaces - Standard Library

### Writer

	type Writer interface {
	        Write(p []byte) (n int, err error)
	}

* The implementation of the Write method should attempt to write the entire length
  of the byte slice that is passed in.

* If that is not possible, then at least 1 byte must be written or the method must
  return an error.

* The number of bytes reported as written can be less then the length of the byte
  slice but never more.

* The byte slice must never be modify in any way.

### Reader

	type Reader interface {
	        Read(p []byte) (n int, err error)
	}

* The implementation should attempt to read the entire length of the byte slice that
  is passed in. It is ok to read less than the entire length and it should not wait
  to read the entire length if that much data is not available at the time of the call.

* When the last byte is read, two options are available. Either Read returns the final
  bytes with the proper count and EOF for the error value or returns the final bytes
  with the proper count and nil for the error value. In the latter case, the next read
  must return no bytes with the count of 0 and EOF for the error value.

* Anytime the Read method returns bytes, those bytes should be processed first before
  checking the error value for an EOF or other error value.

* The implementation must never return a 0 byte read count with an error value of nil.
  Reads that result in no byte read should always return an error.

___
[![GoingGo Training](../../../00-slides/images/ggt_logo.png)](http://www.goinggotraining.net)
[![Ardan Studios](../../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [GNU Free Documentation License](https://github.com/ArdanStudios/gotraining/blob/master/LICENSE).