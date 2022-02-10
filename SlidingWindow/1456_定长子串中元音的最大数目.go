package SlidingWindow

/*
 *  https://leetcode-cn.com/problems/maximum-number-of-vowels-in-a-substring-of-given-length/
 *  滑动窗口
 */


func maxVowels(s string, k int) int {
	if s == "" {
		return 0
	}
	strBytes := []byte(s)
	length := len(strBytes)
	totalNum := 0
	maxNum := 0
	right := 0
	for ;right <= length-1;right++ {
		currentByte := strBytes[right]
		if isFlag(currentByte) {
			totalNum++
		}
		if right < k-1 { //还未达到滑动窗口的大小
			continue
		}
		if totalNum > maxNum {
			maxNum = totalNum
		}
		leftByte := strBytes[right-k+1]
		if isFlag(leftByte) {
			totalNum--
		}
	}
	return maxNum
}

func isFlag(b byte) bool {
	if b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u' {
		return true
	}
	return false
}


func maxVowels2(s string, k int) int {
	if s == "" {
		return 0
	}
	strBytes := []byte(s)
	length := len(strBytes)
	left, right := 0, 0
	totalNum := 0
	maxNum := 0
	for right <= length-1 {
		currentByte := strBytes[right]
		if right < k {
			if isFlag(currentByte) {
				totalNum++
			}
		} else {
			addByte := currentByte
			removeByte := strBytes[left]
			left++
			if isFlag(addByte) {
				totalNum++
			}
			if isFlag(removeByte) {
				totalNum--
			}
		}
		if totalNum > maxNum {
			maxNum = totalNum
		}
		right++
	}
	return maxNum
}

