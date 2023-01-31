package main

import "fmt"

func main() {
	var num int
	_, _ = fmt.Scanln(&num)
	for num != 0 {
		fmt.Println(soda(num))
		_, _ = fmt.Scanln(&num)
	}
}

// 华为编程
func soda(num int) int {
	sum := 0
	for num >= 3 {
		tmp := num / 3
		num = num % 3
		sum += tmp
		num += tmp
	}

	if num == 2 {
		sum++
	}
	return sum
}
