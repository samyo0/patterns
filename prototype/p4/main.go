package prototype

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address struct {
	StreetAddress, City, Country string
}

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

func (p *Person) DeepCopy() *Person {
	// note: no error handling below
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	// peek into structure
	fmt.Println(string(b.Bytes()))

	d := gob.NewDecoder(&b)
	result := Person{}
	_ = d.Decode(&result)
	return &result
}

func main() {
	john := Person{"John",
		&Address{"123 London Rd", "London", "UK"},
		[]string{"Chris", "Matt", "Sam"}}

	jane := john.DeepCopy()
	jane.Name = "Jane"
	jane.Address.StreetAddress = "321 Baker St"
	jane.Friends = append(jane.Friends, "Jill")

	fmt.Println(john, john.Address)
	fmt.Println(jane, jane.Address)

}

/*
Is there a way for the struct is copied including all of its dependencies including the following of the pointers
and the rest of it? Can be done with serialization.

When you use a serializer, it will figure out that you have a string and a pointer, which you will have to follow
that pointer to take a look at the data thats actually in there and serialize that data, not just the point of value.
Serializer know how to unwrap a structure and serialize all of its members.

Wnen you serialize an object, you save all of its states including all the dependencies and when you deserialize it
you construct a brand new object initialized with those same value.

The prototype design pattern is all about take a preconfigured object like John, making a copy, and then customizing it.
Keeps the amount of code that you have to write down to a minimum.
*/
