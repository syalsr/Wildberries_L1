package main

import "testing"

func Test_checkUnqueString(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"upper and lowe case",
			args{str: "Aa"},
			false,
		},
		{
			"upper case",
			args{str: "AA"},
			false,
		},
		{
			"lower case",
			args{str: "aa"},
			false,
		},
		{
			"upper and lowe case",
			args{str: "AbcFg"},
			true,
		},
		{
			"upper case",
			args{str: "ABFE"},
			true,
		},
		{
			"lower case",
			args{str: "abcf"},
			true,
		},
		{
			"empty str",
			args{str: ""},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkUnqueString(tt.args.str); got != tt.want {
				t.Errorf("checkUnqueString() = %v, want %v", got, tt.want)
			}
		})
	}
}
