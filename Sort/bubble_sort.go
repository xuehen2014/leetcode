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

func BubbleSortTwo(list []int) {
	n := len(list)
	didSwap := false
	for i:=0; i<n-1; i++{
		for j:=0; j<n-i-1;j++ {
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

