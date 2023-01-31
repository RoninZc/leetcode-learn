package main

import (
	"fmt"
)

func main() {
	fmt.Println(removeDuplicates("abbaca"))
}

func removeDuplicates(S string) string {
	n := len(S)
	bys := []byte(S)
	stack := make([]byte, 0, n)
	for i := 0; i < n; i++ {
		if len(stack) == 0 {
			stack = append(stack, bys[i])
			continue
		}
		if bys[i] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, bys[i])
		}
	}
	return string(stack)
}
