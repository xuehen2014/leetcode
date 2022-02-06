package Array

/*
 * https://leetcode-cn.com/problems/merge-sorted-array/
 */

func merge(nums1 []int, m int, nums2 []int, n int)  {
	tail := m + n -1
	p1 := m-1
	p2 := n-1
	for p1>=0 || p2 >=0 {
		if p2<0 || p1>=0 && nums1[p1] >= nums2[p2] {
			nums1[tail] = nums1[p1]
			tail--
			p1--
		} else {
			nums1[tail] = nums2[p2]
			tail--
			p2--
		}
	}
}