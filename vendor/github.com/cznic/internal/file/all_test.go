// Copyright 2016 The Internal Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package file

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"runtime"
	"strings"
	"testing"

	"github.com/cznic/fileutil"
	"github.com/cznic/internal/buffer"
)

func caller(s string, va ...interface{}) {
	if s == "" {
		s = strings.Repeat("%v ", len(va))
	}
	_, fn, fl, _ := runtime.Caller(2)
	fmt.Fprintf(os.Stderr, "caller: %s:%d: ", path.Base(fn), fl)
	fmt.Fprintf(os.Stderr, s, va...)
	fmt.Fprintln(os.Stderr)
	_, fn, fl, _ = runtime.Caller(1)
	fmt.Fprintf(os.Stderr, "\tcallee: %s:%d: ", path.Base(fn), fl)
	fmt.Fprintln(os.Stderr)
	os.Stderr.Sync()
}

func dbg(s string, va ...interface{}) {
	if s == "" {
		s = strings.Repeat("%v ", len(va))
	}
	_, fn, fl, _ := runtime.Caller(1)
	fmt.Fprintf(os.Stderr, "dbg %s:%d: ", path.Base(fn), fl)
	fmt.Fprintf(os.Stderr, s, va...)
	fmt.Fprintln(os.Stderr)
	os.Stderr.Sync()
}

func TODO(...interface{}) string { //TODOOK
	_, fn, fl, _ := runtime.Caller(1)
	return fmt.Sprintf("TODO: %s:%d:\n", path.Base(fn), fl) //TODOOK
}

func use(...interface{}) {}

func init() {
	use(caller, dbg, TODO) //TODOOK
}

// ============================================================================

const testSize = 1 << 26

var (
	_ Interface = (*osFile)(nil)

	_ testInterface = (*testFile)(nil)
	_ testInterface = (*testMem)(nil)
	_ testInterface = (*testOSFile)(nil)
)

type testInterface interface {
	new() (Interface, error)
	free() error
}

type osFile struct {
	*os.File
}

func (f *osFile) ReadFrom(r io.Reader) (n int64, err error) { return readFrom(f, r) }
func (f *osFile) WriteTo(w io.Writer) (n int64, err error)  { return writeTo(f, w) }

type testFile struct {
	f    Interface
	name string
}

func (f *testFile) new() (Interface, error) {
	f0, err := ioutil.TempFile("", "dbkit-test")
	if err != nil {
		return nil, err
	}

	if f.f, err = Open(f0); err != nil {
		return nil, err
	}

	f.name = f0.Name()
	return f.f, nil
}

func (f *testFile) free() error {
	if err := f.f.Close(); err != nil {
		return err
	}

	return os.Remove(f.name)
}

type testOSFile struct {
	f *osFile
}

func (f *testOSFile) new() (Interface, error) {
	f0, err := ioutil.TempFile("", "dbkit-test")
	if err != nil {
		return nil, err
	}

	f.f = &osFile{f0}
	return f.f, nil
}

func (f *testOSFile) free() error {
	nm := f.f.Name()
	if err := f.f.Close(); err != nil {
		return err
	}

	return os.Remove(nm)
}

type testMem struct {
	mem Interface
}

func (m *testMem) new() (Interface, error) {
	mem, err := OpenMem("")
	if err != nil {
		return nil, err
	}

	m.mem = mem
	return mem, nil
}

func (m *testMem) free() error {
	if mem := m.mem; mem != nil {
		mem.Close()
	}
	return nil
}

func TestTuncate(t *testing.T) {
	testTruncate(t, &testFile{}, 1)
	testTruncate(t, &testMem{}, 1)
	testTruncate(t, &testOSFile{}, 1)
}

func testTruncate0(t testing.TB, f Interface) {
	// Check Truncate works.
	sz := int64(1e6)
	if err := f.Truncate(sz); err != nil {
		t.Error(err)
		return
	}

	fi, err := f.Stat()
	if err != nil {
		t.Error(err)
		return
	}

	if g, e := fi.Size(), sz; g != e {
		t.Error(g, e)
		return
	}

	sz *= 2
	if err := f.Truncate(sz); err != nil {
		t.Error(err)
		return
	}

	if fi, err = f.Stat(); err != nil {
		t.Error(err)
		return
	}

	if g, e := fi.Size(), sz; g != e {
		t.Error(g, e)
		return
	}

	sz = 0
	if err := f.Truncate(sz); err != nil {
		t.Error(err)
		return
	}

	if fi, err = f.Stat(); err != nil {
		t.Error(err)
		return
	}

	if g, e := fi.Size(), sz; g != e {
		t.Error(g, e)
		return
	}

	// Check Truncate(-1) doesn't work.
	sz = -1
	if err := f.Truncate(sz); err == nil {
		t.Error(err)
	}
}

