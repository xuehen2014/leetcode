package doubly_linkedlist

import "fmt"

/*
 * 双向链表
 */

//节点定义
type ListNode struct {
	val        interface{}
	prev, next *ListNode
}

func newNode(val interface{}) *ListNode {
	return &ListNode{val, nil, nil}
}

//链表定义
type LinkedList struct {
	head *ListNode   //头结点
	size int
}

func New() *LinkedList{
	return &LinkedList{nil, 0}
}

//插入数据到链表头部
func (list *LinkedList) PushFront(val interface{}) *ListNode {
	node := newNode(val)
	if list.head != nil {
		list.head.prev = node
		node.next = list.head
	}

	list.head = node
	list.size++
	return node
}

//插入数据到链表尾部
func (list *LinkedList) PushBack(val interface{}) *ListNode {
	if list.head == nil {
		return list.PushFront(val)
	}
	node := newNode(val)

	p := list.head
	for p.next != nil {
		p = p.next
	}

	p.next = node
	node.prev = p
	list.size++
	
	return node
}

//在某个节点之后插入数据
func (list *LinkedList) PushAfterNode(p *ListNode, val interface{}) *ListNode {
	if p == nil {
		return nil
	}
	node := newNode(val)

	next := p.next
	node.next = next
	node.prev = p
	p.next = node
	if next != nil {
		next.prev = node
	}
	list.size++
	return node
}

//在某个节点之前插入元素
func (list *LinkedList)  PushBefore(p *ListNode, val interface{}) *ListNode {
	if p == nil {
		return nil
	}

	prev := p.prev
	if prev == nil {
		return list.PushFront(val)
	}
	node := newNode(val)

	p.prev = node
	node.next = p
	node.prev = prev
	prev.next = node
	list.size++
	
	return node
}

//删除节点
func (list *LinkedList) Delete(p *ListNode) {
	if p == nil || list.head == nil {
		return
	}
	if p == list.head { //头结点
		list.head = p.next
	} else {
		prev := p.prev
		next := p.next
		prev.next = next
		if next != nil {
			next.prev = prev
		}
	}
	list.size--
}

//根据值查找节点
func (list *LinkedList) Find(val interface{}) *ListNode {
	if list.head == nil {
		return nil
	}
	p := list.head
	for p != nil && p.val != val {
		p = p.next
	}
	return p
}

//打印数据
func (list *LinkedList) PrintData() {
	current := list.head
	i :=1
	for current != nil {
		fmt.Printf("第%d的节点是%v\n", i, current.val)
		i++
		current = current.next
	}
}


