package main

// import "fmt"

type substitution map[string]Term

func (s substitution) String() string {
	retval := "{"
	for v, t := range s {
		retval += "  " + v + " === " + t.String() + ", "
	}
	retval += "}"
	return retval
}
func (this substitution) AddSubstitution(v *variable, t Term) substitution {
	this[v.name] = t
	return this
}

func NewSubstitution() substitution {
	return make(substitution)
}