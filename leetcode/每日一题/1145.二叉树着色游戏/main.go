package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world!")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 深度优先搜索，计算子树的节点数量
func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	lts, rts := 0, 0
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		ls := dfs(node.Left)
		rs := dfs(node.Right)
		if node.Val == x {
			lts, rts = ls, rs
		}
		return ls + rs + 1
	}

	dfs(root)
	return max(max(lts, rts), n-1-lts-rts)*2 > n

}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
