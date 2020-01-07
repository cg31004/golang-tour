package main

import "fmt"

func main() {
	t1()
}

func t1() {
	defer fmt.Println("Hi end")
	fmt.Println("Hi, Simon")
}
