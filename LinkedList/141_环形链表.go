package LinkedList

/*
 *  https://leetcode-cn.com/problems/linked-list-cycle/
 */

//快慢指针
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

//哈希表-记录节点是否被访问过
func hasCycle2(head *ListNode) bool {
	if head == nil {
		return false
	}
	seen := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := seen[head]; ok {
			return true
		}
		seen[head] = struct{}{}
		head = head.Next
	}
	return false
}
