package LinkedList

/*
 *  https://leetcode-cn.com/problems/palindrome-linked-list/
 *
 */

//双指针方法, 首尾对比
func isPalindrome(head *ListNode) bool {
	var vals []int
	for ; head != nil; head = head.Next {
		vals = append(vals, head.Val)
	}
	n := len(vals)
	left :=0
	right := n-1
	for left<=right {
		if vals[left] != vals[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func isPalindrome2(head *ListNode) bool {
	if head == nil {
		return true
	}
	var vals []int
	for cur:=head; cur!= nil; cur=cur.Next {
		vals = append(vals, cur.Val)
	}
	for k, v := range vals[:len(vals)/2] {
		if v != vals[len(vals)-1-k] {
			return false
		}
	}
	return true
}