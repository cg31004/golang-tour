package main

import (
	"fmt"
	"time"
)

func main() {
	a := 0
	times := 10000 // <-- HERE
	start := time.Now()
	c := make(chan bool)

	for i := 0; i < times; i++ {
		go func() {
			a++
			c <- true
		}()
	}

	for i := 0; i < times; i++ {
		<-c
	}
	end := time.Since(start)
	fmt.Printf("a = %d\n", a)
	fmt.Println(end)
}
