package main

import (
	"fmt"
	"math"
)

func main() {
	var num float64 = 0
	var num2 float64 = 0
	boo := true
	for num > -10 {
		fmt.Println(math.Abs(num))
		fmt.Println(num)
		fmt.Println("======================")
		num -= 2
	}
	for boo {
		fmt.Println(math.Abs(num2))
		fmt.Println(num2)
		fmt.Println("-----------------")
		num2 -= 2
		if num2 < -10 {
			boo = false
		}
	}

}
