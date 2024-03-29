package Queue

import "fmt"

/*
 * https://leetcode-cn.com/problems/design-circular-queue/
 */

type MyCircularQueue struct {
	data     []int //循环队列的存储空间
	used     int   //已经使用的元素个数
	front    int   //第一个元素所在位置-队首
	rear     int   //rear是EnQueue可在存放的位置,待插入的位置, [front,rear)
	capacity int   //循环队列最多可以存放的元素个数
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		capacity: k,
		front:    0,
		rear:     0,
		used:     0,
		data:     make([]int, k, k),
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.data[this.rear] = value
	this.rear = (this.rear + 1) % this.capacity
	this.used++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.front = (this.front + 1) % this.capacity
	this.used--
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	res := this.data[this.front]
	return res
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	res := (this.rear - 1 + this.capacity) % this.capacity
	return this.data[res]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.used == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.used == this.capacity
}

//打印循环队列的所有元素
func (this *MyCircularQueue) List() {
	if this.IsEmpty() {
		return
	}
	var list []int
	p := this.front
	for i := 1; i <= this.used; i++ {
		list = append(list, this.data[p])
		p = (p + 1) % this.capacity
	}
	fmt.Println(list)
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
