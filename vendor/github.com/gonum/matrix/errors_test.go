// Copyright ©2013 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matrix

import "testing"

func leaksPanic(fn func()) (panicked bool) {
	defer func() {
		r := recover()
		panicked = r != nil
	}()
	Maybe(fn)
	return
}

func TestMaybe(t *testing.T) {
	for i, test := range []struct {
		fn     func()
		panics bool
		errors bool
	}{
		{
			fn:     func() {},
			panics: false,
			errors: false,
		},
		{
			fn:     func() { panic("panic") },
			panics: true,
			errors: false,
		},
		{
			fn:     func() { panic(Error{"panic"}) },
			panics: false,
			errors: true,
		},
	} {
		panicked := leaksPanic(test.fn)
		if panicked != test.panics {
			t.Errorf("unexpected panic state for test %d: got: panicked=%t want: panicked=%t",
				i, panicked, test.panics)
		}
		if test.errors {
			err := Maybe(test.fn)
			stack, ok := err.(ErrorStack)
			if !ok {
				t.Errorf("unexpected error type: got:%T want:%T", stack, ErrorStack{})
			}
			if stack.StackTrace == "" {
				t.Error("expected non-empty stack trace")
			}
		}
	}
}
