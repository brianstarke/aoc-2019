package main

import "testing"

func Test_getXYHash(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			name: "hash is always the same",
			args: args{
				x: 12387,
				y: -12399,
			},
			want: 2804302117,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getXYHash(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("getXYHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
