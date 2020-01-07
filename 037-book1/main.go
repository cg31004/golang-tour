package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//C01P04 命令參數 echo1
	// C01P04()
	//C01P06 命令參數 echo2
	C01P06()
}

//C01P04 命令參數 echo1
func C01P04() {
	var s, sep string
	fmt.Println(os.Args)
	fmt.Println("-----------------")
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

//C01P06 命令參數 echo2
func C01P06() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
		fmt.Println(counts)
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
