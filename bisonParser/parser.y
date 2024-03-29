%{
#include <iostream>
#include <string>
#include <map>
#include <algorithm>
#include "node.h"
using namespace std;
int yylex(); 
int yyerror(const char *p);
extern char linebuf[500];

RuleSet* program;
std::vector<Term*>* input;
%}

//-- SYMBOL SEMANTIC VALUES -----------------------------
%error-verbose
%define parse.lac
%union {
  int val; 
  char sym;
  string* aString;
  Term* term;
  Rule* rule;
  std::vector<Term*>* terms;
  std::vector<Rule>* rules;
  RuleSet* ruleset;
};
%token REWRITE COMMA LEFT_PAREN RIGHT_PAREN SEMICOLON PIPE
%token <aString> VARIABLE FUNCTION

%type <term> term variable function
%type <rule> rule
%type <rules> rules
%type <terms> children terms
%type <ruleset> program

%locations
// %type  <val> exp term sfactor factor res

//-- GRAMMAR RULES ---------------------------------------
%%
program: rules PIPE terms {
  reverse($1->begin(), $1->end());
  $$ = new RuleSet($1); 
  program = $$;
  input = $3;
}

terms: term SEMICOLON terms { 
  $$ = $3;
  $$->push_back($1);
}
| term SEMICOLON { 
  $$ = new vector<Term*>();
  $$->push_back($1);
}

rules: rule rules { 
  $$ = $2; 
  $$->push_back(*$1);
}
| rule { 
  $$ = new vector<Rule>();
  $$->push_back(*$1);
}

rule: term REWRITE term SEMICOLON { 
  // cout << "--------------" << endl;
  $$ = new Rule($1, $3); 
}
// | error SEMICOLON { yyerrok; }

term: LEFT_PAREN term RIGHT_PAREN { $$ = $2; }
| variable { $$ = $1; }
| function { $$ = $1; }

variable: VARIABLE {
  // cout << "Making variable " << *$1 << endl;
  $$ = new Variable(*$1);
  delete $1;
}

function: FUNCTION { 
  // cout << "Making function " << *$1 << endl;
  $$ = new Function(*$1);
  delete $1;
}
| FUNCTION LEFT_PAREN children RIGHT_PAREN {
  // cout << "Making function " << *$1 << endl;
  reverse($3->begin(), $3->end());
  $$ = new Function(*$1, $3);
  delete $1;
}

children: term { 
    //Variable* v = new Variable(string("Q"));
    $$ = new vector<Term*>();
    $$->push_back($1);
    // cout << v->AsString() << endl;
    //cout << (*$$).size() << endl;
    //Term v2 = $$->at(0);
    //cout <<  << endl;
    //cout << (*$$)[0].AsString() << endl;
    /*delete $1;*/
}
| term COMMA children { 
  $$ = $3;
  $$->push_back($1);
  /*delete $1;*/
}

%%

//-- Lexer prototype required by bison, aka getNextToken()
// int yylex(YYSTYPE, YYLTYPE); 
int yyerror(const char *p) { 
  // cerr << "Error: " << p << endl; 
  if(yylloc.first_line) {
    cerr << "ERROR line " << yylloc.first_line << ": " << p << endl << linebuf << endl;
  }
  exit(1);
}

