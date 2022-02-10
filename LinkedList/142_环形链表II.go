package LinkedList

/*
 *  https://leetcode-cn.com/problems/linked-list-cycle-ii/
 *
 */

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next== nil {
		return nil
	}
	isFlag := false
	slow, fast := head, head
	for fast !=nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			isFlag = true
			break
		}
	}
	if !isFlag {
		return nil
	}
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
