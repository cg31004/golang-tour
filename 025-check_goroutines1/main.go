package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	// start := time.Now()
	// sum := 0
	// for i := 0; i < 10000000000; i++ {
	// 	sum += i
	// }
	// end := time.Since(start)
	// fmt.Println(sum)
	// fmt.Println(end)
	wg.Add(2)
	start := time.Now()
	s := 0

	var top int
	go func() {
		top = sumtop()
		wg.Done()
	}()

	var bot int
	go func() {
		bot = sumbot()
		wg.Done()
	}()
	wg.Wait()
	s = top + bot

	end := time.Since(start)
	fmt.Println(s)
	fmt.Println(end)

}
func sumtop() int {
	sum := 0
	for i := 0; i < 5000000000; i++ {
		sum += i
	}
	return sum
}

func sumbot() int {
	sum := 0
	for i := 5000000001; i < 10000000001; i++ {
		sum += i
	}
	return sum
}
