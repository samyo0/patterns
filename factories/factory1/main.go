package main

func main() {

}

/*
Object creation logic becomes too convoluted
Struct has too many fields, need to initialize all correctly

Wholesale object create (non-pieecewise, unlike Builder,
which is peicewise building up a object step by step) can be
outsourced to.
	- A separate function (Factory Function, aka Constructor)
	- That may exist in a separate struct (Factory)
Mainly creating an object in one single invocation
of the factory function.

A component responsible solely for the wholesale (not peicewise)
creation of objects.
*/
