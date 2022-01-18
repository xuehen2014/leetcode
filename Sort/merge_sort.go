package Sort

/*
 * 归并排序
 */
//自顶向下  从上往下进行递归，直到切分的小数组无法切分了，然后不断地对这些有序数组进行合并
func MergeSort(list []int) {
	n := len(list)
	if n == 0 || n == 1 {
		return
	}
	tmp := make([]int, n, n)
	//开始结束为 左闭右开
	mSort(list, 0, n, tmp)
}

func mSort(list []int, begin int, end int, tmp []int) {
	if begin >= end || begin+1 >= end { //如果区间为空 或者 只有一个元素直接返回
		return
	}
	//计算中间节点
	mid := begin + ((end - begin) >> 1)

	mSort(list, begin, mid, tmp)
	mSort(list, mid, end, tmp)

	i := begin
	j := mid
	to := begin

	for i < mid || j < end {
		if j >= end || (i < mid && list[i] <= list[j]) {
			tmp[to] = list[i]
			to++
			i++
		} else {
			tmp[to] = list[j]
			to++
			j++
		}
	}
	for deal := begin; deal < end; deal++ {
		list[deal] = tmp[deal]
	}
}
