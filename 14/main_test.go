package main

import "testing"

func Test_checkVar(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"int",
			args{v: 12},
			"int",
		},
		{
			"bool",
			args{v: true},
			"bool",
		},
		{
			"string",
			args{v: "str"},
			"string",
		},
		{
			"channel int",
			args{v: make(chan int)},
			"chan int",
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := checkVar(tt.args.v); got != tt.want {
					t.Errorf("checkVar() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func Test_yetAnotherCheckVar(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"int",
			args{v: 12},
			"int",
		},
		{
			"bool",
			args{v: true},
			"bool",
		},
		{
			"string",
			args{v: "str"},
			"string",
		},
		{
			"channel int",
			args{v: make(chan int)},
			"chan int",
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := yetAnotherCheckVar(tt.args.v); got != tt.want {
					t.Errorf("yetAnotherCheckVar() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func Test_yetYetAnotherCheckVar(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"int",
			args{v: 12},
			"int",
		},
		{
			"bool",
			args{v: true},
			"bool",
		},
		{
			"string",
			args{v: "str"},
			"string",
		},
		{
			"channel int",
			args{v: make(chan int)},
			"chan int",
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := yetYetAnotherCheckVar(tt.args.v); got != tt.want {
					t.Errorf(
						"yetYetAnotherCheckVar() = %v, want %v", got, tt.want,
					)
				}
			},
		)
	}
}

func Test_yetYetYetAnotherCheckVar(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"int",
			args{v: 12},
			"int",
		},
		{
			"bool",
			args{v: true},
			"bool",
		},
		{
			"string",
			args{v: "str"},
			"string",
		},
		{
			"channel int",
			args{v: make(chan int)},
			"chan",
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := yetYetYetAnotherCheckVar(tt.args.v); got != tt.want {
					t.Errorf(
						"yetYetYetAnotherCheckVar() = %v, want %v", got,
						tt.want,
					)
				}
			},
		)
	}
}
