#include <iostream>
#include <vector>
#include <string>

using namespace std;

class Term {
public:
    // virtual ~Term() { }
    // virtual Term() { }
};

class Function : public Term {
private:
    std::string _name;
    std::vector<Term>* _children;
public:
    // Function(const char* name) {
    //     _name = std::string(name);
    //     _children = new std::vector<Term>();
    // }
    Function(const char* name, std::vector<Term>* children = new std::vector<Term>()) {
        _name = std::string(name);
        _children = children;
    }
};

class Variable : public Term {
private:
    std::string _name;
public:
    Variable(const char* name) {
        _name = std::string(name);
    }
};