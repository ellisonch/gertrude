%{
#include <iostream>
#include <string>
#include <map>
#include "node.h"
using namespace std;
int yylex(); 
int yyerror(const char *p);
extern char linebuf[500];

RuleSet* program;
%}

//-- SYMBOL SEMANTIC VALUES -----------------------------
%union {
  int val; 
  char sym;
  char* string;
  Term* term;
  Rule* rule;
  std::vector<Term>* children;
  std::vector<Rule>* rules;
  RuleSet* ruleset;
};
%token REWRITE COMMA LEFT_PAREN RIGHT_PAREN SEMICOLON
%token <string> VARIABLE FUNCTION

%type <term> term variable function
%type <rule> rule
%type <rules> run
%type <children> children
%type <ruleset> program

%locations
// %type  <val> exp term sfactor factor res

//-- GRAMMAR RULES ---------------------------------------
%%
program: run { $$ = new RuleSet($1); program = $$; }

run: rule run { $$ = $2; $2->push_back(*$1); }
| rule { $$ = new std::vector<Rule>(); $$->push_back(*$1); }

rule: term REWRITE term SEMICOLON { $$ = new Rule($1, $3); }
// | error SEMICOLON { yyerrok; }

term: LEFT_PAREN term RIGHT_PAREN { $$ = $2; }
| variable
| function

variable: VARIABLE { $$ = new Variable($1); }

function: FUNCTION { $$ = new Function($1); }
| FUNCTION LEFT_PAREN children RIGHT_PAREN { $$ = new Function($1, $3); }

children: term { $$ = new std::vector<Term>(); $$->push_back(*$1); }
| term COMMA children { $$ = $3; $3->push_back(*$1) }

%%

//-- Lexer prototype required by bison, aka getNextToken()
// int yylex(YYSTYPE, YYLTYPE); 
int yyerror(const char *p) { 
  // cerr << "Error: " << p << endl; 
  if(yylloc.first_line) {
    cerr << "ERROR line " <<yylloc.first_line << ": " << p << endl << linebuf << endl;
  }
}

