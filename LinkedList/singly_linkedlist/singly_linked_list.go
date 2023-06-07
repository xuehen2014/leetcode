package singly_linkedlist

import "fmt"

//单项链表

type Node struct {
	Data interface{}
	Next *Node
}

type LList struct {
	Header *Node //指向第一个节点
	Length int
}

func CreateNode(v interface{}) *Node{
	return &Node{v, nil}
}
func CreateList()  *LList{
	header := CreateNode(nil)
	return &LList{header,0}
}

//从链表头部增加一个节点
func (l *LList) Add(data interface{}) {
	newNode := CreateNode(data)
	defer func() {
		l.Length++
	}()

	if l.Length == 0 {    //当前链表为空链表
		l.Header = newNode
	} else {
		newNode.Next = l.Header
		l.Header = newNode
	}
}

//在链表尾部增加一个节点
func (l *LList) Append(data interface{}) {
	newNode := CreateNode(data)
	defer func() {
		l.Length++
	}()

	if l.Length == 0 {
		l.Header = newNode
	} else {
		current := l.Header
		for current.Next != nil { //循环指向最后一个节点
			current = current.Next
		}
		current.Next = newNode
	}
}

//在第i个节点后插入节点
func (l *LList) Insert(i int, data interface{}) {
	defer func() {
		l.Length++
	}()

	if i>= l.Length {
		l.Append(data)
		return
	}

	newNode := CreateNode(data)

	//找到第i个节点
	j:=1
	pre := l.Header
	for j != i {
		pre=pre.Next
		j++
	}
	
	after := pre.Next
	pre.Next = newNode
	newNode.Next = after
	return
}

//遍历链表
func (l *LList) Scan() {
	current := l.Header
	i := 1
	for current != nil {
		fmt.Printf("第%d的节点是%d\n", i, current.Data)
		i++
		current = current.Next
	}
}




