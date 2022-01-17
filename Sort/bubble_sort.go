package Sort

func BubbleSort(list []int) {
	n := len(list)
	didSwap := false
	for i := n-1; i>0; i-- {
		for j:=0; j<i; j++ {
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

