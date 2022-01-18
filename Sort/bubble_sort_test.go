package Sort

import (
	"fmt"
	"testing"
)

//go test -v -run TestBubbleSort bubble_sort.go bubble_sort_test.go
func TestBubbleSort(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "one", args: struct{ list []int }{list: []int{5, 19, 1, 89, 8, 14, 16, 49, 25, 4, 6, 13}}},
		{name:"two", args: struct{ list []int }{list: []int{7,6,1,2,3,4,5}}},
		{name:"three", args: struct{ list []int }{list: []int{0, 2, 9, 8, 7, 10, 8, 100}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MergeSort(tt.args.list)
		})
		fmt.Println(tt.args.list)
	}
}
