package main

import "testing"

func Test_calcFuel(t *testing.T) {
	type args struct {
		mass int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test round down uneccessary",
			args: args{
				mass: 12,
			},
			want: 2,
		},
		{
			name: "test round down",
			args: args{
				mass: 14,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcFuel(tt.args.mass); got != tt.want {
				t.Errorf("calcFuel() = %v, want %v", got, tt.want)
			}
		})
	}
}
