package Sort

/*
 *   https://leetcode-cn.com/problems/insertion-sort-list/
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummyHead := &ListNode{Next: head}
	lastSorted, curr := head, head.Next
	for curr != nil {
		if lastSorted.Val < curr.Val {
			lastSorted = lastSorted.Next
		} else {
			pre := dummyHead
			for pre.Next.Val < curr.Val {
				pre = pre.Next
			}
			lastSorted.Next = curr.Next
			curr.Next = pre.Next
			pre.Next = curr
		}
		curr = lastSorted.Next
	}
	return dummyHead.Next
}
