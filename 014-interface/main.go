package main

import (
	"fmt"
	"math"
)

func main() {
	interface1()
}

//===================    interface1    ==========================
type Abser interface {
	Abs() float64
}

type Vertex1 struct {
	X, Y float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex1) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// func (v Vertex1) Abs() float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

func interface1() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex1{3, 4}

	a = f
	a = &v

	// a = v
	fmt.Println(a.Abs())
}
