package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxEnvelopes([][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}))
}

func maxEnvelopes(envelopes [][]int) int {
	count := len(envelopes)

	// 快速排序
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})

	// 取出h
	height := make([]int, count)
	for i := 0; i < count; i++ {
		height[i] = envelopes[i][1]
	}

	// 计算LIS
	return lengthOfLIS(height)
}

func lengthOfLIS(nums []int) int {
	piles, count := 0, len(nums)
	top := make([]int, count)

	for i := 0; i < count; i++ {
		num := nums[i]
		left, right := 0, piles

		for left < right {
			mid := (left + right) / 2
			if top[mid] >= num {
				right = mid
			} else {
				left = mid + 1
			}
		}
		if left == piles {
			piles++
		}
		top[left] = num
	}
	return piles
}
