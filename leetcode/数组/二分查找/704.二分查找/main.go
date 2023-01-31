package main

import (
	"fmt"
)

/*
	704. 二分查找
	力扣题目链接

	给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。

	示例 1:

	输入: nums = [-1,0,3,5,9,12], target = 9
	输出: 4
	解释: 9 出现在 nums 中并且下标为 4
	示例 2:

	输入: nums = [-1,0,3,5,9,12], target = 2
	输出: -1
	解释: 2 不存在 nums 中因此返回 -1
	提示：

	你可以假设 nums 中的所有元素是不重复的。
	n 将在 [1, 10000]之间。
	nums 的每个元素都将在 [-9999, 9999]之间。
*/

func main() {
	fmt.Println(binarySearch([]int{-1, 0, 3, 5, 9, 12}, 9))
	fmt.Println(binarySearch([]int{-1, 0, 3, 5, 9, 12}, 0))
	fmt.Println(binarySearch([]int{-1, 0, 3, 5, 9, 12}, 12))
	fmt.Println(binarySearch([]int{-1, 0, 3, 5, 9, 12}, 1000))
	fmt.Println(binarySearch([]int{-1, 0, 3, 5, 9, 12}, -1))
	fmt.Println(binarySearch([]int{-1, 0, 3, 5, 9, 12}, -100))
}

func binarySearch(nums []int, target int) int {
	// 当进入的数组长度为 0 时，直接返回 -1
	if len(nums) == 0 {
		return -1
	}

	// 我们定义目标在左闭右开的区间内
	// 即 [left, right)
	// 当 left == right == 1 时，目标在 [1, 1) 这个区间，此时无意义
	// 所以这里需要判断 left >= right 时，退出循环
	// 如果我们定义目标在左闭右闭的区间内
	// 即 [left, right]
	// 当 left == right == 1 时，目标在 [1, 1] 这个区间，此时有意义
	// 所以这里需要判断 left < right 时，才退出循环

	left := 0
	// 左闭右开
	// right := len(nums)
	// 左开右闭
	right := len(nums) - 1

	// for left < right { // 左闭右开
	for left <= right { // 左闭右闭

		mid := left + (right-left)/2

		// 此时我们开始判断 nums[mid] 对应目标的 3 种情况

		// 1. nums[mid] > target, 相当于目标在当前区间的左侧
		// 此时右侧指针将设置为 mid - 1（mid 此时的位置已经判断了）
		if nums[mid] > target {
			// 左闭右开, target 在左区间，所以[left, middle)
			// right = mid
			// 左闭右闭, target 在左区间，所以[left, middle - 1]
			right = mid - 1
		}
		// 2. nums[mid] = target, 不用多说吧
		if nums[mid] == target {
			return mid
		}
		// 3. nums[mid] < target, 相当于目标在当前区间的右侧
		// 此时左侧指针被设置为 mid + 1
		if nums[mid] < target {
			left = mid + 1
		}

	}
	// 当循环结束还未获取到目标位置，返回未找到
	return -1
}
