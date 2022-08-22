package main

import (
	"reflect"
	"testing"
)

func Test_deleteElemFromSlice(t *testing.T) {
	type args struct {
		arr []int64
		idx uint
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			"delete",
			args{arr: []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, idx: 4},
			[]int64{1, 2, 3, 4, 6, 7, 8, 9, 10},
		},
		{
			"delete first",
			args{arr: []int64{1, 2, 3}, idx: 0},
			[]int64{2, 3},
		},
		{
			"delete last",
			args{arr: []int64{1, 2, 3, 4, 1}, idx: 4},
			[]int64{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := deleteElemFromSlice(
					tt.args.arr, tt.args.idx,
				); !reflect.DeepEqual(got, tt.want) {
					t.Errorf(
						"deleteElemFromSlice() = %v, want %v", got, tt.want,
					)
				}
			},
		)
	}
}
