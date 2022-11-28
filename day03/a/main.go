// https://adventofcode.com/2017/day/3
// Day 3: Spiral Memory
package main

import (
	"fmt"
	"math"
)

type coord struct {
	c int // column
	r int // row
}

func main() {
	var (
		grid  = make(map[int]coord, 100000)
		limit = 347991
	)
	buildGrid(grid, limit)
	fmt.Printf("Distance of %d: %d (%v)", limit, getDistance(limit, grid), grid[limit])
}

func buildGrid(grid map[int]coord, limit int) {
	pos := coord{0, 0}
	count := 1
	reps := 1
	grid[count] = pos
	for count <= limit {
		for i := 1; i <= reps; i++ { // Right
			pos.c++
			count++
			grid[count] = pos
		}
		for i := 1; i <= reps; i++ { // Up
			pos.r++
			count++
			grid[count] = pos
		}
		reps++
		for i := 1; i <= reps; i++ { // Left
			pos.c--
			count++
			grid[count] = pos
		}
		for i := 1; i <= reps; i++ { // Down
			pos.r--
			count++
			grid[count] = pos
		}
		reps++
	}

}

func getDistance(limit int, grid map[int]coord) int {
	pos := grid[limit]
	return int(math.Abs(float64(pos.c))) + int(math.Abs(float64(pos.r)))
}
