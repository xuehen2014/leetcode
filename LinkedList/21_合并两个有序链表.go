package LinkedList

/*
 * https://leetcode-cn.com/problems/merge-two-sorted-lists/
 */

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2==nil {
		return nil
	}
	dummy := &ListNode{}
	cur := dummy
	for l1 != nil || l2 != nil {
		if l2 == nil || (l2!=nil && l1 != nil && l1.Val<l2.Val) {
			cur.Next = &ListNode{Val: l1.Val}
			l1 = l1.Next
		} else {
			cur.Next = &ListNode{Val: l2.Val}
			l2 = l2.Next
		}
		cur = cur.Next
	}
	return dummy.Next
}


func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2==nil {
		return nil
	}
	dummy := &ListNode{}
	cur := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = &ListNode{Val: l1.Val}
			l1 = l1.Next
		} else {
			cur.Next = &ListNode{Val: l2.Val}
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return dummy.Next
}
