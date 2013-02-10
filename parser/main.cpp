#include <iostream>
#include "node.h"
extern Term* program;
extern int yyparse();

int main(int argc, char **argv) {
    yyparse();
    // std::cout << program << std::endl;
    return 0;
}