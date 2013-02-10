package terms

type constructor struct {
	name string
}
func (this *constructor) String() string {
	return this.name
}

var constructors map[string]*constructor

func NewConstructor(name string) *constructor {
	if constructors == nil {
		constructors = make(map[string]*constructor)
	}
	if c, ok := constructors[name]; ok {
		return c
	}
	c := constructor{name}
	constructors[name] = &c
	return &c
}