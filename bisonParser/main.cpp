#include <iostream>
#include "node.h"
extern RuleSet* program;
extern int yyparse();

int main(int argc, char **argv) {
    yyparse();
    std::cout << program->AsXML() << std::endl;
   	std::cout << endl << program->AsString() << std::endl;

    // Variable* v1 = new Variable(string("Q"));
    // vector<Term*>* vec1 = new vector<Term*>();
    // vec1->push_back(v1);
    // Term *t1 = (*vec1)[0];
    // cout << t1->AsString() << endl; // works

    // cout << "---------------" << endl;

    // Variable* v2 = new Variable(string("Q"));
    // vector<Term>* vec2 = new vector<Term>();
    // vec2->push_back(*v2);
    // Term t2 = (*vec2)[0];
    // cout << t2.AsString() << endl; // dies

    // for (vector<Term>::iterator it = vec->begin(); it != vec->end(); ++it) {
    // 	Term x = (*it);
    // 	cout << x.AsString() << endl;
    // }
    // cout << v->AsString() << endl;
    // cout << (*$$).size() << endl;
    // Term v2 = $$->at(0);

    return 0;
}