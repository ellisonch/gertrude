%{
#include "node.h"
#include "parser.hpp"
#define YY_USER_ACTION yylloc.first_line = yylloc.last_line = yylineno;
#include <string>
#include <iostream>
using namespace std;

char linebuf[500];
%}

%option warn nodefault
%option noyywrap
%option yylineno

%%
[\n\r]+.*  { strncpy(linebuf, yytext+1, sizeof(linebuf)); /* save the next line */
        yyless(1);      /* give back all but the \n to rescan */
      }
[A-Z][a-zA-Z0-9_]* { yylval.aString = new string(yytext); return VARIABLE; }
[+\-*^/0-9a-z][+\-*^/0-9a-zA-Z]* {
	// cout << "Lexing " << yytext << endl;
	yylval.aString = new string(yytext);
	return FUNCTION; 
}
[,] { return COMMA; }
[\/\/].+ { }
";" { return SEMICOLON; }
"=>" { return REWRITE; }
"(" { return LEFT_PAREN; }
")" { return RIGHT_PAREN; }

<<EOF>>  { return 0; }
[ \t\n\r]+ { }
.        { 
	cerr << "Unrecognized token:\n" << yytext << "\nin:\n" << linebuf << "\n(Line " << yylineno << ")\n" << endl; 
	exit(1);
}
%%