package main

func main() {

}

// MyQueue 队列
type MyQueue struct {
	in  Stack
	out Stack
}

// Constructor MyQueue构造方法
func Constructor() MyQueue {
	in := new(Stack)
	out := new(Stack)
	return MyQueue{
		in:  *in,
		out: *out,
	}
}
func (mq *MyQueue) in2out() {
	for len(mq.in) > 0 {
		mq.out.Push(mq.in.Pop())
	}
}

// Push Push
func (mq *MyQueue) Push(x int) {
	mq.in.Push(x)
}

// Pop Pop
func (mq *MyQueue) Pop() int {
	if mq.out.Len() == 0 {
		mq.in2out()
	}
	return mq.out.Pop()
}

// Peek Peek
func (mq *MyQueue) Peek() int {
	if mq.out.Len() == 0 {
		mq.in2out()
	}
	return mq.out.Top()
}

// Empty Empty
func (mq *MyQueue) Empty() bool {
	return mq.in.IsEmpty() && mq.out.IsEmpty()
}

// Stack 栈
type Stack []int

// Len 获取栈的长度
func (stack *Stack) Len() int {
	return len(*stack)
}

// Push push
func (stack *Stack) Push(value int) {
	*stack = append(*stack, value)
}

// Top 获取栈的第一个
func (stack *Stack) Top() int {
	return (*stack)[len(*stack)-1]
}

// Pop 弹出最后一个
func (stack *Stack) Pop() int {
	value := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return value
}

// IsEmpty 判断是否为空
func (stack *Stack) IsEmpty() bool {
	return len(*stack) == 0
}
