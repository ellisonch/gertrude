package main

import "strings"

type term interface {
	String() string
	IsVariable() bool
}

type function struct {
	constructor *constructor
	children []term
}
func (this *function) String() string {
	retval := this.constructor.String()
	arity := len(this.children)
	if arity == 0 {
		return retval
	}

	stringChildren := []string{}
	for _, child := range this.children {
		stringChildren = append(stringChildren, child.String())
	}

	retval += "("
	retval += strings.Join(stringChildren, ", ")
	retval += ")"
	return retval
}
func (this *function) IsVariable() bool {
	return false
}

type variable struct {
	name string
}
func (this *variable) String() string {
	return this.name
}
func (this *variable) IsVariable() bool {
	return true
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