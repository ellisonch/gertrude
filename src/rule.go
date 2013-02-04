package main

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