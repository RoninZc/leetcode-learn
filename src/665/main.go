package main

import "fmt"

func main() {
	fmt.Println(checkPossibility([]int{4, 2, 3}))
	fmt.Println(checkPossibility([]int{4, 2, 1}))
	fmt.Println(checkPossibility([]int{3, 4, 2, 3}))
}
func checkPossibility(nums []int) bool {
	change, length := 1, len(nums)

	for i := 1; i < length; i++ {
		if nums[i-1] > nums[i] {
			change--
			if i == 1 || nums[i] >= nums[i-2] {
				nums[i-1] = nums[i]
			} else {
				nums[i] = nums[i-1]
			}
		}

		if change < 0 {
			return false
		}
	}

	return true
}
