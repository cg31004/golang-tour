package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	a := Vertex{1, 2}
	fmt.Println(a.X)
	fmt.Println(a.Y)
	a.X = 4
	fmt.Println(a.X)
	fmt.Println(a.Y)
}
