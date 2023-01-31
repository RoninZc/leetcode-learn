package main

import (
	"fmt"
)

func main() {
	// fmt.Println(maxTurbulenceSize([]int{9, 4, 2, 10, 7, 8, 8, 1, 9}))
	fmt.Println(maxTurbulenceSize([]int{37, 199, 60, 296, 257, 248, 115, 31, 273, 176}))
}

func maxTurbulenceSize(arr []int) int {
	length, maxLen := len(arr), 1
	up, down := make([]int, length), make([]int, length)

	for i := 0; i < length; i++ {
		up[i] = 1
		down[i] = 1
	}

	for i := 1; i < length; i++ {
		if arr[i-1] < arr[i] {
			up[i] = down[i-1] + 1
			down[i] = 1
		} else if arr[i-1] > arr[i] {
			up[i] = 1
			down[i] = up[i-1] + 1
		} else {
			up[i] = 1
			down[i] = 1
		}

		maxLen = max(maxLen, max(up[i], down[i]))
	}
	return maxLen
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
