package main

import (
	"fmt"
)

func main() {
	var c []int
	a := []int{1, 2, 3, 1, 3, 2, 1, 3}
	c = a[0:4]
	c = append(c, a[5:7]...)
	fmt.Println(c)
}
