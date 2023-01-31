package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
}

func minSubArrayLen(target int, nums []int) int {
	left, sum, min := 0, 0, math.MaxInt
	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		for sum >= target {
			tmpLen := right - left + 1
			if tmpLen < min {
				min = tmpLen
			}
			sum -= nums[left]
			left++
		}
	}
	if min == math.MaxInt {
		return 0
	}
	return min
}
