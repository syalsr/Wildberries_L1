package main

import (
	"reflect"
	"testing"
)

func Test_groupTemperature(t *testing.T) {
	type args struct {
		arr []float64
	}
	tests := []struct {
		name string
		args args
		want map[int64][]float64
	}{
		{
			"empty",
			args{
				arr: []float64{},
			},
			map[int64][]float64{},
		},
		{
			"not empty",
			args{
				arr: []float64{
					1.0, 2.0, 3.5, 10.9, 11.3, 13.0, 43.6, 41.0, 49.0,
				},
			},
			map[int64][]float64{
				0: {1.0, 2.0, 3.5}, 10: {10.9, 11.3, 13.0},
				40: {43.6, 41.0, 49.0},
			},
		},
		{
			"not empty",
			args{
				arr: []float64{
					-1.0, -2.0, -3.5, -10.9, -11.3, -13.0, -43.6, -41.0, -49.0,
				},
			},
			map[int64][]float64{
				0: {-1.0, -2.0, -3.5}, -10: {-10.9, -11.3, -13.0},
				-40: {-43.6, -41.0, -49.0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := groupTemperature(tt.args.arr); !reflect.DeepEqual(
					got, tt.want,
				) {
					t.Errorf("groupTemperature() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
