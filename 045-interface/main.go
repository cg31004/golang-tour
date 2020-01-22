package main

import "fmt"

func main() {
	// part1()
	part2()
}

// part1
func part1() {
	var a interface{}
	var b interface{}
	var c interface{}

	a = 5
	b = 1.5
	c = "hello world"

	Switch(a)
	Switch(b)
	Switch(c)
	fmt.Println(a, b, c)
}

func Switch(input interface{}) {
	switch v := input.(type) {
	case float64:
		fmt.Printf("Twice %v is %v \n", v, v*2)
	case int:
		fmt.Printf("Twice %v is %v \n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!! \n", v)
	}
}

// part2
type Member interface {
	GetName() string
	GetAge() int
}

type Robot struct {
	name  string
	age   int
	power int
}

func (r *Robot) Work() {
}

func (r *Robot) GetName() string {
	return r.name
}

func (r *Robot) GetAge() int {
	return r.age
}

type Car struct {
	name      string
	age       int
	odomerter int
}

func (r *Car) Run() {
}

func (r *Car) GetName() string {
	return r.name
}

func (r *Car) GetAge() int {
	return r.age
}

func Cleaning(m Member) {
	fmt.Printf("Consumer Name: %s, Age; %d\n", m.GetName(), m.GetAge())
}

func part2() {
	r := &Robot{
		name:  "GUMDAN",
		age:   1,
		power: 100,
	}
	c := &Car{
		name:      "BENZ",
		age:       2,
		odomerter: 80,
	}
	Cleaning(c)
	Cleaning(r)
}
