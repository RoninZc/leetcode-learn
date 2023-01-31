package main

import (
	"fmt"
)

func main() {
	fmt.Println(countBits(5))
}

func countBits(num int) []int {
	nums := make([]int, num+1)
	for i := 1; i <= num; i++ {
		if i&1 == 1 {
			nums[i] = nums[i-1] + 1
		} else {
			nums[i] = nums[i/2]
		}
	}
	return nums
}

// func countBits(num int) []int {
// 	nums := make([]int, num+1)
// 	for i := 0; i <= num; i++ {
// 		nums[i] = bits.OnesCount(uint(i))
// 	}
// 	return nums
// }

// func onesCount(x int) (ones int) {
// 	for ; x > 0; x &= x - 1 {
// 		ones++
// 	}
// 	return
// }
