package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("hello world!")
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
}

func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	left, right, sum, length := 0, 0, 0, math.MaxInt
	for right < len(nums) {
		sum += nums[right]

		for sum >= target {
			length = min(length, right-left+1)
			sum -= nums[left]
			left++
		}
		right++
	}
	if length == math.MaxInt {
		return 0
	}
	return length
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
