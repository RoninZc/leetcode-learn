package main

import "fmt"

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram2("anagram", "nagaram"))
}

func isAnagram2(s, t string) bool {
	dic := make([]int, 26)
	for _, b := range s {
		dic[b-'a']++
	}
	for _, b := range t {
		dic[b-'a']--
	}
	for _, count := range dic {
		if count != 0 {
			return false
		}
	}
	return true
}

// 暴力map
func isAnagram(s string, t string) bool {
	sDic, tDic := make(map[byte]int), make(map[byte]int)
	for _, b := range []byte(s) {
		if _, ok := sDic[b]; ok {
			sDic[b]++
		} else {
			sDic[b] = 1
		}
	}
	for _, b := range []byte(t) {
		if _, ok := tDic[b]; ok {
			tDic[b]++
		} else {
			tDic[b] = 1
		}
	}

	for char, count := range sDic {
		if tCount, ok := tDic[char]; !ok || tCount != count {
			return false
		}
	}
	for char, count := range tDic {
		if sCount, ok := sDic[char]; !ok || sCount != count {
			return false
		}
	}
	return true
}
