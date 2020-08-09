package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) RemoveEntry(index int) {
	//
}

func (j *Journal) Save(filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(j.String()), 0644) //specify permission
}

func (j *Journal) Load(filename string) {
	//
}

func (j *Journal) LoadFromWeb(url string) {
	//
}

var LineSeparator = "\n"

func SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(strings.Join(j.entries, LineSeparator)), 0644)
}

type Persistence struct {
	LineSeparator string
}

func (p *Persistence) SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(strings.Join(j.entries, p.LineSeparator)), 0644)
}

func main() {
	j := Journal{}
	j.AddEntry("I cried today")
	j.AddEntry("I ate a bug")
	fmt.Println(j.String)

	SaveToFile(&j, "journal.txt")

	p := Persistence{"\r\n"}
	p.SaveToFile(&j, "journal.txt")
}

/*
Single Responsibility
The single responsibility principle states that a type should have one
primary responsibility and as a result have one reason to change. That
reason being somehow related to its primary responsibility.

Separation of Concern
We can break this SRP by adding functions which for example deal with
another concern. There is another term in addition to the SRP and that
is the separation of concerns. It basically means that different concerns
or different problems that the system solves have to reside in different
constructs of whether attached to different structures or residing residing
in different packages. You CANNOT put everything in a single package for
example because that is an anti-pattern and this is called a God Object.

How it can break down
Because the responsibility of the journal is to do management of the
entries, trying to persist data in this file is breaking the SRP. This
file/package is not to deal with persistence, which can be handled by
a separate component whether its a separate package (struct that has
some methods related to persistance)

Not convenient or scalable
There might be different types in system that needs to be persisted

Two Ways to Solve This
1. Create a different function
2. Create another struct handling this concern
*/
