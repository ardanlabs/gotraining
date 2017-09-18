// Copyright (C) 2011, Ross Light

package pdf

import (
	"bytes"
	"reflect"
	"testing"
)

const encodingTestData = "%PDF-1.7\r\n" +
	"%\x93\x8c\x8b\x9e\r\n" +
	"1 0 obj\r\n" +
	"(Hello, World!)\r\n" +
	"endobj\r\n" +
	"2 0 obj\r\n" +
	"42\r\n" +
	"endobj\r\n" +
	"xref\r\n" +
	"0 3\r\n" +
	"0000000000 65535 f\r\n" +
	"0000000017 00000 n\r\n" +
	"0000000051 00000 n\r\n" +
	"trailer\r\n" +
	"<< /Size 3 /Root 0 0 R >>\r\n" +
	"startxref\r\n" +
	"72\r\n" +
	"%%EOF\r\n"

func TestEncoder(t *testing.T) {
	var e encoder
	if ref := e.add("Hello, World!"); !reflect.DeepEqual(ref, Reference{1, 0}) {
		t.Errorf("After adding first object, reference is %#v", ref)
	}
	if ref := e.add(42); !reflect.DeepEqual(ref, Reference{2, 0}) {
		t.Errorf("After adding second object, reference is %#v", ref)
	}

	var b bytes.Buffer
	if err := e.encode(&b); err != nil {
		t.Fatalf("Encoding error: %v", err)
	}
	if b.String() != encodingTestData {
		t.Errorf("Encoding result %q, want %q", b.String(), encodingTestData)
	}
}
