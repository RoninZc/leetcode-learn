package main

import "fmt"

func main() {
	head := InitListNode([]int{1, 2, 3, 4, 5})
	head = reverseList(head)
	OutPut(head)
}

func OutPut(head *ListNode) {
	tmp := head
	for tmp != nil {
		fmt.Println(tmp.Val)
		tmp = tmp.Next
	}
}

func InitListNode(data []int) *ListNode {
	head := &ListNode{
		Val:  0,
		Next: nil,
	}
	tmp := head
	for _, datum := range data {
		tmp.Next = &ListNode{
			Val:  datum,
			Next: nil,
		}
		tmp = tmp.Next
	}
	return head.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	nextNode := head.Next
	head.Next = nil
	for nextNode != nil {
		nextNext := nextNode.Next
		nextNode.Next = head
		head, nextNode = nextNode, nextNext
	}
	return head
}
