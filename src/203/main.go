package main

import "fmt"

func main() {
	head := new(ListNode)
	next := head
	for _, datum := range []int{7, 7, 7, 7} {
		tmp := &ListNode{
			Val:  datum,
			Next: nil,
		}
		next.Next, next = tmp, tmp
	}
	head = head.Next

	head = removeElements(head, 7)

	tmp := head
	for tmp != nil {
		fmt.Println(tmp.Val)
		tmp = tmp.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// func removeElements(head *ListNode, val int) *ListNode {
// 	lastNode := &ListNode{
// 		Val:  0,
// 		Next: head,
// 	}
// 	res, node := lastNode, lastNode.Next
// 	for node != nil {
// 		if node.Val == val {
// 			lastNode.Next = node.Next
// 			node = node.Next
// 			continue
// 		}
// 		lastNode, node = node, node.Next
// 	}
//
// 	return res.Next
// }

// 优雅写法
func removeElements(head *ListNode, val int) *ListNode {
	newHead := &ListNode{
		Next: head,
	}
	node := newHead
	for node != nil && node.Next != nil {
		if node.Next.Val == val {
			node.Next = node.Next.Next
		} else {
			node = node.Next
		}
	}
	return newHead.Next
}
