package main

import "fmt"

func main() {
	nm := Constructor([][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	})

	fmt.Println(nm.SumRegion(2, 1, 4, 3))
	fmt.Println(nm.SumRegion(1, 1, 2, 2))
	fmt.Println(nm.SumRegion(1, 2, 2, 4))
}

// NumMatrix 存储
type NumMatrix struct {
	sums [][]int
}

// Constructor 构造NumMatrix方法
func Constructor(matrix [][]int) NumMatrix {
	sums := make([][]int, len(matrix)+1)
	for k, v := range matrix {
		sum := make([]int, len(v)+1)
		for key, value := range v {
			sum[key+1] = sum[key] + value
		}
		sums[k] = sum
	}

	return NumMatrix{sums: sums}
}

// SumRegion 计算区域和
func (nm *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum := 0
	for i := row1; i <= row2; i++ {
		sum += nm.sums[i][col2+1] - nm.sums[i][col1]
	}
	return sum
}
