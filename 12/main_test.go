package main

import (
	"reflect"
	"testing"
)

func Test_createSet(t *testing.T) {
	type args struct {
		arr []string
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{
			"empty",
			args{arr: []string{}},
			map[string]int{},
		},
		{
			"not empty",
			args{arr: []string{"cat", "dog", "cat"}},
			map[string]int{"cat": 2, "dog": 1},
		},
		{
			"not empty",
			args{arr: []string{"cat", "dog", "cat", "cat"}},
			map[string]int{"cat": 3, "dog": 1},
		},
		{
			"not empty",
			args{arr: []string{"cat", "dog", "cat", "dog"}},
			map[string]int{"cat": 2, "dog": 2},
		},
		{
			"not empty",
			args{arr: []string{"cat", "dog", "cat", "dog", "elephant"}},
			map[string]int{"cat": 2, "dog": 2, "elephant": 1},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := createSet(tt.args.arr); !reflect.DeepEqual(
					got, tt.want,
				) {
					t.Errorf("createSet() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func Test_createSet1(t *testing.T) {
	type args struct {
		arr []string
	}
	tests := []struct {
		name    string
		args    args
		wantRes []string
	}{
		{
			"empty",
			args{arr: []string{}},
			[]string{},
		},
		{
			"not empty",
			args{arr: []string{"cat", "dog", "cat"}},
			[]string{"cat", "dog"},
		},
		{
			"not empty",
			args{arr: []string{"cat", "dog", "cat", "cat"}},
			[]string{"cat", "dog"},
		},
		{
			"not empty",
			args{arr: []string{"cat", "dog", "cat", "dog"}},
			[]string{"cat", "dog"},
		},
		{
			"not empty",
			args{arr: []string{"cat", "dog", "cat", "dog", "elephant"}},
			[]string{"cat", "dog", "elephant"},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if gotRes := createSet1(tt.args.arr); !reflect.DeepEqual(gotRes, tt.wantRes) {
					t.Errorf("createSet1() = %v, want %v", gotRes, tt.wantRes)
				}
			},
		)
	}
}
