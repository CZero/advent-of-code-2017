package main

import (
	"fmt"

	"github.com/CZero/gofuncy/lfs"
)

func main() {
	input, err := lfs.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(AddRepeats(input[0]))
}

// AddRepeats adds a character that is repeated on the next position to a sum. It's circular, so the first character is added to the end.
func AddRepeats(input string) int {
	sum := 0
	for i, char := range input {
		nextChar := rune(input[(i+1)%len(input)])
		if char == nextChar {
			sum += lfs.SilentAtoi(string(char))
		}
	}
	return sum
}
