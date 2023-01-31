package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(medianSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}

func medianSlidingWindow(nums []int, k int) []float64 {
	knums := make([]int, k)
	copy(knums, nums[:k])
	sort.Ints(knums)
	ret := make([]float64, len(nums) - k + 1)
	ret[0] = calMid(knums, k)
	for i := k; i < len(nums); i++ {
		delIdx := binarySearch(knums, nums[i-k])
		knums = append(knums[:delIdx], knums[delIdx+1:]...)
		addIdx := binarySearch(knums, nums[i])
		last := append([]int{nums[i]}, knums[addIdx:]...)
		knums = append(knums[:addIdx], last...)
		ret[i-k+1] = calMid(knums, k)
	}
	return ret
}

func binarySearch(nums []int, val int) int {
	low, high := 0, len(nums) - 1
	for low <= high {
		mid := low + (high - low)/2
		if nums[mid] == val {
			return mid
		} else if nums[mid] > val {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}

func calMid(nums []int, k int) float64 {
	n1, n2 := 0, 0
	if k%2 == 0 {
		n1, n2 = nums[k/2-1], nums[k/2]
	} else {
		n1, n2 = nums[k/2], nums[k/2]
	}
	return float64(n1 + n2)/float64(2)
}


//func medianSlidingWindow(nums []int, k int) []float64 {
//	lens := len(nums)
//	res := make([]float64, lens-k+1)
//	window := make([]int, k)
//	// 添加初始值
//	copy(window, nums[:k])
//
//	// 冒泡排序
//	sort.Ints(window)
//	res[0] = getMid(window)
//
//	for i := 0; i < lens-k; i++ {
//		index := search(window, nums[i])
//		window[index] = nums[i+k]
//		sort.Ints(window)
//		res[i+1] = getMid(window)
//	}
//	return res
//}
//
////func sort(window []int) {
////	lens := len(window)
////	for i := 0; i < lens-1; i++ {
////		for j := 0; j < lens-1-i; j++ {
////			if window[j] > window[j+1] {
////				window[j], window[j+1] = window[j+1], window[j]
////			}
////		}
////	}
////}
//
//// 求slice的中位数
//func getMid(window []int) float64 {
//	lens := len(window)
//	if lens%2 == 0 {
//		return float64(window[lens/2-1]+window[lens/2]) / 2.0
//	} else {
//		return float64(window[lens/2])
//	}
//}
//
//// 二分查找
//func search(window []int, target int) int {
//	start := 0
//	end := len(window) - 1
//	for start <= end {
//		mid := start + (end-start)/2
//		if window[mid] > target {
//			end = mid - 1
//		} else if window[mid] < target {
//			start = mid + 1
//		} else {
//			return mid
//		}
//	}
//	return -1
//}
