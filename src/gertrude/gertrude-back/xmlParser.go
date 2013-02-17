package main

import "os"
import "fmt"
import "encoding/xml"

type Query struct {
 	XMLName xml.Name `xml:"Gertrude"`
	Rules []XMLRule `xml:"Rules>Rule"`
	Input XMLTerm `xml:"Input"`
}
func (q Query) getSystem() System {
	rules := make([]Rule, 0)
	for _, rule := range q.Rules {
		rules = append(rules, rule.getRule())
	}
	return NewSystem(rules)
}
func (q Query) getInput() Term {
	return q.Input.getTerm()
}
type XMLRule struct {
	LHS XMLTerm
	RHS XMLTerm
}
func (r XMLRule) getRule() Rule {
	return NewRule(r.LHS.getTerm(), r.RHS.getTerm())
}
type XMLTerm struct {
	Variable XMLVariable
	Function XMLFunction
}
func (t XMLTerm) getTerm() Term {
	if t.Variable.Name != "" {
		return t.Variable.getVariable()
	}
	if t.Function.Name != "" {
		return t.Function.getFunction()
	}
	panic("This should never happen\n")
}
type XMLFunction struct {
	Name string
	Children []XMLTerm `xml:"Children>Child"`
}
func (t *XMLFunction) getFunction() Term {
	// fmt.Printf("function %s with %d children\n", t.Name, len(t.Children))
	children := make([]Term, 0)
	for _, child := range t.Children {
		//_ = child
		children = append(children, child.getTerm())
	}
	return NewFunction(t.Name, children)
}
type XMLVariable struct {
	Name string
}
func (t *XMLVariable) getVariable() Term {
	// fmt.Printf("variable %s\n", t.Name)
	return NewVariable(t.Name)
}
 
func Parse() (System, Term, bool) {
	// xmlFile, err := os.Open("example.xml")
	// if err != nil {
	// 	fmt.Println("Error opening file:", err)
	// 	return nil, false
	// }
	// defer xmlFile.Close()
	_ = fmt.Sprintf("")
	var q Query
	d := xml.NewDecoder(os.Stdin)
	err := d.Decode(&q)
	if err != nil {
		fmt.Printf("Error decoding xml: %s\n", err)
		os.Exit(1)
	}

	// fmt.Printf("%+v\n", q.Input)
	// xml.Unmarshal(xmlFile, &q)
	// out := "----------------------------\n"
	// for _, rule := range q.Rules {
	// 	out += fmt.Sprintf("%+v\n", rule)
	// 	out += "----------------------------\n"
	// }
	// fmt.Printf("%s\n", out)
	sys := q.getSystem()
	fmt.Printf("System:\n%s\n", sys)
	input := q.getInput()
	fmt.Printf("Input:\n%s\n", input)
	fmt.Printf("----------------------------\n")
	return sys, input, true
	// for _, episode := range q.EpisodeList {
	// 	fmt.Printf("\t%s\n", episode)
	// }
}