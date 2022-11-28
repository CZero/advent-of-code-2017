// https://adventofcode.com/2017/day/2
// --- Day 2: Corruption Checksum ---

package main

import "testing"

func TestGetDifferenceBiggestAndSmallest(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example",
			args: args{
				input: "5 1 9 5",
			},
			want: 8,
		},
		{
			name: "example 2",
			args: args{
				input: "7 5 3",
			},
			want: 4,
		},
		{
			name: "example 3",
			args: args{
				input: "2 4 6 8",
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDifferenceBiggestAndSmallest(tt.args.input); got != tt.want {
				t.Errorf("GetDifferenceBiggestAndSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}
