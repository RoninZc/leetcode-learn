package main

import "fmt"

func main() {
	fmt.Println(equalSubstring("abcd", "bcdf", 3))
	fmt.Println(equalSubstring("abcd", "cdef", 3))
	fmt.Println(equalSubstring("abcd", "acde", 0))
	fmt.Println(equalSubstring("krrgw", "zjxss", 19))
}

func equalSubstring(s string, t string, maxCost int) int {
	length := len(s)
	maxLength, left, cost := 0, 0, 0

	for i := 0; i < length; i++ {
		cost += abs(int(s[i]) - int(t[i]))
		if cost > maxCost {
			cost -= abs(int(s[left]) - int(t[left]))
			left++
			continue
		}

		if i-left+1 > maxLength {
			maxLength = i - left + 1
		}
	}

	return maxLength
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
