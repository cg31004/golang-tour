package main

import "fmt"

/*

p["HM"] isn't quite a regular addressable value: hashmaps can grow at runtime, and then their values get moved around in memory, and the old locations become outdated. If values in maps were treated as regular addressable values, those internals of the map implementation would get exposed.

So, instead, p["HM"] is a slightly different thing called a "map index expression" in the spec; if you search the spec for the phrase "index expression" you'll see you can do certain things with them, like read them, assign to them, and use them in increment/decrement expressions (for numeric types). But you can't do everything. They could have chosen to implement more special cases than they did, but I'm guessing they didn't just to keep things simple.


*/

type Person struct {
	name string
	age  int
}

type People map[string]*Person

func main() {
	// p := make(People)
	// p["HM"] = Person{"HANK", 33}
	// p["HM"].age = p["HM"].age + 1
	// fmt.Printf("age : %d\n", p["HM"].age)
	p := make(People)
	p["SS"] = &Person{"Simon Su", 27}
	p["SS"].age += 1
	fmt.Printf("name : %v\tage : %d\n", p["SS"].name, p["SS"].age)
}
