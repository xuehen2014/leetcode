package LinkedList

/*
 *  https://leetcode-cn.com/problems/reverse-linked-list/
 */

//三指针做法
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre *ListNode
	cur := head
	for cur != nil {
		nextNode := cur.Next
		cur.Next = pre
		pre = cur
		cur = nextNode
	}
	return pre
}


//递归
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	last := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

//新建链表用头插法
func reverseList3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{}
	for head != nil {
		nextNode := head.Next
		head.Next = dummy.Next
		dummy.Next = head
		head = nextNode
	}
	return dummy.Next
}
