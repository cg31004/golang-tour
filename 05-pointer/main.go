package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	p := Vertex{1, 2}
	q := &p
	z := p
	q.X = 1e2
	fmt.Println(p)
	fmt.Println(q)

	z.X = 1e3
	fmt.Println(p)
	fmt.Println(z)

	q = new(Vertex)
	q.X = 9
	fmt.Println(p)
	fmt.Println(q)

}
