package main

import "fmt"

func main() {
	for {
		var n, m int
		i, _ := fmt.Scanf("%d%d", &n, &m)
		if i == 0 {
			return
		}
		nums := make(Nums, 0, n)
		for i := 0; i < n; i++ {
			var num int
			_, _ = fmt.Scan(&num)
			nums = append(nums, num)
		}

		for i := 0; i < m; i++ {
			var op string
			var a, b int
			_, _ = fmt.Scanf("%s %d %d", &op, &a, &b)
			switch op {
			case "Q":
				fmt.Println(nums.Find(a, b))
			case "U":
				nums.Update(a, b)
			}
		}
	}
}

type Nums []int

func (n *Nums) Find(left, right int) (max int) {
	if left > right {
		left, right = right, left
	}
	for i := left - 1; i < right; i++ {
		if (*n)[i] > max {
			max = (*n)[i]
		}
	}
	return
}
func (n *Nums) Update(index, num int) {
	(*n)[index-1] = num
}
