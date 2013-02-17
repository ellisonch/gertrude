package main

// import "os"

type constraint struct {
	lhs *variable
	rhs Term
}

type constraints []constraint

func (this constraints) String() string {
	retval := "{"
	for _, constraint := range this {
		retval += "  " + constraint.lhs.String() + " = " + constraint.rhs.String() + ", "
	}
	retval += "}"
	return retval
}

func NewConstraints() constraints {
	return make(constraints, 0)
}

func (this constraints) AddConstraint(t1 *variable, t2 Term) constraints {
	this = append(this, constraint{t1, t2})
	return this
}
// func (this constraints) AddConstraints(t1s []Term, t2s []Term) constraints {
// 	if len(t1s) != len(t2s) { os.Exit(1) }
// 	for i, t1 := range t1s {
// 		this = append(this, constraint{t1, t2s[i]})
// 	}
// 	return this
// }

func (this constraints) BuildSubstitution() (subst substitution, matches bool) {
	// log.Printf("Trying to build substitution from constraints: %s\n", this)

	s := NewSubstitution()

	for _, c := range this {
		lhs := c.lhs
		rhs := c.rhs
		if other, ok := s[lhs.name]; ok {
			if !other.Equals(rhs) {
				return nil, false
			}
		} else {
			s[lhs.name] = rhs
		}
	}

	// log.Printf("Done.  Subst: %s\n", s)

	return s, true
}

// func (this constraints) BuildSubstitution_aux(s substitution) (subst substitution, matches bool) {
// 	fmt.Printf("Trying to build substitution from constraints: %s.  Already decided on %s", this, s)
// 	if len(this) == 0 {
// 		return s, true
// 	}
	
// 	// first := this[0]
// 	// others := this[1:]
// 	// lhs := first.lhs
// 	// rhs := first.rhs

// 	// if rhs.ContainsVariable(lhs) {
// 	// 	return nil, false
// 	// }

// 	// if s, ok := lhs.BuildSubstitution(rhs, s); ok {
// 	// 	return constraints.BuildSubstitution_aux(s)
// 	// } else {
// 	// 	return nil, false
// 	// }
// 	return nil, false
// }

// func (this constraints) ApplyOne(lhs *variable, rhs Term) constraints {
// 	return this
// }