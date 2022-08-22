package main

import "testing"

func Test_binarySearch(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"empty",
			args{arr: []int{}, target: 2},
			-1,
		},
		{
			"unique",
			args{
				arr:    []int{1, 2, 3, 5, 8, 10, 12, 155, 1634, 43241},
				target: 155,
			},
			155,
		},
		{
			"unique",
			args{
				arr:    []int{1, 2, 3, 8, 10, 12, 155, 1634, 43241},
				target: 2222,
			},
			-1,
		},
		{
			"not unique",
			args{arr: []int{1, 1, 1}, target: 1},
			1,
		},
		{
			"not unique",
			args{arr: []int{1, 1, 1}, target: 2},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := binarySearch(
					tt.args.arr, tt.args.target,
				); got != tt.want {
					t.Errorf("binarySearch() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
