package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("hello world!")
	fmt.Println(getMaximumConsecutive([]int{1, 4, 10, 3, 1}))
}

func getMaximumConsecutive(coins []int) (m int) {
	// 一开始只能构造出 0
	sort.Ints(coins)
	for _, c := range coins {
		if c > m+1 { // coins 已排序，后面没有比 c 更小的数了
			break // 无法构造出 m+1，继续循环没有意义
		}
		m += c // 可以构造出区间 [0,m+c] 中的所有整数
	}
	return m + 1 // [0,m] 中一共有 m+1 个整数
}
