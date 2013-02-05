package main

import "os"

type constraint struct {
	lhs term
	rhs term
}

type constraints []constraint

func (this constraints) String() string {
	retval := "{\n"
	for _, constraint := range this {
		retval += "  " + constraint.lhs.String() + " = " + constraint.rhs.String() + "\n"
	}
	retval += "}\n"
	return retval
}

func NewConstraints() constraints {
	return make(constraints, 0)
}

func (this constraints) AddConstraint(t1 term, t2 term) constraints {
	this = append(this, constraint{t1, t2})
	return this
}
func (this constraints) AddConstraints(t1s []term, t2s []term) constraints {
	if len(t1s) != len(t2s) { os.Exit(1) }
	for i, t1 := range t1s {
		this = append(this, constraint{t1, t2s[i]})
	}
	return this
}

func (this constraints) BuildSubstitution() (subst substitution, matches bool) {
	return this.BuildSubstitution_aux(NewSubstitution())
}

func (this constraints) BuildSubstitution_aux(s substitution) (subst substitution, matches bool) {
	if len(this) == 0 {
		return s, true
	}
	
	constraint := this[0]
	constraints := this[1:]
	lhs := constraint.lhs
	rhs := constraint.rhs
	if rhs.IsVariable() {
		lhs, rhs = rhs, lhs
	}
	if lhs.IsVariable() {
		constraints = constraints.ApplyOne(lhs, rhs)
		s = s.AddSubstitution(lhs.AsVariable(), rhs)
	} else {
		if lhs.GetConstructor() == rhs.GetConstructor() {
			constraints = constraints.AddConstraints(lhs.GetChildren(), rhs.GetChildren())
		} else {
			return nil, false
		}
	}
	
	return constraints.BuildSubstitution_aux(s)
}

// lhs is supposed to be a variable, but can't figure out how to get around type system
func (this constraints) ApplyOne(lhs term, rhs term) constraints {
	return this
}