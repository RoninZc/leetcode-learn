package main

import (
	"fmt"
)

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
}

// 暴力
// func sortedSquares(nums []int) []int {
// 	for i := 0; i < len(nums); i++ {
// 		if nums[i] < 0 {
// 			nums[i] = -nums[i]
// 			for j := i + 1; j < len(nums); j++ {
// 				if nums[j] < nums[j-1] {
// 					nums[j], nums[j-1] = nums[j-1], nums[j]
// 				} else {
// 					break
// 				}
// 			}
// 			i--
// 		} else {
// 			nums[i] = nums[i] * nums[i]
// 		}
// 	}
// 	return nums
// }

// 双指针
func sortedSquares(nums []int) []int {
	left, right, dataIndex := 0, len(nums)-1, len(nums)-1
	data := make([]int, len(nums))

	for left <= right {
		if nums[left]*nums[left] > nums[right]*nums[right] {
			data[dataIndex] = nums[left] * nums[left]
			left++
		} else {
			data[dataIndex] = nums[right] * nums[right]
			right--
		}
		dataIndex--
	}

	return data
}
