package main

import "fmt"

func main() {
	twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

func twoSum(nums []int, target int) []int {
	tar := make(map[int]int)
	for i, v := range nums {
		tar[target-v] = i
	}
	for i := 0; i < len(nums); i++ {
		if _, ok := tar[nums[i]]; ok {
			return []int{tar[nums[i]], i}
		}
	}
	return []int{77777, 11111}
}
