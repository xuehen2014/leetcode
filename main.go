package main

import "github.com/xuehen2014/leetcode/LinkedList/singly_linkedlist"

func main()  {
	//list := []int{1,2,3,5}
	//Sort.BubbleSort(list)
	//fmt.Println(list)
	//circularQueue := Queue.Constructor(5)
	//circularQueue.EnQueue(1)
	//circularQueue.EnQueue(2)
	//circularQueue.EnQueue(3)
	//circularQueue.EnQueue(4)
	//circularQueue.EnQueue(5)
	//circularQueue.List()
	//circularQueue.DeQueue()
	//circularQueue.DeQueue()
	//circularQueue.List()
	//circularQueue.EnQueue(6)
	//circularQueue.List()
	//circularQueue.DeQueue()
	//circularQueue.List()
	singlyLinkedList := singly_linkedlist.CreateList()
	singlyLinkedList.Add(1)
	singlyLinkedList.Add(0)
	singlyLinkedList.Append(2)
	singlyLinkedList.Scan()
}
