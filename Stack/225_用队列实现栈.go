package Stack

/*
 * https://leetcode-cn.com/problems/implement-stack-using-queues/
 */

type MyStack struct {
	dataQueue []int
}


func Constructor() MyStack {
	return MyStack{
		dataQueue: []int{},
	}
}


func (this *MyStack) Push(x int)  {
	length := len(this.dataQueue)
	this.dataQueue = append(this.dataQueue, x)
	for ;length>0;length-- {
		val := this.dataQueue[0]
		this.dataQueue = append(this.dataQueue, val)
		this.dataQueue = this.dataQueue[1:]
	}
}


func (this *MyStack) Pop() int {
	val := this.dataQueue[0]
	this.dataQueue = this.dataQueue[1:]
	return val
}


func (this *MyStack) Top() int {
	return this.dataQueue[0]
}


func (this *MyStack) Empty() bool {
	return len(this.dataQueue) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */