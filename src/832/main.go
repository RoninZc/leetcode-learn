package main

import "fmt"

func main() {
	fmt.Println(flipAndInvertImage([][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}))
	fmt.Println(flipAndInvertImage([][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}}))
}

func flipAndInvertImage(A [][]int) [][]int {
	for _, v := range A {
		length := len(v)
		isDual := length % 2

		for i := 0; i < length/2+isDual; i++ {
			if v[i] == 0 && v[length-i-1] == 0 {
				v[i], v[length-i-1] = 1, 1
			} else if v[i] == 1 && v[length-i-1] == 1 {
				v[i], v[length-i-1] = 0, 0
			}
		}
	}
	return A
}
