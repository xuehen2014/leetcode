package Sort

//快速排序
func QuickSort(list []int) {
	if len(list) <= 1 {
		return
	}
	qSort(list, 0, len(list))
}

func qSort(list []int, begin int, end int) { //左闭右开区间
	if begin >= end || begin+1 >= end { //区间为空或者区间只有一个元素
		return
	}
	mid := begin + ((end - begin) >> 1)
	x := list[mid]

	l := begin
	i := begin
	r := end - 1
	//三路切分
	for i <= r {
		if list[i] < x {
			list[l], list[i] = list[i], list[l]
			l++
			i++
		} else if list[i] == x {
			i++
		} else {
			list[i], list[r] = list[r], list[i]
			r--
		}
	}
	qSort(list, begin, l)
	qSort(list, i, end)
}
