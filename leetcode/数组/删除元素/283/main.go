package main

import "fmt"

func main() {
	fmt.Println("hello world!")
	fmt.Println(moveZeroes1([]int{0, 1, 0, 3, 12}))
	fmt.Println(moveZeroes2([]int{0, 1, 0, 3, 12}))
}

// 两次遍历
func moveZeroes2(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}

	slowIndex := 0
	for fastIndex := 0; fastIndex < len(nums); fastIndex++ {
		if nums[fastIndex] != 0 {
			nums[slowIndex] = nums[fastIndex]
			slowIndex++
		}
	}
	for i := slowIndex; i < len(nums); i++ {
		nums[i] = 0
	}

	return nums
}

// 一次遍历
func moveZeroes1(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}

	slowIndex := 0
	for fastIndex := 0; fastIndex < len(nums); fastIndex++ {
		if nums[fastIndex] != 0 {
			nums[slowIndex], nums[fastIndex] = nums[fastIndex], nums[slowIndex]
			slowIndex++
		}
	}

	return nums
}
