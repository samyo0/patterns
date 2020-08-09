package factories

import "fmt"

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

//predefined employee template
const (
	Developer = iota
	Manager
)

// functional
func NewEmployee(role int) *Employee {
	switch role {
	case Developer:
		return &Employee{"", "Developer", 60000}
	case Manager:
		return &Employee{"", "Manager", 80000}
	default:
		panic("unsupported role")
	}
}

func main() {
	m := NewEmployee(Manager)
	m.Name = "Sam"
	fmt.Println(m)
}

/*
Sometimes you'll find yourself in a situation where you're creating lots of very similar objects
For example, you have several position in a company

Summary
- A factory function (aka constructor) is a helper function for making struct instances
- A factory is any entity that can take care of object creation
- Can be a function or a dedicated struct
*/
