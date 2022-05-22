package LinkedList

/*
 *  https://leetcode-cn.com/problems/design-linked-list/
 */

type MyLinkedList struct {
	ListLength int
	dummy *Node
	tail *Node
}

type Node struct {
	Val int
	NextNode *Node
}


/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	list := MyLinkedList{
	}
	preNode := &Node{}
	list.ListLength = 0
	list.dummy = preNode
	list.tail = preNode
	return list
}


/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if !(index>=0 && index<this.ListLength) {
		return -1
	}
	cur := this.dummy
	for i:=-1; i<index; i++ {
		cur = cur.NextNode
	}
	return cur.Val
}


/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int)  {
	insetNode := &Node{
		Val: val,
		NextNode: nil,
	}
	insetNode.NextNode = this.dummy.NextNode
	this.dummy.NextNode = insetNode
	if this.dummy == this.tail {
		this.tail = insetNode
	}
	this.ListLength++
}


/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int)  {
	insetNode := &Node{
		Val: val,
		NextNode: nil,
	}
	this.tail.NextNode = insetNode
	this.tail = insetNode
	this.ListLength++
}


/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int)  {
	if index > this.ListLength {
		return
	}
	insetNode := &Node{
		Val: val,
		NextNode: nil,
	}
	if index == this.ListLength {
		this.AddAtTail(val)
		return
	}
	if index <= 0 {
		this.AddAtHead(val)
		return
	}
	cur := this.dummy
	for i:=-1; i<index-1; i++ {
		cur = cur.NextNode
	}
	insetNode.NextNode = cur.NextNode
	cur.NextNode = insetNode
	this.ListLength++
}


/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int)  {
	if !(index>=0 && index<this.ListLength) {
		return
	}
	cur := this.dummy
	for i:=-1; i<index-1; i++ {
		cur = cur.NextNode
	}
	if index == this.ListLength-1 {
		cur.NextNode = nil
		this.tail = cur
		this.ListLength--
		return
	}
	cur.NextNode = cur.NextNode.NextNode
	this.ListLength--
}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */