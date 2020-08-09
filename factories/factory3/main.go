package main

import "fmt"

//define the behaviors that are available on the person
type Person interface {
	SayHello()
}

type person struct {
	name string
	age  int
}

type tiredPerson struct {
	name string
	age  int
}

func (p *person) SayHello() {
	fmt.Printf("Hi, my name is %s, I am %d years old \n", p.name, p.age)
}

func (p *tiredPerson) SayHello() {
	fmt.Printf("Sorry, I'm too tired \n")
}

func NewPerson(name string, age int) Person {
	if age > 100 {
		return &tiredPerson{name, age}
	}
	return &person{name, age}
}

func main() {
	p := NewPerson("James", 34)
	p.SayHello()
}

/*
When you have a factory function you don't always have to return a struct.
Instead what you can do is you can return an interface that the struct
conforms to.

When it come sto making a factory function what you do is you don't return
person struct but rather the Person interface.

What's the difference?
Since you have an interface to work with, you can't modify the underlying
type. We're not exposing the person struct. We're only exposing the Person
interface.
*/
