package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				3,
				&ListNode{
					4,
					nil,
				},
			},
		},
	}
	l2 := &ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				3,
				nil,
			},
		},
	}
	t := addTwoNumbers(l1, l2)
	for t != nil {
		fmt.Println(t.Val)
		t = t.Next
	}
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root := &ListNode{}
	res := root
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		l1Val := 0
		if l1 != nil {
			l1Val = l1.Val
		}
		l2Val := 0
		if l2 != nil {
			l2Val = l2.Val
		}
		sum := carry + l1Val + l2Val
		carry = sum / 10
		res.Next = &ListNode{sum % 10, nil}
		res = res.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	return root.Next
}

