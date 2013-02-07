package main

import "strings"
import "fmt"

type term interface {
	String() string
	match_aux(t term, c constraints) (constraints, bool)
	match_with_function(t *function, c constraints) (constraints, bool)
	ContainsVariable(v *variable) bool
	Equals(t term) bool
	EqualsFunction(t *function) bool
	Apply(s substitution) (term, bool)
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
	fmt.Printf("Trying to match %s with %s given constraints %s\n", t1, t2, c)
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
func (t *function) Apply(s substitution) (term, bool) {
	newChildren := []term{}
	for _, child := range t.children {
		if newC, ok := child.Apply(s); ok {
			newChildren = append(newChildren, newC)
		} else {
			return nil, false
		}
	}
	t.children = newChildren
	return t, true
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
func (t *variable) Apply(s substitution) (term, bool) {
	if result, ok := s[t.name]; ok {
		return result, true
	}
	fmt.Printf("This should never happen")
	return nil, false
}

// ----------------------------------------------------------------------------------------
// statically check that functions and variables are terms
var _ term = &function{}
var _ term = &variable{}
