/*

This file is a modified excerpt from the GNU Bison Manual examples originally found here:
http://www.gnu.org/software/bison/manual/html_node/Infix-Calc.html#Infix-Calc

The Copyright License for the GNU Bison Manual can be found in the "fdl-1.3" file.

*/

/* Infix notation calculator. */

%{

package parser

import (
    "bufio"
    "fmt"
    "os"
    "gertrude/terms"
)

var program terms.System

%}

%union{
    value float64
    string string
    term terms.Term
    system terms.System
    rules []terms.Rule
    children []terms.Term
    rule terms.Rule
}

%token NUM SEMICOLON REWRITE LEFT_PAREN RIGHT_PAREN COMMA
%token <string> VARIABLE FUNCTION
%type <system> program
%type <term> variable term function
%type <rule> rule
%type <rules> run
%type <children> children

%% /* The grammar follows. */

program: run { 
    $$ = terms.NewSystem($1)
    program = $$
}
;

run: rule run { 
    $$ = append($2, $1) 
}
| rule { 
    $$ = []terms.Rule{$1}
}
;

rule: term REWRITE term SEMICOLON { 
    $$ = terms.NewRule($1, $3)
    fmt.Printf("Parsed rule as:\n-----\n%s\n-----\n", $$)
}
;

term: LEFT_PAREN term RIGHT_PAREN { $$ = $2 }
| variable
| function
;

variable: VARIABLE { 
    $$ = terms.NewVariable($1)
    fmt.Printf("Parsed variable as:\n-----\n%s\n-----\n", $$)
}
;

function: FUNCTION { $$ = terms.NewFunction($1, []terms.Term{}) }
| FUNCTION LEFT_PAREN children RIGHT_PAREN { $$ = terms.NewFunction($1, $3) }
;

children: term { $$ = []terms.Term{$1} }
| term COMMA children { $$ = append($3, $1) }
;

%%

func main() {
    os.Exit(yyParse(NewLexer(bufio.NewReader(os.Stdin))))
}