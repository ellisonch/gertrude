%{
#include <iostream>
#include <string>
#include <map>
#include "node.h"
using namespace std;
int yylex(); 
int yyerror(const char *p);
%}

//-- SYMBOL SEMANTIC VALUES -----------------------------
%union {
  int val; 
  char sym;
  char* string;
  Term* term;
  std::vector<Term>* children;
};
%token REWRITE COMMA LEFT_PAREN RIGHT_PAREN SEMICOLON
%token <string> VARIABLE FUNCTION

%type <term> term variable rule run function
%type <children> children

%locations
// %type  <val> exp term sfactor factor res

//-- GRAMMAR RULES ---------------------------------------
%%
run: rule run | rule    /* forces bison to process many stmts */

rule: term REWRITE term SEMICOLON { cout << "a rule\n" << endl; }
// | error SEMICOLON { yyerrok; }

term: LEFT_PAREN term RIGHT_PAREN { $$ = $2; }
| variable
| function

variable: VARIABLE { $$ = new Variable($1); }

function: FUNCTION { $$ = new Function($1); }
| FUNCTION LEFT_PAREN children RIGHT_PAREN { $$ = new Function($1, $3); }

children: {$$ = new std::vector<Term>()}
| term COMMA children { $$ = $3; $3->push_back(*$1) }

*/

%%

//-- Lexer prototype required by bison, aka getNextToken()
// int yylex(YYSTYPE, YYLTYPE); 
int yyerror(const char *p) { 
  // cerr << "Error: " << p << endl; 
  if(yylloc.first_line) {
    cerr << "ERROR line " <<yylloc.first_line << ": " << p << endl;
  }
}

