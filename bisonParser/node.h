#include <iostream>
#include <vector>
#include <string>

using namespace std;

static string wrapWith(string s, const char* name) {
    // cout << "Wrapping " << s << " with " << name << endl;
    string retval = string("<");
    retval.append(name);
    retval.append(">");
    retval.append(s);
    retval.append("</");
    retval.append(name);
    retval.append(">\n");
    return retval;
}

class Term {
public:
    // virtual ~Term() {}
    // virtual Term() {}
    virtual string AsXML() = 0;
    virtual string AsString() = 0;
};

class Rule {
private:
    Term* _lhs;
    Term* _rhs;
public:    
    Rule(Term* lhs, Term* rhs) {
        // cout << "Rule LHS " << lhs->AsString() << endl;
        // cout << "Rule RHS " << rhs->AsString() << endl;
        _lhs = lhs;
        _rhs = rhs;
    }
    string AsXML() {
        // cout << "Visiting Rule " << endl;
        string lhs = wrapWith(_lhs->AsXML(), "LHS");
        string rhs = wrapWith(_rhs->AsXML(), "RHS");
        return wrapWith(lhs.append(rhs), "Rule");
    }
    string AsString() {
        return _lhs->AsString().append(" => ").append(_rhs->AsString()).append("\n");
    }
};

class RuleSet {
private:
    vector<Rule>* _rules;
public:
    RuleSet(vector<Rule>* rules) {
        _rules = rules;
    }
    string AsXML() {
        string retval = string("<?xml version=\"1.0\"?>\n");

        string rules = string("");
        for(vector<Rule>::iterator it = _rules->begin(); it != _rules->end(); ++it) {
            rules.append((*it).AsXML());
        }

        retval.append(wrapWith(rules, "Rules"));
        return retval;
    }
    string AsString() {
        string rules = string("");
        for(vector<Rule>::iterator it = _rules->begin(); it != _rules->end(); ++it) {
            rules.append((*it).AsString());
        }

        return rules;
    }
};

class Function : public Term {
private:
    string _name;
    vector<Term*>* _children;
public:
    // Function(const char* name) {
    //     _name = string(name);
    //     _children = new vector<Term>();
    // }
    Function(string name, vector<Term*>* children = new vector<Term*>()) {
        _name = name;
        _children = children;
    }
    string AsXML() {
        // cout << "Visiting: " << _name << endl;
        string children = string("");
        for (vector<Term*>::iterator it = _children->begin(); it != _children->end(); ++it) {
            string childXML = (*it)->AsXML();
            // cout << childXML << endl;
            children.append(wrapWith(childXML, "Child"));
        }
        string childrenNode = wrapWith(children, "Children");
        string nameNode = wrapWith(_name, "Name");
        return wrapWith(nameNode.append(childrenNode), "Function");
    }
    string AsString() {
        string children = string("");
        for(vector<Term*>::iterator it = _children->begin(); it != _children->end(); ++it) {
            children.append((*it)->AsString());
        }
        return string(_name).append("(").append(children).append(")");
    }
};

class Variable : public Term {
private:
    string _name;
public:
    Variable(string name) {
        // cout << "Building " << name << endl;
        _name = name;
    }
    string AsXML() {
        // cout << "Visiting: " << _name << endl;
        return wrapWith(wrapWith(_name, "Name"), "Variable");
    }
    string AsString() {
        return string(_name);
    }
};