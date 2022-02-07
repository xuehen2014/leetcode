package Stack

/*
 *  https://leetcode-cn.com/problems/valid-parentheses/
 *
 */

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	var strStack []rune
	for _, val := range s {
		if val == '(' || val == '[' || val == '{' {
			strStack = append(strStack, val)
		}
		if val == ')' {
			if len(strStack) ==0 || strStack[len(strStack)-1] != '(' {
				return false
			}
			strStack = strStack[:len(strStack)-1]
		}
		if val == ']' {
			if len(strStack) ==0 || strStack[len(strStack)-1] != '[' {
				return false
			}
			strStack = strStack[:len(strStack)-1]
		}
		if val == '}' {
			if len(strStack) ==0 || strStack[len(strStack)-1] != '{' {
				return false
			}
			strStack = strStack[:len(strStack)-1]
		}
	}
	return len(strStack) == 0
}
