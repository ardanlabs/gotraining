package pdf

import (
	"compress/lzw"
	"compress/zlib"
	"io/ioutil"
	"testing"
)

const streamTestString = "Hello, 世界!\n"

func TestUnfilteredStream(t *testing.T) {
	st := newStream(streamNoFilter)
	st.WriteString(streamTestString)
	st.Close()

	if st.String() != streamTestString {
		t.Errorf("Stream is %q, wanted %q", st.String(), streamTestString)
	}
}

func TestLZWStream(t *testing.T) {
	st := newStream(streamLZWDecode)
	st.WriteString(streamTestString)
	st.Close()

	output, _ := ioutil.ReadAll(lzw.NewReader(st, lzw.MSB, 8))
	if string(output) != streamTestString {
		t.Errorf("Stream is %q, wanted %q", output, streamTestString)
	}
}

func TestFlateStream(t *testing.T) {
	st := newStream(streamFlateDecode)
	st.WriteString(streamTestString)
	st.Close()

	r, _ := zlib.NewReader(st)
	output, _ := ioutil.ReadAll(r)
	if string(output) != streamTestString {
		t.Errorf("Stream is %q, wanted %q", output, streamTestString)
	}
}

const expectedMarshalStreamOutput = "<< /Length 15 >> stream\r\n" + streamTestString + "\r\nendstream"

func TestMarshalStream(t *testing.T) {
	b, err := marshalStream(nil, streamInfo{Length: len(streamTestString)}, []byte(streamTestString))
	if err == nil {
		if string(b) != expectedMarshalStreamOutput {
			t.Errorf("marshalStream(...) != %q (got %q)", expectedMarshalStreamOutput, b)
		}
	} else {
		t.Errorf("Error: %v", err)
	}
}
