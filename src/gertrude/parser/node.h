#include <iostream>
#include <vector>
#include <string>

using namespace std;

class Term {
public:
    // virtual ~Term() {}
    // virtual Term() {}
};

class Rule {
private:
    Term* _lhs;
    Term* _rhs;
public:    
    Rule(Term* lhs, Term* rhs) {
        _lhs = lhs;
        _rhs = rhs;
    }
    std::string AsXML() {
        return std::string("a rule\n");
    }
};

static std::string wrapWith(std::string s, const char* name) {
    std::string retval = std::string("<");
    retval.append(name);
    retval.append(">");
    retval.append(s);
    retval.append("</");
    retval.append(name);
    retval.append(">\n");
    return retval;
}

class RuleSet {
private:
    std::vector<Rule>* _rules;
public:    
    RuleSet(std::vector<Rule>* rules) {
        _rules = rules;
    }
    std::string AsXML() {
        std::string retval = std::string("<?xml version=\"1.0\"?>\n");

        std::string rules = std::string("");
        for(std::vector<Rule>::iterator it = _rules->begin(); it != _rules->end(); ++it) {
            rules.append((*it).AsXML());
        }

        retval.append(wrapWith(rules, "Rules"));
        return retval;
    }
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
    string AsXML() {
        return _name;
    }
};