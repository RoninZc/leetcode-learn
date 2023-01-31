package main

import "fmt"

func main() {
	head := InitListNode([]int{1, 2, 3, 4, 5})
	head = removeNthFromEnd(head, 2)
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

// 栈版本
// func removeNthFromEnd(head *ListNode, n int) *ListNode {
// 	nodeStack, tmpNode := make([]*ListNode, 0), head
// 	for tmpNode != nil {
// 		nodeStack = append(nodeStack, tmpNode)
// 		tmpNode = tmpNode.Next
// 	}
// 	if n == len(nodeStack) {
// 		return head.Next
// 	}
// 	nodeStack[len(nodeStack)-1-n].Next = nodeStack[len(nodeStack)-n].Next
// 	return head
// }

// 双指针
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	startNode := &ListNode{
		Val:  0,
		Next: head,
	}
	fast, slow, counter := startNode, startNode, 0

	for fast != nil {
		fast = fast.Next
		if counter > n {
			slow = slow.Next
		}
		counter++
	}
	slow.Next = slow.Next.Next
	return startNode.Next
}
