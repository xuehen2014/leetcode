package DoublePointer

/*
 * https://leetcode-cn.com/problems/container-with-most-water/
 */

func maxArea(height []int) int {
	length := len(height)
	l :=0
	r := length-1
	ans := 0
	for l < r {
		minHeight := min(height[l], height[r])
		width := r - l
		ans = max(minHeight * width, ans)
		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}
	return ans
}

func min(x, y int) int {
	if x>y {
		return y
	}
	return x
}
func max(x, y int) int {
	if x>y {
		return x
	}
	return y
}

