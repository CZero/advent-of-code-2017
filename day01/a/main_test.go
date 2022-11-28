package main

import "testing"

func TestAddRepeats(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				input: "1122",
			},
			want: 3,
		},
		{
			name: "Test 2",
			args: args{
				input: "1111",
			},
			want: 4,
		},
		{
			name: "Test 3",
			args: args{
				input: "1234",
			},
			want: 0,
		},
		{
			name: "Test 4",
			args: args{
				input: "91212129",
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddRepeats(tt.args.input); got != tt.want {
				t.Errorf("AddRepeats() = %v, want %v", got, tt.want)
			}
		})
	}
}
