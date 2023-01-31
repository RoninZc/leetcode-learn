package main

import "fmt"

func main() {
	var str []byte
	for {
		i, _ := fmt.Scan(&str)
		if i == 0 {
			return
		}
		foo(str)
	}
}

func foo(str []byte) {
	if len(str) <= 2 {
		return
	}
	str = str[2:]
	num, p := 0, len(str)-1
	base := 1
	for i := 0; i < p; i++ {
		base *= 16
	}

	for _, b := range str {
		bNum := int(b)
		if bNum > 57 {
			bNum -= 55
		} else {
			bNum -= 48
		}
		num += base * bNum

		base = base / 16
	}
	fmt.Println(num)
}
