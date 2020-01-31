package main

import "fmt"

func main() {
	var count int
	var shutdown bool
	for !shutdown {
		count++
		if count > 10 {
			shutdown = true
		}
		fmt.Println(count)
	}

}
