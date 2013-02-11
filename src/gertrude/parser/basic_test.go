package parser

import "os"
import _ "testing"
import "bufio"

func ExampleBasic() {
    if file, err := os.Open("sample.grt"); err == nil {
        defer file.Close(); 
        yyParse(newLexer(bufio.NewReader(file)))
    } else {
        os.Exit(1)
    }
    // Output: hello
}