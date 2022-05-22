package Greedy

/*
 * https://leetcode-cn.com/problems/split-a-string-in-balanced-strings/
 */

func balancedStringSplit(s string) int {
	result := 0
	length := 0
	for _, row := range s {
		if row == 'L' {
			length++
		} else {
			length--
		}
		if length == 0 {
			result++
		}
	}
	return result
}
