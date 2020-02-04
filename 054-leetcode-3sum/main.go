package main

import (
	"fmt"
	"sort"
)

func main() {
	a := threeSum([]int{-1, 0, 1, 2, -1, -4})
	fmt.Println(a)
}

func threeSum(nums []int) [][]int {
	var result [][]int
	var sliceTemp []int
	var target int
	c := make(chan []int, 100)
	for i, v := range nums {
		sliceTemp = nums[0 : i+1]
		sliceTemp = append(sliceTemp, nums[i:]...)
		target = 0 - v
		go twoSum(sliceTemp, target, c, i)
	}
	for z := range c {
		// result = append(result, i)
		fmt.Println(z)
	}

	return result
}

func twoSum(nums []int, target int, c chan []int, ind int) {
	tar := make(map[int]int)
	for i, v := range nums {
		tar[target-v] = i
	}
	fmt.Printf("ind: %d\t target:  %d\ttar:  %d\n", ind, target, tar)
	for i := 0; i < len(nums); i++ {
		if _, ok := tar[nums[i]]; ok && tar[nums[i]] != i {
			r := []int{tar[nums[i]], i}
			r = append(r, ind)
			sort.Ints(r)
			c <- r
			close(c)
		}
	}
}
