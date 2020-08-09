package factories

import "fmt"

type Person struct {
	Name string
	Age  int
}

// factory function
//func NewPerson(name string, age int) Person {
//  return Person{name, age}
//}

func NewPerson(name string, age int) *Person {
	return &Person{name, age}
}

func main_() {
	// initialize directly
	p := Person{"John", 22}
	fmt.Println(p)

	// use a constructor
	p2 := NewPerson("Jane", 21)
	p2.Age = 30
	fmt.Println(p2)
}

/*
Most cases when working with go, you can take a struct and full initialize struct
using the curly braces.

There are situation when you want to have some sort of behavior happen as an object
is created.

You want a default value.
Factory function is nothing more than a freestanding function which returns an
instance of the struct you want to create.

If you want simple objects than you can initialize those objects just using curly
brace notation but sometimes you want additional logic added as you create a struct
*/
