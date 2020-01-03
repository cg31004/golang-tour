package main

import (
	"fmt"
	"math/cmplx"
)

func Cbrt(x complex128) complex128 {
	temp := cmplx.Pow(x, 3)
	return temp
}

func main() {
	fmt.Println(Cbrt(2))
}
