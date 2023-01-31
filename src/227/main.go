package main

// import (
// 	"math"
// 	"strconv"
// 	"strings"
// )

// func main() {

// }

// // 使用map维护运算符的优先级
// var m = map[byte]int{
// 	'-': 1,
// 	'+': 1,
// 	'*': 2,
// 	'/': 2,
// 	'%': 2,
// 	'^': 3,
// }

// func calculate(s string) int {
// 	// 去除所有的空格
// 	s = strings.Replace(s, " ", "", -1)
// 	// 将所有的（-替换成（0-
// 	s = strings.Replace(s, "(-", "(0-", -1)
// 	n := len(s)

// 	// 存放所有数字,长度为1 ，防止第一个数是负数
// 	nums := make([]int, 1)
// 	// 存放所有操作
// 	ops := make([]byte, 0)

// 	for i := 0; i < n; i++ {
// 		char := s[i]

// 		if char == '(' {
// 			ops = append(ops, char)
// 		} else if char == ')' {
// 			for len(ops) > 0 {
// 				if ops[len(ops)-1] != '(' {
// 					calc(&nums, &ops)
// 				} else {
// 					ops = ops[:len(ops)-1]
// 					break
// 				}
// 			}
// 		} else {
// 			if isNum(char) {
// 				u, j := 0, i
// 				for j < n && isNum(byte(nums[j])) {
// 					u = u*10 + nums[j] - '0'
// 					j++
// 				}
// 				nums = append(nums, u)
// 				i = j - 1
// 			} else {

// 			}
// 		}
// 	}
// 	return 0
// }

// func calc(_nums *[]int, _ops *[]byte) {
// 	nums, ops := *_nums, *_ops

// 	if len(nums) < 2 || len(ops) == 0 {
// 		return
// 	}

// 	op := ops[len(ops)-1]
// 	ops = ops[:len(ops)-1]

// 	b := nums[len(nums)-1]
// 	nums = nums[:len(nums)-1]

// 	a := nums[len(nums)-1]
// 	nums = nums[:len(nums)-1]

// 	ans := 0
// 	switch op {
// 	case '+':
// 		ans = a + b
// 	case '-':
// 		ans = a - b
// 	case '*':
// 		ans = a * b
// 	case '/':
// 		ans = a / b
// 	case '^':
// 		ans = int(math.Pow(float64(a), float64(b)))
// 	case '%':
// 		ans = a % b
// 	}

// 	nums = append(nums, ans)
// }

// func isNum(s byte) bool {
// 	_, err := strconv.ParseInt(string(s), 10, 32)
// 	return err == nil
// }
