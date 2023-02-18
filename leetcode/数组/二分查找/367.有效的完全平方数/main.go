package main

import (
	"fmt"
)

/*
给你一个正整数 num 。如果 num 是一个完全平方数，则返回 true ，否则返回 false 。

完全平方数 是一个可以写成某个整数的平方的整数。换句话说，它可以写成某个整数和自身的乘积。

不能使用任何内置的库函数，如  sqrt 。



示例 1：

输入：num = 16
输出：true
解释：返回 true ，因为 4 * 4 = 16 且 4 是一个整数。
示例 2：

输入：num = 14
输出：false
解释：返回 false ，因为 3.742 * 3.742 = 14 但 3.742 不是一个整数。


提示：

1 <= num <= 231 - 1

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/valid-perfect-square
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

func main() {
	fmt.Println("hello world!")
	fmt.Println(isPerfectSquare(16))
	fmt.Println(isPerfectSquare(14))
}

func isPerfectSquare(num int) bool {

	// 二分法，左闭右开
	left, right := 0, num+1

	for left < right {

		mid := (left + right) / 2
		if mid*mid == num {
			return true
		}
		if mid*mid > num {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return false
}
