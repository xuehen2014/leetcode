package Greedy

import "sort"

/*
 *  https://leetcode-cn.com/problems/assign-cookies/
 */
//贪心
func findContentChildren(g []int, s []int) int {
	glength := len(g)
	slength := len(s)
	sort.Ints(g)
	sort.Ints(s)
	gi:=0
	si:=0
	ans := 0
	for gi<glength && si<slength {
		if s[si]>=g[gi] {
			ans++
			si++
			gi++
		} else {
			si++
		}
	}
	return ans
}