package main

import (
	"fmt"
)

func main() {
	linkedList := Constructor()
	linkedList.AddAtHead(1)
	linkedList.AddAtTail(3)
	OutPut(&linkedList)
	linkedList.AddAtIndex(1, 2) // 链表变为1-> 2-> 3
	OutPut(&linkedList)
	fmt.Println(linkedList.Get(1)) // 返回2
	linkedList.DeleteAtIndex(1)    // 现在链表是1-> 3
	OutPut(&linkedList)
	fmt.Println(linkedList.Get(1)) // 返回3
}

type MyLinkedList struct {
	Val  int
	Next *MyLinkedList
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func OutPut(list *MyLinkedList) {
	tmp := list.Next
	for tmp != nil {
		fmt.Print(tmp.Val)
		if tmp.Next != nil {
			fmt.Print("->")
		}
		tmp = tmp.Next
	}
	fmt.Println()
}

func (this *MyLinkedList) Get(index int) int {
	tmp, counter := this.Next, 0
	for tmp != nil {
		if counter == index {
			return tmp.Val
		}
		tmp = tmp.Next
		counter++
	}
	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.Next = &MyLinkedList{
		Val:  val,
		Next: this.Next,
	}
}

func (this *MyLinkedList) AddAtTail(val int) {
	tmp := this
	for tmp != nil {
		if tmp.Next == nil {
			tmp.Next = &MyLinkedList{Val: val}
			return
		}
		tmp = tmp.Next
	}
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index <= 0 {
		this.AddAtHead(val)
		return
	}
	tmp, counter := this, 0
	for tmp != nil {
		if counter == index {
			tmp.Next = &MyLinkedList{
				Val:  val,
				Next: tmp.Next,
			}
			return
		}
		tmp = tmp.Next
		counter++
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	node, counter := this, 0
	for node != nil && node.Next != nil {
		if counter == index {
			node.Next = node.Next.Next
		} else {
			node = node.Next
		}
		counter++
	}
}
