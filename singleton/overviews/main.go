package main

func main() {

}

/*
For some components it only makes sense to have one in the system
- Database repository
- Object factory

The construction call is expensive
Why only do it once
We give everyone the same instance

Want to prevent anyone creating additional copies

Need to take care of lazy instantiation

A component which is instantiate only once
*/
