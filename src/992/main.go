package main

import "fmt"

func main() {
	fmt.Println(subarraysWithKDistinct([]int{1, 2, 1, 2, 3}, 2))
}

func subarraysWithKDistinct(A []int, K int) int {
	return atMostK(A, K) - atMostK(A, K-1)
}

func atMostK(A []int, K int) int {
	nums := make(map[int]int) // 字典 存储窗口内数字出现的次数
	distinct := 0             // 窗口内不同数字的个数
	length := len(A)          // A长度
	left, right := 0, 0       // 左右指针
	res := 0                  // 结果
	for right < length {
		if nums[A[right]] == 0 {
			distinct++
		}
		nums[A[right]]++

		for distinct > K {
			nums[A[left]]--
			if nums[A[left]] == 0 {
				distinct--
			}
			left++
		}
		res += right - left + 1
		right++
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
