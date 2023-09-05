package linked_stack

import "errors"

// 基于链表实现栈

// 链表节点
type ListNode struct {
	next *ListNode
	val  interface{}
}

type LinkedStack struct {
	head *ListNode
	size int
}

func NewLinkedStack() *LinkedStack {
	return &LinkedStack{
		head: nil,
		size: 0,
	}
}

// 入栈 -- 采用头插法
func (stack *LinkedStack) Push(val interface{}) {
	node := &ListNode{
		next: nil,
		val: val,
	}
	if stack.Empty() {
		stack.head = node
	} else {
		node.next = stack.head
		stack.head = node
	}
	stack.size++
}

// 出栈 -- 弹出链表首部
func (stack *LinkedStack) Pop() (interface{}, error) {
	if stack.Empty() {
		return nil, errors.New("stack is empty")
	}

	result := stack.head.val
	stack.head = stack.head.next
	stack.size--

	return result, nil
}

// 获取栈顶元素
func (stack *LinkedStack) Peek() (interface{}, error) {
	 if stack.Empty() {
	 	return nil, errors.New("stack is empty")
	 }
	 return stack.head.val, nil
}


// 判断是否为空
func (stack *LinkedStack) Empty() bool {
	return stack.size <= 0
}

func (stack *LinkedStack) Size() int {
	return stack.size
}


