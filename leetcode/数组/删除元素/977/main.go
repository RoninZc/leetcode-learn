package main

import "fmt"

func main() {
	fmt.Println("hello world!")
	fmt.Println(sortedSquares1([]int{-4, -1, 0, 3, 10}))
	fmt.Println(sortedSquares2([]int{-4, -1, 0, 3, 10}))
}

func sortedSquares1(nums []int) []int {
	// 找到最后一个为负数的位置
	negIndex := -1
	for i := 0; i < len(nums) && nums[i] < 0; i++ {
		negIndex = i
	}

	res := make([]int, 0, len(nums))

	// left 向左移动，right 向右移动
	for left, right := negIndex, negIndex+1; left >= 0 || right < len(nums); {
		// 当 left < 0，即该数组后续无负数，直接平方后返回
		if left < 0 {
			res = append(res, nums[right]*nums[right])
			right++
		} else if right == len(nums) {
			// 当 right 超过了数组长度， 相当于右侧无正数，直接平方后返回
			res = append(res, nums[left]*nums[left])
			left--
		} else if nums[left]*nums[left] > nums[right]*nums[right] {
			// 两侧进行比较，确定插入顺序
			res = append(res, nums[right]*nums[right])
			right++
		} else {
			res = append(res, nums[left]*nums[left])
			left--
		}
	}

	return res
}

func sortedSquares2(nums []int) []int {
	res := make([]int, len(nums))

	left, right := 0, len(nums)-1

	for pos := len(nums) - 1; pos >= 0; pos-- {
		if left2, right2 := nums[left]*nums[left], nums[right]*nums[right]; left2 > right2 {
			res[pos] = left2
			left++
		} else {
			res[pos] = right2
			right--
		}
	}

	return res
}
