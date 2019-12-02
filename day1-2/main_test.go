package main

import (
	"testing"
)

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
			name: "no additional fuel",
			args: args{
				mass: 12,
			},
			want: 2,
		},
		{
			name: "much additional fuel",
			args: args{
				mass: 1969,
			},
			want: 966,
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

func Test_calcAdditionalFuel(t *testing.T) {
	type args struct {
		fuel    int
		fuelSum int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name: "test 2",
			args: args{
				fuel:    2,
				fuelSum: 0,
			},
			want:  0,
			want1: 0,
		},
		{
			name: "test 654",
			args: args{
				fuel:    654,
				fuelSum: 0,
			},
			want:  0,
			want1: 966 - 654,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := calcAdditionalFuel(tt.args.fuel, tt.args.fuelSum)
			if got != tt.want {
				t.Errorf("calcAdditionalFuel() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("calcAdditionalFuel() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
