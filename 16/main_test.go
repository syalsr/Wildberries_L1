package main

import (
	"reflect"
	"testing"
)

func Test_quickSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"empty",
			args{arr: []int{}},
			[]int{},
		},
		{
			"unique elements",
			args{arr: []int{3, 5, 1, 2, 0}},
			[]int{0, 1, 2, 3, 5},
		},
		{
			"not unique elements",
			args{arr: []int{3, 5, 5, 2}},
			[]int{2, 3, 5, 5},
		},
		{
			"all unique elements",
			args{arr: []int{5, 5, 5}},
			[]int{5, 5, 5},
		},
		{
			"negative elements",
			args{arr: []int{-5, -15, -1}},
			[]int{-15, -5, -1},
		},
		{
			"positive and negative elements",
			args{arr: []int{-5, -15, -1, 3, 6}},
			[]int{-15, -5, -1, 3, 6},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := quickSort(tt.args.arr); !reflect.DeepEqual(
					got, tt.want,
				) {
					t.Errorf("quickSort() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
