package Stack

/*
 * 单调栈
 * https://leetcode-cn.com/problems/daily-temperatures/
 */

func dailyTemperatures(T []int) []int {
	res := make([]int, len(T))
	var toCheck []int    //用数组模拟单调栈
	for i,t := range T {
		for len(toCheck) > 0 && t > T[toCheck[len(toCheck)-1]] {
			idx := toCheck[len(toCheck) - 1]
			res[idx] = i-idx
			toCheck = toCheck[:len(toCheck)-1]
		}
		toCheck = append(toCheck, i)
	}
	return res
}

func dailyTemperatures2(T []int) []int {
	res := make([]int, len(T))
	var toCheck []int
	for i:=0; i<len(T); i++ {
		if len(toCheck) == 0 || T[i] <= T[toCheck[len(toCheck)-1]] {
			toCheck = append(toCheck, i)
		} else {
			for len(toCheck) > 0 && T[i] > T[toCheck[len(toCheck)-1]] {
				idx := toCheck[len(toCheck) -1 ]
				res[idx] = i - idx
				toCheck = toCheck[:len(toCheck)-1]
			}
			toCheck = append(toCheck, i)
		}
	}
	return res
}


