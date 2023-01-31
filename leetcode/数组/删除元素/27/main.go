package main

import "fmt"

func main() {
	var nums = []int{3, 2, 2, 3}
	fmt.Println(removeElement(nums, 2))
	fmt.Println(nums)

	nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement(nums, 2))
	fmt.Println(nums)
}

// 暴力
// func removeElement(nums []int, val int) int {
// 	size := len(nums)
// 	for i := 0; i < size; i++ {
// 		if nums[i] == val {
// 			for j := i + 1; j < size; j++ {
// 				nums[j-1] = nums[j]
// 			}
// 			size--
// 			i--
// 		}
// 	}
// 	return size
// }

// 双指针
func removeElement(nums []int, val int) int {
	var slowIndex = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[slowIndex] = nums[i]
			slowIndex++
		}
	}
	return slowIndex
}
