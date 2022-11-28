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
		checksum += GetSumEvenDivsors(line)
	}
	return checksum
}

func GetSumEvenDivsors(line string) int {
	numbers := strings.Fields(line)
	for i, n := range numbers {
		for j, m := range numbers {
			if i == j {
				continue
			}
			if lfs.SilentAtoi(n)%lfs.SilentAtoi(m) == 0 {
				return lfs.SilentAtoi(n) / lfs.SilentAtoi(m)
			}
		}
	}
	panic("No divisor found")
}
