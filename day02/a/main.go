// https://adventofcode.com/2017/day/2
// --- Day 2: Corruption Checksum ---

package main

import (
	"fmt"
	"strings"

	"github.com/CZero/gofuncy/lfs"
)

func main() {
	input, err := lfs.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d\n", GetChecksum(input))
}

// GetChecksum returns the checksum of the spreadsheet by adding the biggest difference for each line
func GetChecksum(input []string) int {
	var checksum int
	for _, line := range input {
		checksum += GetDifferenceBiggestAndSmallest(line)
	}
	return checksum
}

func GetDifferenceBiggestAndSmallest(input string) int {
	var biggest, smallest int
	numbers := strings.Fields(input)
	for _, number := range numbers {
		n := lfs.SilentAtoi(number)
		if n > biggest {
			biggest = n
		}
		if n < smallest {
			smallest = n
		} else if smallest == 0 {
			smallest = n
		}
	}
	return biggest - smallest
}
