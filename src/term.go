package main

import "strings"
import "fmt"

type transformer func (term) (term, bool)

type term interface {
	String() string
	match_aux(t term, c constraints) (constraints, bool)
	match_with_function(t *function, c constraints) (constraints, bool)
	ContainsVariable(v *variable) bool
	Equals(t term) bool
	EqualsFunction(t *function) bool
	ApplySubstitution(s substitution) (term, bool)
	TransformOnceRecursively(trans transformer) (term, bool)
	Copy() term
	// EqualsVariable(t *variable) bool
}

func Match(t1 term, t2 term) (substitution, bool) {
	if c, ok := t1.match_aux(t2, NewConstraints()); ok {
		return c.BuildSubstitution()
	}
	return nil, false
}

// ----------------------------------------------------------------------------------------

type function struct {
	constructor *constructor
	children []term
}

func NewFunction(constructorName string, children []term) term {
	c := NewConstructor(constructorName)
	var f term = &function{c, children}
	return f
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
func (t1 *function) match_aux(t2 term, c constraints) (constraints, bool) {
	fmt.Printf("Trying to match %s with %s given constraints %s\n", t1, t2, c)
	return t2.match_with_function(t1, c)
}
func (t1 *function) match_with_function(t2 *function, c constraints) (constraints, bool) {
	fmt.Printf("Trying to match %s with %s given constraints %s\n", t2, t1, c)
	if t1.constructor != t2.constructor {
		return nil, false
	}
	if len(t1.children) != len(t2.children) {
		return nil, false
	}
	for i, child := range t1.children {
		if newC, ok := child.match_aux(t2.children[i], c); ok {
			c = newC
		} else {
			return nil, false
		}
	}
	return c, true
}

// need to consider whether we're replacing in place or what
func (t *function) ApplySubstitution(s substitution) (term, bool) {
	newChildren := []term{}
	fmt.Printf("Applying %s to %s\n", s, t)
	for _, child := range t.children {
		if newC, ok := child.ApplySubstitution(s); ok {
			newChildren = append(newChildren, newC)
		} else {
			return nil, false
		}
	}
	t.children = newChildren
	fmt.Printf("ApplySubstitution returning %s\n", t)
	return t, true
}

func (t *function) TransformOnceRecursively(trans transformer) (term, bool) {
	fmt.Printf("Trying top...\n")
	if tNew, ok := trans(t); ok {
		fmt.Printf("TransformOnceRecursively returning %s, %v\n", tNew, ok)
		return tNew, true
	}
	fmt.Printf("Trying children...\n")
	for i, child := range t.children {
		if newC, ok := child.TransformOnceRecursively(trans); ok {
			t.children[i] = newC
			fmt.Printf("TransformOnceRecursively returning %s, %v\n", t, ok)
			return t, true
		}
	}
	fmt.Printf("Giving up...\n")
	fmt.Printf("TransformOnceRecursively returning %s, %v\n", nil, false)
	return nil, false
}


func (this *function) ContainsVariable(v *variable) bool {
	for _, child := range this.children {
		if child.ContainsVariable(v) {
			return true
		}
	}
	return false
}
func (this *function) Equals(t term) bool {
	return t.EqualsFunction(this)
}
func (this *function) EqualsFunction(t *function) bool {
	if this.constructor != t.constructor {
		return false
	}
	if len(this.children) != len(this.children) {
		return false
	}
	for i, child := range this.children {
		if !child.Equals(t.children[i]) {
			return false
		}
	}
	return true
}

func (this *function) Copy() term {
	newChildren := []term{}
	for _, child := range this.children {
		newChildren = append(newChildren, child.Copy())
	}
	return NewFunction(this.constructor.name, newChildren)
}

// ----------------------------------------------------------------------------------------

type variable struct {
	name string
}

func NewVariable(name string) term {
	var x term = &variable{name}
	return x
}

func (this *variable) String() string {
	return this.name
}
func (t1 *variable) match_aux(t2 term, c constraints) (constraints, bool) {
	fmt.Printf("Trying to match %s with %s given constraints %s\n", t1, t2, c)
	return c.AddConstraint(t1, t2), true
}
func (t1 *variable) match_with_function(t2 *function, c constraints) (constraints, bool) {
	fmt.Printf("Trying to match %s with %s given constraints %s\n", t1, t2, c)
	return c.AddConstraint(t1, t2), true
}
func (this *variable) ContainsVariable(v *variable) bool {
	return this.name == v.name // TODO this is dangerous
}
func (this *variable) Equals(t term) bool {
	fmt.Printf("This should never happen")
	return false
}
func (this *variable) EqualsFunction(t *function) bool {
	fmt.Printf("This should never happen")
	return false
}
// need to consider whether we're replacing in place or what
func (t *variable) ApplySubstitution(s substitution) (term, bool) {
	if result, ok := s[t.name]; ok {
		return result, true
	}
	fmt.Printf("This should never happen")
	return nil, false
}
func (t *variable) TransformOnceRecursively(trans transformer) (term, bool) {
	if tNew, ok := trans(t); ok {
		return tNew, true
	}
	return nil, false
}
func (this *variable) Copy() term {
	return NewVariable(this.name)
}

// ----------------------------------------------------------------------------------------
// statically check that functions and variables are terms
var _ term = &function{}
var _ term = &variable{}
