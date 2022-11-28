// https://adventofcode.com/2017/day/2
// --- Day 2: Corruption Checksum ---

package main

import "testing"

func TestGetSumEvenDivsors(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				line: "5 9 2 8",
			},
			want: 4,
		},
		{
			name: "Test 2",
			args: args{
				line: "9 4 7 3",
			},
			want: 3,
		},
		{
			name: "Test 3",
			args: args{
				line: "3 8 6 5",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSumEvenDivsors(tt.args.line); got != tt.want {
				t.Errorf("GetSumEvenDivsors() = %v, want %v", got, tt.want)
			}
		})
	}
}
