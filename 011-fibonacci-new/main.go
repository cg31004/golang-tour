package main

import "fmt"

func main() {
	for i := 0; i < 11; i++ {
		fmt.Printf("n=%d, fin=%d\n", i, fin(i))
	}
}

func fin(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	a, b := 0, 1
	temp := 0
	for i := 2; i < x+1; i++ {
		temp = b
		b = a + b
		a = temp
	}
	return b
}
