package main

import "testing"

func Test_reverseStr(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
	}{
		{
			"empty",
			args{str: ""},
			"",
		},
		{
			"space",
			args{str: " "},
			" ",
		},
		{
			"str",
			args{str: "abcde"},
			"edcba",
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if gotResult := reverseStr(tt.args.str); gotResult != tt.wantResult {
					t.Errorf(
						"reverseStr() = %v, want %v", gotResult, tt.wantResult,
					)
				}
			},
		)
	}
}

func Test_yetAnotherReverseStr(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"empty",
			args{str: ""},
			"",
		},
		{
			"space",
			args{str: " "},
			" ",
		},
		{
			"str",
			args{str: "abcde"},
			"edcba",
		},
		{
			"str",
			args{str: "abcd"},
			"dcba",
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := yetAnotherReverseStr(tt.args.str); got != tt.want {
					t.Errorf(
						"yetAnotherReverseStr() = %v, want %v", got, tt.want,
					)
				}
			},
		)
	}
}
