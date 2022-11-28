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
				input: "1212",
			},
			want: 6,
		},
		{
			name: "Test 2",
			args: args{
				input: "1221",
			},
			want: 0,
		},
		{
			name: "Test 3",
			args: args{
				input: "123425",
			},
			want: 4,
		},
		{
			name: "Test 4",
			args: args{
				input: "123123",
			},
			want: 12,
		},
		{
			name: "Test 5",
			args: args{
				input: "12131415",
			},
			want: 4,
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
