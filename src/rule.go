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

func (r rule) Apply(t term) (term, bool) {
	if subst, ok := Match(r.lhs, t); ok {
		return r.rhs.Copy().ApplySubstitution(subst)
	}
	return nil, false
}

func (r rule) ApplyAnywhere(t term) (term, bool) {
	transform := func(t term) (term, bool) {
		if subst, ok := Match(r.lhs, t); ok {
			fmt.Printf("Found a match with subst: %s\n", subst)
			a, b := r.rhs.Copy().ApplySubstitution(subst)
			fmt.Printf("Transform returning %s, %v\n", a, b)
			return a, b
		}
		fmt.Printf("Transform returning %s, %v\n", nil, false)
		return nil, false
	}
	a, b := t.TransformOnceRecursively(transform)
	fmt.Printf("ApplyAnywhere returning %s, %v\n", a, b)
	return a, b
}


// func (this rule) match_aux(t term, c constraints) (subst substitution, matches bool) {
// 	lhs := this.lhs
// 	// rhs := this.rhs

// 	return c.AddConstraint(lhs, t).BuildSubstitution()
// }