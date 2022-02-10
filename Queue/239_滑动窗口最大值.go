package Queue

/*
 *  239. 滑动窗口最大值   https://leetcode-cn.com/problems/sliding-window-maximum/
 *  单调队列 -- 单调递减队列, 求最大值, 队列首部的即是最大值
 *  
 *
 */

func maxSlidingWindow(nums []int, k int) []int {
	var res []int
	queue := NewMonotoneDecreasingQueue(k)
	for key, num := range nums {
		queue.Push(num)
		if key < k-1 {
			continue
		}
		res = append(res, queue.GetFirst())
		queue.Pop(nums[key-k+1])
	}
	return res
}

//单调队列--非严格单调队列
type MonotoneDecreasingQueue struct {
	dataArray []int
	length int
}

func NewMonotoneDecreasingQueue(k int) *MonotoneDecreasingQueue {
	return &MonotoneDecreasingQueue{
		dataArray: make([]int, 0, k),
		length: 0,
	}
}

func (t *MonotoneDecreasingQueue) Push(num int) {
	for !t.IsEmpty() && t.GetLast() < num {
		t.RemoveLast()
	}
	t.dataArray = append(t.dataArray, num)
	t.length++
}

func (t *MonotoneDecreasingQueue) Pop(num int) {
	if !t.IsEmpty() && t.GetFirst() == num {
		t.RemoveFirst()
	}
}

func (t *MonotoneDecreasingQueue) RemoveFirst() {
	t.dataArray = t.dataArray[1:]
	t.length--
}

func (t *MonotoneDecreasingQueue) RemoveLast() {
	t.dataArray = t.dataArray[:t.length-1]
	t.length--
}

func (t *MonotoneDecreasingQueue) GetFirst() int {
	return t.dataArray[0]
}

func (t *MonotoneDecreasingQueue) GetLast() int {
	return t.dataArray[t.length-1]
}

func (t *MonotoneDecreasingQueue) IsEmpty() bool {
	return t.length == 0
}