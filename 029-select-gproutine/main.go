package main

import (
	"fmt"
	"time"
)

func fib(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}

	}
}

func main() {
	start := time.Now()
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 90; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fib(c, quit)
	end := time.Since(start)
	fmt.Println(end)
}
