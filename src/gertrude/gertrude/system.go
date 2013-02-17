package main

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

var rewrites uint64

func (this System) Rewrite(t1 Term) (Term, uint64, bool) {
	applications := 1

	dots := ""
	dots += "digraph G {\n"
	dots += "subgraph {\n"
	dots += t1.AsDot()
	dots += "}\n"

	for applications > 0 {
		applications = 0
		for _, rule := range this.rules {
			log.Printf("Trying to apply %s somewhere in term %s\n", rule, t1)
			if t2, ok := rule.ApplyAnywhere(t1); ok {
				log.Printf("Got %s\n", t2)
				t1 = t2
				applications++
				rewrites++
				dots += "subgraph {\n"
				dots += t2.AsDot()
				dots += "}\n"
			} else {
				// log.Printf("Failed\n")
			}
		}
	}

	dots += "}\n"
	dot.WriteString(dots)

	return t1, rewrites-1, true
}