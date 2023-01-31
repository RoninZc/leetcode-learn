package main

import "fmt"

func main() {
	fmt.Println(lastRemaining(5, 3))
}

func lastRemaining(n int, m int) int {
	if n == 1 {
		return 0
	}
	x := lastRemaining(n-1, m)
	return (m + x) % n
}
