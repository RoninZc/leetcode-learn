package main

import (
	"fmt"
)

func main() {
	fmt.Println("*:", calculate("1 + 1"))
	// fmt.Println("*:", calculate("2-(5-6)"))
	// fmt.Println("*:", calculate("(1+(4+5+2)-3)+(6+8)"))
	// fmt.Println("*:", calculate("2147483647"))
	// fmt.Println("*:", calculate("1-(2+3-(4+(5-(1-(2+4-(5+6))))))"))
}

func calculate(s string) int {
	count := len(s)
	symbols := make([]int, 1)
	symbols[0] = 1
	symbol := 1 // 存储上一个符号
	sum := 0
	num := 0
	for i := 0; i < count; i++ {
		if s[i] >= '0' && s[i] <= '9' {
			num = num*10 + int(s[i]-'0')
			continue
		}

		sum += num * symbol
		fmt.Println(sum, symbol)
		num = 0

		switch s[i] {
		case '+':
			// 加号不变符号 还是括号前的符号
			symbol = symbols[len(symbols)-1]
		case '-':
			// 减号需要变号 加号变减号 减号变加号
			symbol = symbols[len(symbols)-1] * -1
		case '(':
			// 将括号前的符号保存
			symbols = append(symbols, symbol)
		case ')':
			// 出括号了，删除最近括号外的符号
			symbols = symbols[:len(symbols)-1]
		}
	}
	sum += num * symbol
	return sum
}

// func calculate(s string) int {
// 	for {
// 		left := strings.LastIndexByte(s, '(')
// 		if left == -1 {
// 			break
// 		}
// 		right := strings.IndexByte(s[left:], ')')
// 		num := calculateStr(string(s[left+1 : left+right]))
// 		s = s[:left] + strconv.Itoa(num) + s[left+right+1:]
// 	}

// 	return calculateStr(s)
// }

// // 计算无括号的字符串
// func calculateStr(s string) (num int) {
// 	symbol := true
// 	tmp := 0
// 	digit := 0
// 	isLastSub := false
// 	isLastAdd := false
// 	for _, v := range s {
// 		switch v {
// 		case '+':
// 			isLastSub = false
// 			isLastAdd = true
// 			fmt.Println("+", num, symbol, tmp)
// 			num = symbolRes(num, symbol, tmp)
// 			digit = 0
// 			symbol = true
// 		case '-':
// 			if isLastSub {
// 				symbol = true
// 			} else if isLastAdd {
// 				symbol = false
// 			} else {
// 				fmt.Println("-", num, symbol, tmp)
// 				num = symbolRes(num, symbol, tmp)
// 				symbol = false
// 			}
// 			digit = 0
// 			isLastSub = true
// 			isLastAdd = false
// 		case ' ':
// 			continue
// 		default:
// 			isLastSub = false
// 			isLastAdd = false
// 			tmp = tmp*10*digit + int(v-'0')
// 			digit = 1
// 		}
// 	}
// 	fmt.Println(num, symbol, tmp)
// 	num = symbolRes(num, symbol, tmp)
// 	return
// }

// func symbolRes(num1 int, symbol bool, num2 int) int {
// 	if symbol {
// 		return num1 + num2
// 	}
// 	return num1 - num2
// }
