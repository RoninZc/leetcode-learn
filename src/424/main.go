package main

import (
	"fmt"
)

func main() {
	fmt.Println(characterReplacement("AABABBA", 1))
}

func characterReplacement(s string, k int) int {
	cnt := [26]int{}     // 窗口内每个字符出现的次数
	maxCnt, left := 0, 0 // 最长长度, 左边界

	// 右边界
	for right, ch := range s {
		cnt[ch-'A']++ // 每次循环 对应的字符出现次数增加1

		// 判断最长的长度和当前字符谁更长
		maxCnt = max(maxCnt, cnt[ch-'A'])
		// 如果右边距 - 左边距 [窗口大小] + 1 - 最长的长度 超过了可以替换的次数
		// 则左边距向右移动一格 同时 cnt 窗口内出现次数减少1
		if right-left+1-maxCnt > k {
			cnt[s[left]-'A']--
			left++
		}
	}

	return len(s) - left
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
