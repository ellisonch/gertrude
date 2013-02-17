package main

import "fmt"
import logPackage "log"
// import "encoding/xml"
import "os"
import "time"
import "flag"

var doTrace = flag.Bool("trace", false, "Trace exection to log file")
var inputFile = flag.String("input", "", "File containing a Gertrude definition")

var log *logPackage.Logger

func main() {
	flag.Parse()
	if *inputFile == "" {
		fmt.Fprintf(os.Stderr, "Must specify an input file!\n")
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
    	flag.PrintDefaults()
		os.Exit(1)
	}

	if sys, input, ok := Parse(*inputFile); ok {
		// fmt.Printf("%s\n", "parsed!")
		// fmt.Printf("%s\n", sys.String())
		logFile, err := os.Create("rewriting.log")
		if err != nil {
			fmt.Println("Error opening file: %s", err)
			return
		}
		defer logFile.Close()
		// l := log.New(logFile, "", log.LstdFlags)
		log = logPackage.New(logFile, "", 0)
		time1 := time.Now()
		t2, rewrites, ok := sys.Rewrite(input)
		time2 := time.Now()
		delta := time2.Sub(time1).Seconds()
		fmt.Printf("%d rewrites; %0.3f rewrites per second\n", rewrites, float64(rewrites)/delta)
		if ok {
			fmt.Printf("%s\n", t2)
		}
	}
}