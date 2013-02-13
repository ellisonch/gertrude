package parser

import "os"
import _ "testing"
import "bufio"
import "fmt"

func ExampleBasic() {
    if file, err := os.Open("sample.grt"); err == nil {
        defer file.Close(); 
        ok := yyParse(NewLexer(bufio.NewReader(file)))
        if ok == 0 {
     	   fmt.Printf("%s\n", program)
    	}
    } else {
        os.Exit(1)
    }
    // Output: hello
}