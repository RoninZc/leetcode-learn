package main

import "fmt"

func main() {
	head := InitListNode([]int{1, 2})
	head = swapPairs(head)
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

func swapPairs(head *ListNode) *ListNode {
	dummyNode := &ListNode{
		Val:  0,
		Next: head,
	}
	tmpNode := dummyNode
	for tmpNode.Next != nil && tmpNode.Next.Next != nil {
		one := tmpNode.Next
		two := tmpNode.Next.Next.Next

		tmpNode.Next = tmpNode.Next.Next
		tmpNode.Next.Next = one
		tmpNode.Next.Next.Next = two

		tmpNode = tmpNode.Next.Next
	}

	return dummyNode.Next
}
