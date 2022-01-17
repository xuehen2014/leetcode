package Sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
	}{
		struct {
			name string
			args args
		}{name: "one", args: struct{ list []int }{list: []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BubbleSort(tt.args.list)
		})
		fmt.Println(tt.args.list)
	}
}
