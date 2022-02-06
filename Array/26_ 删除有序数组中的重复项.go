package Array

/*
 * 26--> https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
 */

func removeDuplicates(nums []int) int {
	count := len(nums)
	if count<=1 {
		return count
	}
	fast := 1
	slow := 1
	for ; fast<count;  {
		if nums[fast] == nums[fast-1] {
			fast++
		} else {
			nums[slow] = nums[fast]
			slow++
			fast++
		}
	}
	return slow
}