package Sort

//gotests -only InsertSort -w inser_sort_test.go .
func InsertSort(list []int) {
	n := len(list)

	for i := 1; i <= n-1; i++ {
		deal := list[i] // 待排序的数
		j := i - 1      // 待排序的数左边的第一个数的位置
		// 如果第一次比较，比左边的已排好序的第一个数小，那么进入处理
		if deal < list[j] {
			// 一直往左边找，比待排序大的数都往后挪，腾空位给待排序插入
			for ; j >= 0 && deal < list[j]; j-- {
				list[j+1] = list[j]
			}
		}
		list[j+1] = deal
	}
}

func InsertSortTwo(list []int) {
	n := len(list)

	for i := 1; i <= n-1; i++ {
		deal := list[i]
		if deal < list[i-1] {
			for j := i - 1; j >= 0 && deal < list[j]; j-- {
				list[j+1], list[j] = list[j], list[j+1]
			}
		}
	}
}
