#include <iostream>
#include "node.h"
extern RuleSet* program;
extern std::vector<Term*>* input;
extern int yyparse();

int main(int argc, char **argv) {
    yyparse();
    std::cout << program->AsXML(input) << std::endl;
    return 0;
}
