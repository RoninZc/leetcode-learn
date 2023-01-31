package main

import "fmt"

func main() {
	matrix := generateMatrix(4)
	for _, ints := range matrix {
		fmt.Println(ints)
	}
}

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	counter, state := 1, 0
	length, lengthLeft, lengthRight := 0, 0, n-1
	wide, wideHigh, wideEnd := 0, n-1, 0
	for counter <= n*n {
		switch state {
		case 0:
			matrix[wide][length] = counter
			if length == lengthRight {
				wideEnd++
				state = 1
				continue
			}
			length++
		case 1:
			matrix[wide][length] = counter
			if wide == wideHigh {
				lengthRight--
				state = 2
				continue
			}
			wide++
		case 2:
			matrix[wide][length] = counter
			if length == lengthLeft {
				wideHigh--
				state = 3
				continue
			}
			length--
		case 3:
			matrix[wide][length] = counter
			if wide == wideEnd {
				lengthLeft++
				state = 0
				continue
			}
			wide--
		}
		counter++
	}

	return matrix
}
