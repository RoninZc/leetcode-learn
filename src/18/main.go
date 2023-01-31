package main

import "fmt"

func main() {
	fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8))
}

func fourSum(nums []int, target int) [][]int {
	data := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			left, right := j+1, len(nums)-1
			for left < right {
				if nums[i]+nums[j] > target-nums[left]-nums[right] {
					left++
				} else if nums[i]+nums[j] < target-nums[left]-nums[right] {
					right--
				} else {
					data = append(data, []int{nums[i], nums[j], nums[left], nums[right]})
					for left < right && nums[left] == nums[left+1] { // 去重
						left++
					}
					for left < right && nums[right] == nums[right-1] { // 去重
						right--
					}
					left++
					right--
				}
			}
		}
	}

	return data
}
