package main

import "fmt"

type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct {
	//
}

func (f *Filter) FilterByColor(product []Product, color Color) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

func (f *Filter) FilterBySize(product []Product, size Size) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.size == size {
			result = append(result, &products[i])
		}
	}
	return result
}

type Specification interface {
	IsSatisfied(p *Product) bool
}

//if you want to check for color you would make a color specification
type ColorSpecification struct {
	color Color
}

func (c ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == c.color
}

type SizeSpecification struct {
	size Size
}

func (s SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == s.size
}

type BetterFilter struct{}

func (f *BetterFilter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}
	return result
}

type AndSpecification struct {
	first, second Specification
}

func (a AndSpecification) IsSatisfied(p (*Product) bool {
	return a.first.IsSatisfied(p) && a.second.IsSatisfied(p)
}

func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}
	fmt.Printf("Green products (old): \n")
	f := Filter{}
	for _, v := range f.FilterByColor(products, green) {
		fmt.Printf(" - %s is green \n", v.name)
	}

	fmt.Printf("Green products (new): \n")
	greenSpec := ColorSpecification{green}
	bf := BetterFilter{}
	for _, v := range bf.Filter(products, greenSpec) {
		fmt.Printf("- %s is green \n", v.name)
	}

	largeSpec := SizeSpecification{large}
	lgSpec := AndSpecification{greenSpec, largeSpec}
	fmt.Printf("Large green products: \n")
	for _, v := range bf.Filter(products, lgSpec) {
		fmt.Printf("- %s is large and green \n". v.name)
	}
}

/*
Type should be open for extension but closed for modification.

Enterprise pattern called specification (good example of OCP)

Problem
Initial code only filters by color but later wants to implement
filter by size. You can do this by adding another method to the
Filter struct, which shares a lot of the same code as the other
method. And more changes and updates happen.

Open closed principle is all about beiing open for extensions,
so you want to be able to extend a scenario by adding additional
types and freestanding functions but without modifying something
that you already written and you've already tested.

You want to leave the filter type alone and you don't want to keep
adding more and more methods to it. You want to have an extendable
setup. You can do that with the specification pattern.

Because it has a bunch of interfaces and additional element for
flexibility

The Better filter approach gives you more flexibility. If you want
filter by a particular new type all you have to do is you have
to make a new specification. If you want to filter by size then
all you have to do is make a size specification and make sure
that it conforms to the specification.

The interface type is open for extension, mean you can implement this
interface but its close to modification which means that you're
unlikely to ever modify the specification interface and in a 
similar fashion you are unlikely to ever modify better filter
because there is no reason for us to do so.

composite - combine
*/
