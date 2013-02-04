package main

import "strings"

type term interface {
	String() string
}

type function struct {
	constructor *constructor
	children []term
}
func (f *function) String() string {
	retval := f.constructor.String()
	arity := len(f.children)
	if arity == 0 {
		return retval
	}

	stringChildren := []string{}
	for _, child := range f.children {
		stringChildren = append(stringChildren, child.String())
	}

	retval += "("
	retval += strings.Join(stringChildren, ", ")
	retval += ")"
	return retval
}

type variable struct {
	name string
}
func (f *variable) String() string {
	return f.name
}


// statically check that functions and variables are terms
var _ term = &function{}
var _ term = &variable{}


func NewFunction(constructorName string, children []term) term {
	c := NewConstructor(constructorName)
	var f term = &function{c, children}
	return f
}

func NewVariable(name string) term {
	var x term = &variable{name}
	return x
}