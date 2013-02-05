package main

import "fmt"

type rule struct {
	lhs term
	rhs term
}
func (this rule) String() string {
	return this.lhs.String() + " => " + this.rhs.String()
}

func NewRule(lhs term, rhs term) rule {
	return rule{lhs, rhs}
}

func (this rule) Match(t term) (subst substitution, matches bool) {
	return this.match_aux(t, NewConstraints())
} 

func (this rule) match_aux(t term, c constraints) (subst substitution, matches bool) {
	lhs := this.lhs
	// rhs := this.rhs

	if lhs.IsVariable() {
		return buildSubstitution(c.AddConstraint(lhs, t))
	}
	return nil, false
}

func buildSubstitution(c constraints) (subst substitution, matches bool) {
	fmt.Printf("%s", c.String())
	
	return nil, false
}