package main

import "fmt"

func main() {
	// slice index
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	// slice range
	p := []int{2, 3, 4, 5, 6, 7, 8, 11, 13}
	fmt.Println("p ==", p)

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n", i, p[i])
	}
	fmt.Println("==================")
	for index, element := range p[5:7] {
		fmt.Printf("p[%d] == %d\n", index, element)
	}
	for index, element := range p {
		fmt.Printf("p[%d] == %d\n", index, element)
	}
	p2 := []int{2, 3, 4, 5, 6, 7, 8, 11, 13}
	fmt.Println("p2 ==", p2)
	fmt.Println("p2[1:4] ==", p2[1:4])
	fmt.Println("p2[:3] ==", p2[:3])
	fmt.Println("p2[6:] ==", p2[6:])

	make_a := make([]int, 5)
	printSlice("make_a", make_a)
	make_b := make([]int, 1, 5)
	printSlice("make_b", make_b)

	// 空slice
	var nil_slice []int
	fmt.Println(nil_slice, len(nil_slice), cap(nil_slice))
	if nil_slice == nil {
		fmt.Println("nil!")
	}
	nil_slice = append(nil_slice, 6)
	if nil_slice == nil {
		fmt.Println("nil!")
	} else {
		fmt.Println("append success")
		fmt.Println(nil_slice)
	}

	// for range
	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	// for range ignore index or element
	new_pow := make([]int, 6)
	for i := range new_pow {
		pow[i] = 1 << uint(i)
	}
	for _, v := range new_pow {
		fmt.Printf("%d\n", v)
	}

}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
