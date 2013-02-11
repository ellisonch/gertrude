package terms

import "fmt"

type Rule struct {
	lhs Term
	rhs Term
}
func (this Rule) String() string {
	return this.lhs.String() + " => " + this.rhs.String()
}

func NewRule(lhs Term, rhs Term) Rule {
	return Rule{lhs, rhs}
}

func (r Rule) Apply(t Term) (Term, bool) {
	if subst, ok := Match(r.lhs, t); ok {
		return r.rhs.Copy().ApplySubstitution(subst)
	}
	return nil, false
}

func (r Rule) ApplyAnywhere(t Term) (Term, bool) {
	transform := func(t Term) (Term, bool) {
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


// func (this Rule) match_aux(t term, c constraints) (subst substitution, matches bool) {
// 	lhs := this.lhs
// 	// rhs := this.rhs

// 	return c.AddConstraint(lhs, t).BuildSubstitution()
// }