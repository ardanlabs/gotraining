// All material is licensed under the Apache License Version 2.0, January 2004
import "errors"

// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/E-Bb5cRuyz

// Sample code to show how the standard library in general,
// does not pass reference types via a pointer unless the function is
// implementing an unmarshal functionality.

// *******************************************************

// http://golang.org/src/net/ip.go
//
// This is a named type from the net package called IP with a base type that
// is a slice of bytes. There is value in using a named type when you need
// to declare behavior around a built-in or reference type.
type IP []byte

// http://golang.org/src/net/ip.go
//
// MarshalText is using a value receiver. This is exactly what I would expect
// to see because we don’t pass reference types with a pointer.
func (ip IP) MarshalText() ([]byte, error) {
	if len(ip) == 0 {
		return []byte(""), nil
	}
	if len(ip) != IPv4len && len(ip) != IPv6len {
		return nil, errors.New("invalid IP address")
	}
	return []byte(ip.String()), nil
}

// http://golang.org/src/net/ip.go
//
// ipEmptyString accepts a value of named type IP. No pointer is used to pass
// this value since the base type for IP is a slice of bytes and therefore a
// reference type.
func ipEmptyString(ip IP) string {
	if len(ip) == 0 {
		return ""
	}
	return ip.String()
}

// http://golang.org/src/net/ip.go
//
// Anytime you are unmarshaling data into a reference type, you will need to
// pass that reference type value with a pointer.
func (ip *IP) UnmarshalText(text []byte) error {
	if len(text) == 0 {
		*ip = nil
		return nil
	}
	s := string(text)
	x := ParseIP(s)
	if x == nil {
		return &ParseError{"IP address", s}
	}
	*ip = x
	return nil
}
