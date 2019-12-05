package main

import (
	"reflect"
	"testing"
)

func Test_numToList(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name        string
		args        args
		wantListNum []int
	}{
		{
			name: "big number",
			args: args{
				num: 309128301983,
			},
			wantListNum: []int{3, 0, 9, 1, 2, 8, 3, 0, 1, 9, 8, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotListNum := numToList(tt.args.num); !reflect.DeepEqual(gotListNum, tt.wantListNum) {
				t.Errorf("numToList() = %v, want %v", gotListNum, tt.wantListNum)
			}
		})
	}
}

func Test_obeysRules(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "too many ones",
			args: args{num: 111111},
			want: false,
		},
		{
			name: "decreases",
			args: args{num: 223450},
			want: false,
		},
		{
			name: "no adjacent",
			args: args{num: 123789},
			want: false,
		},
		{
			name: "3 groups",
			args: args{num: 112233},
			want: true,
		},
		{
			name: "too many fours",
			args: args{num: 123444},
			want: false,
		},
		{
			name: "nice group",
			args: args{num: 111122},
			want: true,
		},
		{
			name: "group up front",
			args: args{num: 223333},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := obeysRules(tt.args.num); got != tt.want {
				t.Errorf("obeysRules() = %v, want %v", got, tt.want)
			}
		})
	}
}
