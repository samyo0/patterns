package main

import "fmt"

type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	width, height int
}

func (r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

type Square struct {
	Rectangle
}

func NewSquare(size int) *Square {
	sq := Square{}
	sq.width = size
	sq.height = size
	return &sq
}

func (s *Square) SetWidth(width int) {
	s.width = width
	s.height = width
}

func (s *Square) SetHeight(height int) {
	s.width = height
	s.height = height
}

func UseIt(sized Sized) {
	width := sized.GetWidth()
	sized.SetHeight(10)
	expectedArea := 10 * width
	actualArea := sized.GetWidth() * sized.GetHeight()
	fmt.Printf("Expected area of ", expectedArea, ", but got ", actualArea, "\n")
}

func main() {
	rc := &Rectangle{2, 3}
	UseIt(rc)
}

/*
Liskov Substitution Principle
API takes a base class and works correctly with that base class, then it should
also work correctly with the derived class. In Go, we don't have base classes in
derived classes.

The Issue
For the method UseIt, if it was expect something up the heiracrchy then it should
continue to work even if you proceed to extend object which are already size. So
we took a rectangle and we decided to sort of extend the rectangle and make a
square, it should also continue to work. Everything should continue to work but
it doesnt.

If you use generaliztion like interfaces then you should nothave in heritance or
you should not have implementers of those generalizations break some of the
assumptions which are set up at the higher level.

There is no right answer. You can use different approaches.

The behavior of implementers of a particular type like the Sized interface should
not break the core fundamental behaviors that you rely on.

We've broken certain assumptions about the type and as a result we got incorrect
behavior.
*/
