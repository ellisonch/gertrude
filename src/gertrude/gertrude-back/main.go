package main

import "fmt"
import "log"
// import "encoding/xml"
import "gertrude/terms"
import "os"
import "time"



func main() {
	// zero := terms.NewFunction("0", []terms.Term{})
	// one := terms.NewFunction("s", []terms.Term{zero})
	// two := terms.NewFunction("s", []terms.Term{one})
	// three := terms.NewFunction("s", []terms.Term{two})
	// four := terms.NewFunction("s", []terms.Term{three})

	// _ = two
	// _ = three
	// _ = four

	// onePlusThree := terms.NewFunction("+", []terms.Term{one, three})

	// x := terms.NewVariable("X")
	// y := terms.NewVariable("Y")
	// xp0 := terms.NewFunction("+", []terms.Term{x, zero})
	// sy := terms.NewFunction("s", []terms.Term{y})
	// xpy := terms.NewFunction("+", []terms.Term{x, y})

	// xpsy := terms.NewFunction("+", []terms.Term{x, sy})
	// sxpy := terms.NewFunction("s", []terms.Term{xpy})

	// //allToZero := NewRule(x, zero)

	// add1 := terms.NewRule(xp0, x)
	// add2 := terms.NewRule(xpsy, sxpy)

	// addition := terms.NewSystem([]terms.Rule{add2, add1})

	// // subst, ok := Match(add2.lhs, onePlusThree)
	// // fmt.Printf("%v\n", ok)
	// // fmt.Printf("%s\n", subst.String())

	// // fmt.Printf("-------------\n");
	// // fmt.Printf("\nApplying %s to %s\n", add2, onePlusThree)
	// // t1, ok := add2.Apply(onePlusThree)
	// // fmt.Printf("%s, %v\n", t1, ok)
	// fmt.Printf("-------------\n");
	// t2, ok := addition.Rewrite(onePlusThree)
	// fmt.Printf("Final Answer: %s, %v\n", t2, ok)

	// // fmt.Printf("%s\n", add1.String())
	// // fmt.Printf("%s\n", add2.String())
	// fmt.Printf("%s\n", addition)

	// fmt.Printf("%s\n", terms.Parse(os.Stdin))
	if sys, input, ok := terms.Parse(); ok {
		// fmt.Printf("%s\n", "parsed!")
		// fmt.Printf("%s\n", sys.String())
		logFile, err := os.Create("rewriting.log")
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer logFile.Close()
		l := log.New(logFile, "", log.LstdFlags)
		time1 := time.Now()
		t2, rewrites, ok := sys.Rewrite(input, l)
		time2 := time.Now()
		delta := time2.Sub(time1).Seconds()
		fmt.Printf("%d rewrites; %0.3f rewrites per second\n", rewrites, float64(rewrites)/delta)
		if ok {
			fmt.Printf("%s\n", t2)
		}
	}
}