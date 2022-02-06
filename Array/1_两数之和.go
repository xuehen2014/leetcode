package Array

/*
 * https://leetcode-cn.com/problems/two-sum/
 */

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for k, v := range nums {
		otherNum := target - v
		if otherKey, ok := numMap[otherNum]; ok {
			return []int{k, otherKey}
		}
		numMap[v] = k
	}
	return []int{}
}