func testTruncate(t testing.TB, ti testInterface, n int) {
	f, err := ti.new()
	if err != nil {
		t.Error(err)
		return
	}

	defer func() {
		if err := ti.free(); err != nil {
			t.Error(err)
		}
	}()

	for i := 0; i < n; i++ {
		testTruncate0(t, f)
	}
}

func BenchmarkTruncateFile(b *testing.B)   { testTruncate(b, &testFile{}, b.N) }
func BenchmarkTruncateMem(b *testing.B)    { testTruncate(b, &testMem{}, b.N) }
func BenchmarkTruncateOSFile(b *testing.B) { testTruncate(b, &testOSFile{}, b.N) }

func TestReadWrite(t *testing.T) {
	testReadWrite(t, &testFile{}, 1)
	testReadWrite(t, &testMem{}, 1)
	testReadWrite(t, &testOSFile{}, 1)
}

func testReadWrite0(t testing.TB, f Interface) {
	const (
		N = 1 << 21
		M = 10
	)

	ps := buffer.Get(N)
	defer buffer.Put(ps)
	s := *ps

	pe := buffer.CGet(N)
	defer buffer.Put(pe)
	e := *pe

	rnd := rand.New(rand.NewSource(42))
	for i := range e {
		s[i] = byte(rnd.Intn(256))
	}
	n2 := 0
	for i := 0; i < M; i++ {
		var from, to int
		for from == to {
			from = rnd.Intn(N)
			to = rnd.Intn(N)
		}
		if from > to {
			from, to = to, from
		}
		for i := range s[from:to] {
			s[from+i] = byte(rnd.Intn(256))
		}
		copy(e[from:to], s[from:to])
		if to > n2 {
			n2 = to
		}
		n, err := f.WriteAt(s[from:to], int64(from))
		if err != nil {
			t.Error(err)
			return
		}

		if g, e := n, to-from; g != e {
			t.Error(g, e)
			return
		}
	}

	fi, err := f.Stat()
	if err != nil {
		t.Error(err)
		return
	}

	if g, e := fi.Size(), int64(n2); g != e {
		t.Error(g, e)
		return
	}

	pb := buffer.Get(n2)
	defer buffer.Put(pb)
	b := *pb

	for i := 0; i <= M; i++ {
		from := rnd.Intn(n2)
		to := rnd.Intn(n2)
		if from > to {
			from, to = to, from
		}
		if i == M {
			from, to = 0, n2
		}
		n, err := f.ReadAt(b[from:to], int64(from))
		if err != nil && (!fileutil.IsEOF(err) && n != 0) {
			fi, err = f.Stat()
			if err != nil {
				t.Error(err)
				return
			}

			t.Error(fi.Size(), from, to, err)
			return
		}

		if g, e := n, to-from; g != e {
			t.Error(g, e)
			return
		}

		if g, e := b[from:to], e[from:to]; !bytes.Equal(g, e) {
			t.Errorf(
				"i %d from %d to %d len(g) %d len(e) %d\n---- got ----\n%s\n---- exp ----\n%s",
				i, from, to, len(g), len(e), hex.Dump(g), hex.Dump(e),
			)
			return
		}
	}

	pbuf := buffer.Get(n2)
	defer buffer.Put(pbuf)
	buf := bytes.NewBuffer(*pbuf)
	buf.Reset()
	if _, err := f.WriteTo(buf); err != nil {
		t.Error(err)
		return
	}

	if g, e := buf.Bytes(), e[:n2]; !bytes.Equal(g, e) {
		t.Errorf("\nlen %d\n%s\nlen %d\n%s", len(g), hex.Dump(g), len(e), hex.Dump(e))
		return
	}

	if err := f.Truncate(0); err != nil {
		t.Error(err)
		return
	}

	if _, err := f.ReadFrom(buf); err != nil {
		t.Error(err)
		return
	}

	p := buffer.Get(n2)
	defer buffer.Put(p)
	roundTrip := *p

	if n, err := f.ReadAt(roundTrip, 0); err != nil && n == 0 {
		t.Error(err)
		return
	}

	if g, e := roundTrip, e[:n2]; !bytes.Equal(g, e) {
		t.Errorf("\nlen %d\n%s\nlen %d\n%s", len(g), hex.Dump(g), len(e), hex.Dump(e))
		return
	}
}

