package main

import (
	"fmt"
	"strconv"
)

var aaa []string
var count int
var sss string

func main() {
	// defer defTest()
	fmt.Println(aaa)
	defer fmt.Println(aaa)
	defer defString(count)
	defer defString(count)
	defer defString(count)
	defer fmt.Println(aaa)
}

func defTest() {
	if ok := recover(); ok != nil {
		fmt.Println("recover")
	}
}

func defString(layer int) {
	count++
	sss = strconv.Itoa(layer)
	aaa = append(aaa, sss+sss+sss)
	fmt.Println(strconv.Itoa(layer))
	fmt.Println(count)
}
