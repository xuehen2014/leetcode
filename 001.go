package leetcode

//冒泡排序
func BubbleSort(list []int) {
	len := len(list)
	didSwap := false

	for i := len - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
				didSwap = true
			}
		}
		if !didSwap {
			return
		}
	}
}

func BubbleSort2(list []int) {
	len := len(list)
	didSwap := false

	for i := 0; i < len-1; i++ {
		for j := 0; j < len-1-i; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
				didSwap = true
			}
		}
		if !didSwap {
			return
		}
	}
}

//选择排序
func SelectSort(list []int) {
	n := len(list)
	for i := 0; i < n-1; i++ {
		min := list[i]
		minIndex := i
		for j := i + 1; j < n; j++ {
			if list[j] < min {
				min = list[j]
				minIndex = j
			}
		}
		if i != minIndex {
			list[i], list[minIndex] = list[minIndex], list[i]
		}
	}
}
