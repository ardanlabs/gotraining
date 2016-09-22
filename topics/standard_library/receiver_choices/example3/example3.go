// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

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
// To16 converts the IP address ip to a 16-byte representation.
// If ip is not an IP address (it is the wrong length), To16 returns nil.
func (ip IP) To16() IP {
	if len(ip) == IPv4len {
		return IPv4(ip[0], ip[1], ip[2], ip[3])
	}
	if len(ip) == IPv6len {
		return ip
	}
	return nil
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
