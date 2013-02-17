package main

import "fmt"
import logPackage "log"
// import "encoding/xml"
import "os"
import "time"
import "flag"

var doTrace = flag.Bool("trace", false, "Trace exection to log file")


var log *logPackage.Logger


func main() {
	if sys, input, ok := Parse("sample.grt"); ok {
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