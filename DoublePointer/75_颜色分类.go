package DoublePointer

/*
 * https://leetcode-cn.com/problems/sort-colors/
 * 
 */

//三路切分
func sortColors(nums []int)  {
	length := len(nums)
	l :=0
	i :=0
	r := length-1
	for i<=r {
		if nums[i] < 1 {
			nums[i], nums[l] = nums[l], nums[i]
			i++
			l++
		} else if nums[i] == 1 {
			i++
		} else {
			nums[i], nums[r] = nums[r], nums[i]
			r--
		}
	}
}

