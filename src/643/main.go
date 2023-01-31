package main

import "fmt"

func main() {
	fmt.Println(findMaxAverage([]int{0, 4, 0, 3, 2}, 1))
}

func findMaxAverage(nums []int, k int) float64 {
	sum := 0.0
	for i := 0; i < k; i++ {
		sum += float64(nums[i])
	}
	maxSum := sum
	// avg := sum / float64(k)
	// fmt.Println(avg)
	length := len(nums)
	for i := k; i < length; i++ {
		sum = sum - float64(nums[i-k]) + float64(nums[i])
		// tmp := sum / float64(k)
		if sum > maxSum {
			maxSum = sum
		}
	}

	return float64(maxSum) / float64(k)
}
