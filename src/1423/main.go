package main

import "fmt"

func main() {
	fmt.Println(maxScore([]int{1, 2, 3, 4, 5, 6, 1}, 3))
	fmt.Println(maxScore([]int{2, 2, 2}, 2))
	fmt.Println(maxScore([]int{9, 7, 7, 9, 7, 7, 9}, 7))
	fmt.Println(maxScore([]int{1, 1000, 1}, 1))
	fmt.Println(maxScore([]int{1, 79, 80, 1, 1, 1, 200, 1}, 3))
	fmt.Println(maxScore([]int{11, 49, 100, 20, 86, 29, 72}, 4))
	fmt.Println(maxScore([]int{96, 90, 41, 82, 39, 74, 64, 50, 30}, 8)) // 536
}

func maxScore(cardPoints []int, k int) int {
	cardPointsSum := sum(cardPoints)
	length := len(cardPoints)
	minSum := sum(cardPoints[:length-k])
	sum := minSum

	for i := 0; i < k; i++ {
		sum = sum - cardPoints[i] + cardPoints[length-k+i]
		if sum < minSum {
			minSum = sum
		}
	}
	return cardPointsSum - minSum
}

func sum(nums []int) (sum int) {
	for _, v := range nums {
		sum += v
	}
	return sum
}
