package main

import (
	"reflect"
	"testing"
)

func TestArray(t *testing.T) {
	type args struct {
		arr1 []int
		arr2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"empty",
			args{arr1: []int{}, arr2: []int{}},
			[]int{},
		},
		{
			"one of empty",
			args{arr1: []int{}, arr2: []int{2, 1, 4}},
			[]int{},
		},
		{
			"one of empty",
			args{arr1: []int{2, 1, 4}, arr2: []int{}},
			[]int{},
		},
		{
			"one bigger than other",
			args{arr1: []int{2, 1, 4, 5, 7}, arr2: []int{10, 1, 3, 4}},
			[]int{1, 4},
		},
		{
			"one bigger than other",
			args{arr1: []int{2, 1, 4}, arr2: []int{1, 2, 6, 10, 32}},
			[]int{1, 2},
		},
		{
			"equal",
			args{arr1: []int{2, 1, 4}, arr2: []int{4, 2, 1}},
			[]int{1, 2, 4},
		},
		{
			"same length",
			args{arr1: []int{2, 1, 4, 5, 3}, arr2: []int{2, 5, 8, 9, 1}},
			[]int{1, 2, 5},
		},
		{
			"unique",
			args{arr1: []int{4, 3, 2, 1}, arr2: []int{10, 20, 30, 40}},
			[]int{},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := Array(
					tt.args.arr1, tt.args.arr2,
				); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Array() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestMap(t *testing.T) {
	type args struct {
		arr1 []int
		arr2 []int
	}
	tests := []struct {
		name string
		args args
		want map[int]byte
	}{
		{
			"empty",
			args{arr1: []int{}, arr2: []int{}},
			map[int]byte{},
		},
		{
			"one of empty",
			args{arr1: []int{}, arr2: []int{2, 1, 4}},
			map[int]byte{},
		},
		{
			"one of empty",
			args{arr1: []int{2, 1, 4}, arr2: []int{}},
			map[int]byte{},
		},
		{
			"one bigger than other",
			args{arr1: []int{2, 1, 4, 5, 7}, arr2: []int{10, 1, 3, 4}},
			map[int]byte{1: 2, 4: 2},
		},
		{
			"one bigger than other",
			args{arr1: []int{2, 1, 4}, arr2: []int{1, 2, 6, 10, 32}},
			map[int]byte{2: 2, 1: 2},
		},
		{
			"equal",
			args{arr1: []int{2, 1, 4}, arr2: []int{4, 2, 1}},
			map[int]byte{2: 2, 1: 2, 4: 2},
		},
		{
			"same length",
			args{arr1: []int{2, 1, 4, 5, 3}, arr2: []int{2, 5, 8, 9, 1}},
			map[int]byte{2: 2, 1: 2, 5: 2},
		},
		{
			"unique",
			args{arr1: []int{4, 3, 2, 1}, arr2: []int{10, 20, 30, 40}},
			map[int]byte{},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := Map(
					tt.args.arr1, tt.args.arr2,
				); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Map() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
