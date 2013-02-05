package main

// import "fmt"

type substitution map[variable]term
func (s substitution) String() string {
	retval := "{\n"
	for v, t := range s {
		retval += "  " + v.String() + " === " + t.String() + "\n"
	}
	retval += "}\n"
	return retval
}


type constraint struct {
	lhs term
	rhs term
}

type constraints struct {
	constraints []constraint
	substitution substitution
}
func (cs constraints) String() string {
	retval := "{\n"
	for _, c := range cs.constraints {
		retval += "  " + c.lhs.String() + " = " + c.rhs.String() + "\n"
	}
	retval += "}\n"
	return retval
}

func NewConstraints() constraints {
	return constraints{make([]constraint, 0), make(substitution)}
}

func (cs constraints) AddConstraint(t1 term, t2 term) constraints {
	cs.constraints = append(cs.constraints, constraint{t1, t2})
	return cs
}