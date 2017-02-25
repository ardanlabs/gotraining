// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample code to show how the standard library in general,
// does not pass built-in type values with a pointer.

// *******************************************************

// http://golang.org/src/os/env.go
//
// For the caller to use this function, they must pass a copy of their uint8 type
// value into the function. This is the same for the return value. A copy of the
// function’s bool type value is being returned back to the caller.
func isShellSpecialVar(c uint8) bool {
    switch c {
    case '*', '#', '$', '@', '!', '?', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
        return true
    }
    return false
}

// http://golang.org/src/os/env.go
//
// It is best to treat values of type string the same way you treat boolean
// and numeric type values, as a primitive data value.
func getShellName(s string) (string, int) {
    switch {
    case s[0] == '{':
        . . .
        return "", 1 // Bad syntax; just eat the brace.
    case isShellSpecialVar(s[0]):
        return s[0:1], 1
    }
    . . .
    return s[:i], i
}

// http://golang.org/src/strings/strings.go
//
// All of the functions in the strings package accept copies of the caller’s
// string values and return copies of the string values they create.
func Trim(s string, cutset string) string {
    if s == "" || cutset == "" {
        return s
    }
    return TrimFunc(s, makeCutsetFunc(cutset))
}