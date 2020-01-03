package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("===================    method1    ==========================")
	method1()
	fmt.Println("===================    method2    ==========================")
	method2()
	fmt.Println("===================    method3    ==========================")
	method3()
}

//===================    method1    ==========================
type Vertex1 struct {
	X, Y float64
}

func method1() {
	v := Vertex1{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(v.Even())
}

func (v *Vertex1) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex1) Even() float64 {
	return (v.X + v.Y) / 2
}

//===================    method2    ==========================
type MyFloat float64

func method2() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

//===================    method3    ==========================
type Vertex2 struct {
	X, Y float64
}

func (v *Vertex2) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex2) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func method3() {
	v := Vertex2{3, 4}
	v.Scale(5)
	fmt.Println(v, v.Abs())
}
