package main

import "testing"

func Test_setBit(t *testing.T) {
	type args struct {
		num int64
		pos uint8
		b   bool
	}
	tests := []struct {
		name  string
		args  args
		wantN int64
	}{
		{
			"all one",
			args{num: 255, pos: 7, b: false},
			127,
		},
		{
			"all one",
			args{num: 255, pos: 0, b: false},
			254,
		},
		{
			"mix",
			args{num: 326, pos: 8, b: false},
			70,
		},
		{
			"pos > len(bit num)",
			args{num: 326, pos: 12, b: true},
			4422,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if gotN := setBit(
					tt.args.num, tt.args.pos, tt.args.b,
				); gotN != tt.wantN {
					t.Errorf("setBit() = %v, want %v", gotN, tt.wantN)
				}
			},
		)
	}
}
