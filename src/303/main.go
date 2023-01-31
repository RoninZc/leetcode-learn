package main

// NumArray a
type NumArray struct {
	sums []int
}

// Constructor a
func Constructor(nums []int) NumArray {
	sums := make([]int, len(nums)+1)
	for i, v := range nums {
		sums[i+1] = sums[i] + v
	}
	return NumArray{sums}
}

// SumRange a
func (na *NumArray) SumRange(i int, j int) int {
	return na.sums[j+1] - na.sums[i]
}
