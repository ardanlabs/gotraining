// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample code to show how the standard library in general,
// treats struct types based on the implement of the type.

// *******************************************************
// Structs as primitive data type

// http://golang.org/src/time/time.go
//
// Struct types can be implemented to behave like built-in types. When
// they are, you should treat them as such. 
type Time struct {
    // sec gives the number of seconds elapsed since
    // January 1, year 1 00:00:00 UTC.
    sec int64

    // nsec specifies a non-negative nanosecond
    // offset within the second named by Seconds.
    // It must be in the range [0, 999999999].
    nsec int32

    // loc specifies the Location that should be used to
    // determine the minute, hour, month, day, and year
    // that correspond to this Time.
    // Only the zero Time has a nil Location.
    // In that case it is interpreted to mean UTC.
    loc *Location
}

// http://golang.org/src/time/time.go
//
// If you look at the factory function Now, you will see it returns a value
// of type Time, not a pointer:
func Now() Time {
    sec, nsec := now()
    return Time{sec + unixToInternal, nsec, Local}
}

// http://golang.org/src/time/time.go
//
// Like the built-in types, the Add method is called against a copy of the
// caller’s Time value.
func (t Time) Add(d Duration) Time {
    t.sec += int64(d / 1e9)
    nsec := int32(t.nsec) + int32(d%1e9)
    if nsec >= 1e9 {
        t.sec++
        nsec -= 1e9
    } else if nsec < 0 {
        t.sec--
        nsec += 1e9
    }
    t.nsec = nsec
    return t
}

// http://golang.org/src/time/time.go
//
// Again, values of type Time are treated like a primitive data type and are
// copied when passed.
func div(t Time, d Duration) (qmod2 int, r Duration) {

// *******************************************************
// Structs as a non-primitive data type

// http://golang.org/src/os/file.go
//
// When you have a factory function like Open that is providing you a pointer,
// it is a good sign that you should not be making copies of the referenced
// value being returned.
func Open(name string) (file *File, err error) {
    return OpenFile(name, O_RDONLY, 0)
}

// http://golang.org/src/os/file_unix.go
//
// Open is returning a pointer because it is not safe to make copies of the
// referenced File value being returned. The value should always be used and
// passed through the pointer.

// File represents an open file descriptor.
type File struct {
    *file
}

// file is the real representation of *File.
// The extra level of indirection ensures that no clients of os
// can overwrite this data, which could cause the finalizer
// to close the wrong file descriptor.
type file struct {
    fd int
    name string
    dirinfo *dirInfo // nil unless directory being read
    nepipe int32 // number of consecutive EPIPE in Write
}

// http://golang.org/src/os/file_unix.go
//
// Even if a function or method is not changing the state of a File struct
// type value, it still needs to be passed with a pointer.
func epipecheck(file *File, e error) {
    if e == syscall.EPIPE {
        if atomic.AddInt32(&file.nepipe, 1) >= 10 {
            sigpipe()
        }
    } else {
        atomic.StoreInt32(&file.nepipe, 0)
    }
}

// http://golang.org/src/os/file.go
//
// This applies as well for the methods declared for the File type.
func (f *File) Chdir() error {
    if f == nil {
        return ErrInvalid
    }
    if e := syscall.Fchdir(f.fd); e != nil {
        return &PathError{"chdir", f.name, e}
    }
    return nil
}