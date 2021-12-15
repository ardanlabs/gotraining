package main

// Compiler Error:  methods cannot have type parameters

type value struct {
	s string
}

func (v value) Method[Input any, Result any](i Input) Result {
	
	// Write some code here.
}

// The design draft does not permit methods to declare type parameters that
// are specific to the method. The receiver may have type parameters, but the
// method not add any type parameters.