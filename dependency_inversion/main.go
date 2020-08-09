package main

import "fmt"

//example is modeling relationship between people in family tree

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	name string
	//
}

type Info struct {
	from         *Person
	relationship Relationship
	to           *Person
}

//store data in type
//low level module
type Relationships struct {
	relations []Info
}

func (r *Relationships) AddParentAndChild(parent, child *Person) {
	r.relations = append(r.relations, Info{parent, Parent, child})
	r.relations = append(r.relations, Info{child, Child, parent})
}

func (r *Relationships) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0)
	for i, v := range r.relations {
		if v.relationship == Parent && v.from.name == name {
			result = append(result, r.relations[i].to)
		}
	}
	return result
}

//high level module
// type Research struct {
// 	//break DIP by depending on low level module
// 	relationships *Relationships
// }

type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

//Rewrite
type Research struct {
	browser RelationshipBrowser
}

// func (r *Research) Investigate() {
// 	relations := r.relationships.relations
// 	for _, rel := range relations {
// 		if rel.from.name == "John" && rel.relationship == Parent {
// 			fmt.Println("John has a child called ", rel.to.name)
// 		}
// 	}
// }

//rewrite
func (r *Research) Investigate() {
	for _, p := range r.browser.FindAllChildrenOf("John") {
		fmt.Println("John has a child called", p.name)
	}
}

func main() {
	parent := Person{"John"}
	child1 := Person{"Chris"}
	child2 := Person{"Matt"}

	relationships := Relationships{}
	relationships.AddParentAndChild(&parent, &child1)
	relationships.AddParentAndChild(&parent, &child2)

	r := Research{&relationships}
	r.Investigate()
}

/*
Does not have anything to do with dependency injection

States two things:
- High level modules should not depend on low level modules
- Both of them should depend on abstractions

Issues
If relationships change from slice to database, then the code which depends
on the low level module actually breaks. Everything in the Investigate
method will break (for loop will break).

Rewrite solution
Avoid the situation where we're exposing the internals of relationships and
will define a new interface called relationship browser and will implement
this interface on relationships struct.

FindAllChildrenOf
Now all the finding of the children is actually put into the low level
module and then we can rewrite the high level module, so the high level
module now depends on an abstraction

Low level - closer to hardware sort of data storage communication and system
High level - business log
Abstractions - interfaces
*/
