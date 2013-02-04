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

	add1 := NewRule(xp0, x)
	add2 := NewRule(xpsy, sxpy)

	fmt.Printf("%v\n", add1.String())
	fmt.Printf("%v\n", add2.String())
}