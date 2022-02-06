package Array

/*
 *  169 --> https://leetcode-cn.com/problems/majority-element/
 */
func majorityElement(nums []int) int {
	cnt, major := 0, 0
	for _, num := range nums{
		if cnt == 0 {
			cnt = 1
			major = num
		} else {
			if num == major {
				cnt++
			} else {
				cnt--
			}
		}
	}
	return major
}
