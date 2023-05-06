package DoublePointer

/*
 * https://leetcode.cn/problems/reverse-string/
 */

func reverseString(s []byte)  {
	length := len(s)
	left :=0
	right := length-1
	for left<right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}