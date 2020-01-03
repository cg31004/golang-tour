package main

import (
	"fmt"
	"math"
)

func main() {
	func1()
	fmt.Println("=====================")
	func2()
}

func func1() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(3, 4))
}

func func2() {
	pos, neg := func2_sub(), func2_sub()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func func2_sub() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}

}
