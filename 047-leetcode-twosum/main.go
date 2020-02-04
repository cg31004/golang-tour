package main

import (
	"fmt"
	"sort"
)

func main() {
	a := twoSum([]int{3, 2, 4}, 6)
	fmt.Println(a)
}

func twoSum(nums []int, target int) []int {
	tar := make(map[int]int)
	for i, v := range nums {
		tar[target-v] = i
	}
	for i := 0; i < len(nums); i++ {
		if _, ok := tar[nums[i]]; ok && tar[nums[i]] != i {
			r := []int{tar[nums[i]], i}
			sort.Ints(r)
			return r
		}
	}
	return []int{77777, 11111}
}
