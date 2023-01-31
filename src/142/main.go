package main

import "fmt"

func main() {
	head := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	head.Next.Next.Next.Next = head.Next

	fmt.Println(detectCycle(head))
}

// 双指针
func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			tmpFast, tmpHead := fast, head
			for tmpFast != tmpHead {
				tmpFast, tmpHead = tmpFast.Next, tmpHead.Next
			}
			return tmpHead
		}
	}
	return nil
}

// 暴力
// func detectCycle(head *ListNode) *ListNode {
// 	stack, tmpNode := make([]*ListNode, 0), &ListNode{Next: head}
// 	for tmpNode.Next != nil {
// 		for _, node := range stack {
// 			if node == tmpNode.Next {
// 				return node
// 			}
// 		}
// 		stack = append(stack, tmpNode.Next)
// 		tmpNode = tmpNode.Next
// 	}
// 	return nil
// }

type ListNode struct {
	Val  int
	Next *ListNode
}
