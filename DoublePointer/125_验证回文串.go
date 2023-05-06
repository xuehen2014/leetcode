package DoublePointer

import "strings"

/*
 * https://leetcode.cn/problems/valid-palindrome/
 */


func isPalindrome(s string) bool {
	if s == "" {
		return true
	}
	var validBytes []byte
	for _, v := range []byte(s) {
		if isalnum(v) {
			validBytes = append(validBytes, v)
		}
	}

	if len(validBytes) == 0 {
		return true
	}

	left := 0
	right := len(validBytes)-1
	for left <= right {
		v1 := strings.ToLower(string(validBytes[left]))
		v2 := strings.ToLower(string(validBytes[right]))
		if v1 != v2 {
			return false
		}
		left++
		right--
	}
	return true
}

func isalnum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}