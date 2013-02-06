package main

// import "fmt"

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


// func (this rule) match_aux(t term, c constraints) (subst substitution, matches bool) {
// 	lhs := this.lhs
// 	// rhs := this.rhs

// 	return c.AddConstraint(lhs, t).BuildSubstitution()
// }