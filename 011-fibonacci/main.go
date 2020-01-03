package main

import (
	"fmt"
	"time"
)

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(int) int {

	return func(x int) int {
		if x == 1 || x == 0 {
			return x
		}
		a, b, temp := 0, 1, 0
		for i := 2; i < x+1; i++ {
			temp = b
			b = a + temp
			a = temp
		}
		return b

	}
}

func main() {
	start := time.Now()
	f := fibonacci()
	for i := 0; i < 80; i++ {
		fmt.Println(f(i))
	}
	end := time.Since(start)
	fmt.Println(end)
}
