package main

import "context"

// Compiler Error: function type must have no type parameters

func example1() {
	go func[T any](v T) {
		
		// Write some code here.

	}(10)
}

// In a case like example 1, the argument type T is known at the time this
// literal function is being declared since the function is also being called.
// There's no need to make this literal function generic.

func example2() {
	f := func[Result any](ctx context.Context) Result {
		
		// Write some code here.

	}

	result := retry(ctx, time.Second, f)
}

// There may be cases where this is useful, unfortunately it's really hard to
// implement in general.
// * The type of f would have to be a new "generic function" type.
// * When f is called, we don't know which function is stored in it.
// * When we call f with type parameters, since we don't know what function f
//     is, we'd have to devise a mechanism to pass the generic arguments
//     (based on the type parameters) to f as well.