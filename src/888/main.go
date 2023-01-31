package main

import "fmt"

func main() {
	fmt.Println(fairCandySwap([]int{1, 1}, []int{2, 2}))
}

func fairCandySwap(A []int, B []int) []int {
	sumA, sumB, mB := 0, 0, make(map[int]struct{}, len(B))

	for _, v := range A {
		sumA += v
	}
	for _, v := range B {
		sumB += v
		mB[v] = struct{}{}
	}

	changeA := (sumB - sumA) / 2

	for _, v := range A {
		if _, ok := mB[v+changeA]; ok {
			return []int{v, v + changeA}
		}
	}
	return nil
}
