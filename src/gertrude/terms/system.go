package terms

import "fmt"

type System struct {
	rules []Rule
}

func NewSystem(rules []Rule) System {
	return System{rules}
}

func (this System) String() string {
	result := ""
	for _, rule := range this.rules {
		result += rule.String() + "\n"
	}
	return result
}

func (this System) Rewrite(t1 Term) (Term, bool) {
	applications := 1
	for applications > 0 {
		applications = 0
		for _, rule := range this.rules {
			fmt.Printf("Trying to apply %s somewhere in term %s\n", rule, t1)
			if t2, ok := rule.ApplyAnywhere(t1); ok {
				fmt.Printf("Got %s\n", t2)
				t1 = t2
				applications++
			} else {
				fmt.Printf("Failed\n")
			}
		}
	}
	return t1, true
}