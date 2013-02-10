#include <iostream>
#include "node.h"
extern RuleSet* program;
extern int yyparse();

int main(int argc, char **argv) {
    yyparse();
    std::cout << program->AsXML() << std::endl;
    return 0;
}