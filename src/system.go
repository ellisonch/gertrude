package main

type system struct {
	rules []rule
}
func (this system) String() string {
	result := ""
	for _, rule := range this.rules {
		result += rule.String() + "\n"
	}
	return result
}