package main

import "fmt"

func main() {
	x := NewVariable("X")
	y := NewVariable("Y")
	zero := NewFunction("0", []term{})
	xp0 := NewFunction("+", []term{x, zero})
	sy := NewFunction("s", []term{y})
	xpy := NewFunction("+", []term{x, y})

	xpsy := NewFunction("+", []term{x, sy})
	sxpy := NewFunction("s", []term{xpy})

	allToZero := NewRule(x, zero)

	add1 := NewRule(xp0, x)
	add2 := NewRule(xpsy, sxpy)

	addition := system{[]rule{add1, add2}}

	subst, ok := allToZero.Match(sy)
	fmt.Printf("%v\n", ok)
	fmt.Printf("%s\n", subst.String())

	// fmt.Printf("%s\n", add1.String())
	// fmt.Printf("%s\n", add2.String())
	fmt.Printf("%s\n", addition)
}