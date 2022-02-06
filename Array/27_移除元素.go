package Array


/*
 * 27---> https://leetcode-cn.com/problems/remove-element/
 */

func removeElement(nums []int, val int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	slow := 0
	fast := 0
	for ; fast<length;  {
		if nums[fast] == val {
			fast++
		} else {
			nums[slow] = nums[fast]
			slow++
			fast++
		}
	}
	return slow
}


func removeElement2(nums []int, val int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	left := 0
	right := length-1
	for left <= right {
		if nums[left]== val {
			nums[left]= nums[right]
			right--
		} else {
			left++
		}
	}
	return left
}