package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world!")
	fmt.Println(totalFruit([]int{1, 2, 1}))
	fmt.Println(totalFruit([]int{0, 1, 2, 2}))
	fmt.Println(totalFruit([]int{1, 2, 3, 2, 2}))
	fmt.Println(totalFruit([]int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}))
}

// 该题可以转换成 求一个数组内只包含两种元素的最长子数组
func totalFruit(fruits []int) (length int) {
	resMap := map[int]int{}
	left := 0
	for right, x := range fruits {
		resMap[x]++
		for len(resMap) > 2 {
			y := fruits[left]
			resMap[y]--
			if resMap[y] == 0 {
				delete(resMap, y)
			}
			left++
		}
		length = max(length, right-left+1)
	}
	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
