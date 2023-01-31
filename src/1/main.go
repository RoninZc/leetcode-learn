package main

import (
	"fmt"
)

func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}

func twoSum(nums []int, target int) []int {
	length := len(nums)
	if length < 2 {
		return []int{0, 0}
	}
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{0, 0}
}

// 使用map方式解题，降低时间复杂度
func twoSum1(nums []int, target int) []int {
	m := make(map[int]int)
	for index, val := range nums {
		if preIndex, ok := m[target-val]; ok {
			return []int{preIndex, index}
		} else {
			m[val] = index
		}
	}
	return []int{}
}
