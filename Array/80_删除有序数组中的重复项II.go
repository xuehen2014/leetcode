package Array

/*
 * https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array-ii/
 */

func removeDuplicates2(nums []int) int {
	length := len(nums)
	if length <= 2 {
		return length
	}
	slow := 2
	fast := 2
	for ;fast<length; {
		if nums[fast] == nums[slow-2] {
			fast++
		} else {
			nums[slow] = nums[fast]
			slow++
			fast++
		}
	}
	return slow
}
