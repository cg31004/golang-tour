package main

import (
	"fmt"
	"time"
)

func sum(a int, c chan int) {
	sum := 0
	for i := a; i < (a + 5000000000); i++ {
		sum += i
	}
	c <- sum
}

func all_sum() int {
	sum := 0
	for i := 0; i < 10000000001; i++ {
		sum += i
	}
	return sum
}

func main() {
	start := time.Now()
	c := make(chan int)
	go sum(1, c)
	go sum(5000000001, c)
	x, y := <-c, <-c
	sum := x + y
	end := time.Since(start)
	fmt.Println(sum)
	fmt.Println(end)
	fmt.Println("===============")
	start = time.Now()
	nsum := all_sum()
	end = time.Since(start)
	fmt.Println(nsum)
	fmt.Println(end)

}
