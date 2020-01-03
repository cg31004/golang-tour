package main

import (
	"fmt"
	"math/big"
	"time"
)

func main() {
	start := time.Now()
	r := new(big.Int)
	fmt.Println(r.Binomial(1000, 10))
	time.Sleep(100 * time.Millisecond)
	elaspsed := time.Since(start)
	fmt.Printf("Binomial took %s", elaspsed)
}
