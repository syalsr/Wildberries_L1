package main

import "testing"

func Test_reverse(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
	}{
		{
			"default",
			args{str: "snow dog sun"},
			"sun dog snow",
		},
		{
			"one word",
			args{str: "snow"},
			"snow",
		},
		{
			"empty",
			args{str: ""},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if gotResult := reverse(tt.args.str); gotResult != tt.wantResult {
					t.Errorf(
						"reverse() = %v, want %v", gotResult, tt.wantResult,
					)
				}
			},
		)
	}
}
