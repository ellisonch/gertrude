package main

import "strings"

type term interface {
	String() string
	IsVariable() bool
	GetConstructor() *constructor // wtfffff
	GetChildren() []term
	AsVariable() *variable
	AsFunction() *function
}

type function struct {
	constructor *constructor
	children []term
}
func (this *function) GetConstructor() *constructor {
	return this.constructor
}
func (this *function) GetChildren() []term {
	return this.children
}
func (this *function) AsVariable() *variable {
	return nil
}
func (this *function) AsFunction() *function {
	return this
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
func (this *variable) GetConstructor() *constructor {
	return nil
}
func (this *variable) GetChildren() []term {
	return nil
}
func (this *variable) AsVariable() *variable {
	return this
}
func (this *variable) AsFunction() *function {
	return nil
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