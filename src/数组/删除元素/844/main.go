package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world!")
	fmt.Println(backspaceCompare1("ab#c", "ad#c"))
	fmt.Println(backspaceCompare1("ab##", "c#d#"))
	fmt.Println(backspaceCompare1("a#c", "b"))
	fmt.Println(backspaceCompare2("ab#c", "ad#c"))
	fmt.Println(backspaceCompare2("ab##", "c#d#"))
	fmt.Println(backspaceCompare2("a#c", "b"))
}

// 快慢指针法
func backspaceCompare1(s string, t string) bool {
	return string(fmtString([]byte(s))) == string(fmtString([]byte(t)))
}

func fmtString(str []byte) []byte {
	slowIndex := 0
	for faseIndex := 0; faseIndex < len(str); faseIndex++ {
		if str[faseIndex] != '#' {
			str[faseIndex], str[slowIndex] = str[slowIndex], str[faseIndex]
			slowIndex++
		} else if slowIndex >= 1 {
			slowIndex--
		}
	}
	return str[:slowIndex]
}

// 倒叙双指针法
func backspaceCompare2(s string, t string) bool {
	skipS, skipT := 0, 0
	sIndex, tIndex := len(s)-1, len(t)-1

	for sIndex >= 0 || tIndex >= 0 {
		// 从末尾开始，找到不被删除的字符
		for sIndex >= 0 {
			if s[sIndex] == '#' {
				skipS++
				sIndex--
			} else if skipS > 0 {
				skipS--
				sIndex--
			} else {
				break
			}
		}
		for tIndex >= 0 {
			if t[tIndex] == '#' {
				skipT++
				tIndex--
			} else if skipT > 0 {
				skipT--
				tIndex--
			} else {
				break
			}
		}

		if sIndex >= 0 && tIndex >= 0 {
			if s[sIndex] != t[tIndex] {
				return false
			}
		} else if sIndex >= 0 || tIndex >= 0 {
			// 当任意一字符串遍历到末尾, 另一字符串未遍历到末尾时，代表两字符串长度不一致
			return false
		}
		sIndex--
		tIndex--
	}
	return true
}
