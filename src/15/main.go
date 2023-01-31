package main

import "fmt"

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

func threeSum(nums []int) [][]int {
	data := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

	left, right := 1, len(nums)-1
	for i := 0; i < len(nums); i++ {
		left = i + 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				data = append(data, []int{nums[i], nums[left], nums[right]})
				right--
				left++
			} else if sum > 0 {
				right--
			} else {
				left++
			}
		}
	}

	return data
}
