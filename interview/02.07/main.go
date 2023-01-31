package main

import "fmt"

// https://leetcode-cn.com/problems/intersection-of-two-linked-lists-lcci/
func main() {
	headA, headB := InitListNode([]int{4, 1, 8, 4, 5}), InitListNode([]int{5, 0, 1, 8, 4, 5})
	OutPut(getIntersectionNode(headA, headB))
}

// 双指针
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	tmpHeadA, lenA, tmpHeadB, lenB := headA, 0, headB, 0
	for tmpHeadA != nil {
		lenA++
		tmpHeadA = tmpHeadA.Next
	}
	for tmpHeadB != nil {
		lenB++
		tmpHeadB = tmpHeadB.Next
	}
	counter := 0
	var fast *ListNode
	var slow *ListNode
	if lenA > lenB {
		counter = lenA - lenB
		fast, slow = headA, headB
	} else {
		counter = lenB - lenA
		fast, slow = headB, headA
	}
	for counter > 0 || fast != slow {
		if counter > 0 {
			counter--
			fast = fast.Next
			continue
		}
		fast = fast.Next
		slow = slow.Next
	}

	return fast
}

// 栈
// func getIntersectionNode(headA, headB *ListNode) *ListNode {
// 	stackA, stackB := make([]*ListNode, 0), make([]*ListNode, 0)
// 	tmpHeadA, tmpHeadB := headA, headB
// 	for tmpHeadA != nil {
// 		stackA = append(stackA, tmpHeadA)
// 		tmpHeadA = tmpHeadA.Next
// 	}
// 	for tmpHeadB != nil {
// 		stackB = append(stackB, tmpHeadB)
// 		tmpHeadB = tmpHeadB.Next
// 	}
//
// 	var node *ListNode
// 	for a, b := len(stackA)-1, len(stackB)-1; a >= 0 && b >= 0; {
// 		if stackA[a] == stackB[b] {
// 			node = stackA[a]
// 		} else {
// 			break
// 		}
// 		a--
// 		b--
// 	}
//
// 	return node
// }

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
