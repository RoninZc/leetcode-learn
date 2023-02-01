package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world!")
	fmt.Println(decodeMessage("the quick brown fox jumps over the lazy dog", "vkbs bs t suepuv"))
	fmt.Println(decodeMessage("eljuxhpwnyrdgtqkviszcfmabo", "zwx hnfx lqantp mnoeius ycgk vcnjrdb"))
}

func decodeMessage(key string, message string) string {
	m := key2map(key)
	str := make([]byte, len(message))

	for i := 0; i < len(message); i++ {
		if v, ok := m[message[i]]; ok {
			str[i] = v
		} else {
			str[i] = ' '
		}
	}
	return string(str)
}

// 获取密码表
func key2map(key string) map[byte]byte {
	m := make(map[byte]byte)
	b := byte('a')
	space := byte(' ')

	for i := 0; i < len(key); i++ {
		if key[i] == space {
			continue
		}

		if _, ok := m[key[i]]; !ok {
			m[key[i]] = b
			b++
		}
	}
	return m
}