func testReadWrite(t testing.TB, ti testInterface, n int) {
	f, err := ti.new()
	if err != nil {
		t.Error(err)
		return
	}

	defer func() {
		if err := ti.free(); err != nil {
			t.Error(err)
		}
	}()

	for i := 0; i < n; i++ {
		testReadWrite0(t, f)
	}
}

func BenchmarkReadWriteFile(b *testing.B)   { testReadWrite(b, &testFile{}, b.N) }
func BenchmarkReadWriteMem(b *testing.B)    { testReadWrite(b, &testMem{}, b.N) }
func BenchmarkReadWriteOSFile(b *testing.B) { testReadWrite(b, &testOSFile{}, b.N) }

func benchmarkWriteAt(b *testing.B, ti testInterface) {
	f, err := ti.new()
	if err != nil {
		b.Error(err)
		return
	}

	defer func() {
		if err := ti.free(); err != nil {
			b.Error(err)
		}
	}()

	p := buffer.Get(testSize)
	defer buffer.Put(p)
	rand.Read(*p)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := f.WriteAt(*p, 0); err != nil {
			b.Error(err)
			return
		}
	}
	b.SetBytes(testSize)
}

func BenchmarkWriteAtFile(b *testing.B)   { benchmarkWriteAt(b, &testFile{}) }
func BenchmarkWriteAtMem(b *testing.B)    { benchmarkWriteAt(b, &testMem{}) }
func BenchmarkWriteAtOSFile(b *testing.B) { benchmarkWriteAt(b, &testOSFile{}) }

func benchmarkReadAt(b *testing.B, ti testInterface) {
	f, err := ti.new()
	if err != nil {
		b.Error(err)
		return
	}

	defer func() {
		if err := ti.free(); err != nil {
			b.Error(err)
		}
	}()

	p := buffer.Get(testSize)
	defer buffer.Put(p)
	rand.Read(*p)
	if _, err := f.WriteAt(*p, 0); err != nil {
		b.Error(err)
		return
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := f.ReadAt(*p, 0); err != nil && !fileutil.IsEOF(err) {
			b.Error(err)
			return
		}
	}
	b.SetBytes(testSize)
}

func BenchmarkReadAtFile(b *testing.B)   { benchmarkReadAt(b, &testFile{}) }
func BenchmarkReadAtMem(b *testing.B)    { benchmarkReadAt(b, &testMem{}) }
func BenchmarkReadAtOSFile(b *testing.B) { benchmarkReadAt(b, &testOSFile{}) }

func benchmarkWriteAtOne(b *testing.B, ti testInterface) {
	const size = 1 << 17

	f, err := ti.new()
	if err != nil {
		b.Error(err)
		return
	}

	defer func() {
		if err := ti.free(); err != nil {
			b.Error(err)
		}
	}()

	p := buffer.Get(size)
	defer buffer.Put(p)
	rb := *p
	rand.Read(rb)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for off := 0; off < size-1; off++ {
			if _, err := f.WriteAt(rb[off:off+1], int64(off)); err != nil {
				b.Error(err)
				return
			}
		}
	}
	b.SetBytes(size)
}

func BenchmarkWriteAtOneFile(b *testing.B)   { benchmarkWriteAtOne(b, &testFile{}) }
func BenchmarkWriteAtOneMem(b *testing.B)    { benchmarkWriteAtOne(b, &testMem{}) }
func BenchmarkWriteAtOneOSFile(b *testing.B) { benchmarkWriteAtOne(b, &testOSFile{}) }

func benchmarkReadAtOne(b *testing.B, ti testInterface) {
	const size = 1 << 18

	f, err := ti.new()
	if err != nil {
		b.Error(err)
		return
	}

	defer func() {
		if err := ti.free(); err != nil {
			b.Error(err)
		}
	}()

	p := buffer.Get(size)
	defer buffer.Put(p)
	rb := *p
	rand.Read(rb)
	if _, err := f.WriteAt(rb, 0); err != nil {
		b.Error(err)
		return
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for off := 0; off < size-1; off++ {
			if _, err := f.ReadAt(rb[off:off+1], int64(off)); err != nil && !fileutil.IsEOF(err) {
				b.Error(err)
				return
			}
		}
	}
	b.SetBytes(size)
}

func BenchmarkReadAtOneFile(b *testing.B)   { benchmarkReadAtOne(b, &testFile{}) }
func BenchmarkReadAtOneMem(b *testing.B)    { benchmarkReadAtOne(b, &testMem{}) }
func BenchmarkReadAtOneOSFile(b *testing.B) { benchmarkReadAtOne(b, &testOSFile{}) }
