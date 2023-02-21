package Sort

import (
	"fmt"
	"testing"
)

//go test -v -run  TestInsertSort insert_sort.go insert_sort_test.go
func TestInsertSort(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "a", args: args{list: []int{9, 8, 7, 6, 4, 0}}},
		{name: "b", args: args{list: []int{1, 2, 3, 5, 9, 8, 7, 6, 4, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InsertSort(tt.args.list)
			fmt.Println(tt.args.list)
		})
	}
}

func TestInsertSortTwo(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InsertSortTwo(tt.args.list)
		})
	}
}
