package main

// import "fmt"

type substitution map[string]term

func (s substitution) String() string {
	retval := "{\n"
	for v, t := range s {
		retval += "  " + v + " === " + t.String() + "\n"
	}
	retval += "}\n"
	return retval
}
func (this substitution) AddSubstitution(v *variable, t term) substitution {
	this[v.name] = t
	return this
}

func NewSubstitution() substitution {
	return make(substitution)
}