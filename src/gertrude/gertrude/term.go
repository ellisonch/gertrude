package main

import "strings"
// import "fmt"

type transformer func (Term) (Term, bool)

type Term interface {
	String() string
	match_aux(t Term, c constraints) (constraints, bool)
	match_with_function(t *function, c constraints) (constraints, bool)
	ContainsVariable(v *variable) bool
	Equals(t Term) bool
	EqualsFunction(t *function) bool
	ApplySubstitution(s substitution) (Term, bool)
	TransformOnceRecursively(trans transformer) (Term, bool)
	Copy() Term
	// EqualsVariable(t *variable) bool
}

func Match(t1 Term, t2 Term) (substitution, bool) {
	if c, ok := t1.match_aux(t2, NewConstraints()); ok {
		return c.BuildSubstitution()
	}
	return nil, false
}

// ----------------------------------------------------------------------------------------

type function struct {
	constructor *constructor
	children []Term
}

func NewFunction(constructorName string, children []Term) Term {
	c := NewConstructor(constructorName)
	var f Term = &function{c, children}
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
func (t1 *function) match_aux(t2 Term, c constraints) (constraints, bool) {
	// log.Printf("Trying to match %s with %s given constraints %s\n", t1, t2, c)
	return t2.match_with_function(t1, c)
}
func (t1 *function) match_with_function(t2 *function, c constraints) (constraints, bool) {
	// log.Printf("Trying to match %s with %s given constraints %s\n", t2, t1, c)
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
func (t *function) ApplySubstitution(s substitution) (Term, bool) {
	newChildren := []Term{}
	// log.Printf("Applying %s to %s\n", s, t)
	for _, child := range t.children {
		if newC, ok := child.ApplySubstitution(s); ok {
			newChildren = append(newChildren, newC)
		} else {
			return nil, false
		}
	}
	t.children = newChildren
	// log.Printf("ApplySubstitution returning %s\n", t)
	return t, true
}

func (t *function) TransformOnceRecursively(trans transformer) (Term, bool) {
	// log.Printf("Trying top...\n")
	if tNew, ok := trans(t); ok {
		// log.Printf("TransformOnceRecursively returning %s, %v\n", tNew, ok)
		return tNew, true
	}
	// log.Printf("Trying children...\n")
	for i, child := range t.children {
		if newC, ok := child.TransformOnceRecursively(trans); ok {
			t.children[i] = newC
			// log.Printf("TransformOnceRecursively returning %s, %v\n", t, ok)
			return t, true
		}
	}
	// log.Printf("Giving up...\n")
	// log.Printf("TransformOnceRecursively returning %s, %v\n", nil, false)
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
func (this *function) Equals(t Term) bool {
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

func (this *function) Copy() Term {
	newChildren := []Term{}
	for _, child := range this.children {
		newChildren = append(newChildren, child.Copy())
	}
	return NewFunction(this.constructor.name, newChildren)
}

// ----------------------------------------------------------------------------------------

type variable struct {
	name string
}

func NewVariable(name string) Term {
	var x Term = &variable{name}
	return x
}

func (this *variable) String() string {
	return this.name
}
func (t1 *variable) match_aux(t2 Term, c constraints) (constraints, bool) {
	// log.Printf("Trying to match %s with %s given constraints %s\n", t1, t2, c)
	return c.AddConstraint(t1, t2), true
}
func (t1 *variable) match_with_function(t2 *function, c constraints) (constraints, bool) {
	// log.Printf("Trying to match %s with %s given constraints %s\n", t1, t2, c)
	return c.AddConstraint(t1, t2), true
}
func (this *variable) ContainsVariable(v *variable) bool {
	return this.name == v.name // TODO this is dangerous
}
func (this *variable) Equals(t Term) bool {
	panic("This should never happen\n")
}
func (this *variable) EqualsFunction(t *function) bool {
	panic("This should never happen\n")
}
// need to consider whether we're replacing in place or what
func (t *variable) ApplySubstitution(s substitution) (Term, bool) {
	if result, ok := s[t.name]; ok {
		return result, true
	}
	panic("This should never happen\n")
}
func (t *variable) TransformOnceRecursively(trans transformer) (Term, bool) {
	if tNew, ok := trans(t); ok {
		return tNew, true
	}
	return nil, false
}
func (this *variable) Copy() Term {
	return NewVariable(this.name)
}

// ----------------------------------------------------------------------------------------
// statically check that functions and variables are terms
var _ Term = &function{}
var _ Term = &variable{}
