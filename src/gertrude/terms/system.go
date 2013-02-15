package terms

import logPackage "log"

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

func (this System) Rewrite(t1 Term, aLog *logPackage.Logger) (Term, uint64, bool) {
	log = aLog
	applications := 1
	var rewrites uint64
	for applications > 0 {
		applications = 0
		for _, rule := range this.rules {
			log.Printf("Trying to apply %s somewhere in term %s\n", rule, t1)
			if t2, ok := rule.ApplyAnywhere(t1); ok {
				log.Printf("Got %s\n", t2)
				t1 = t2
				applications++
				rewrites++
			} else {
				log.Printf("Failed\n")
			}
		}
	}
	return t1, rewrites-1, true
}