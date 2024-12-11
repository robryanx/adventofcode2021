package main

import (
	"fmt"
	//"sort"
	"github.com/robryanx/adventofcode2021/modules/readinput"
)

func main() {
	positions := readinput.ReadInts("inputs/7/input.txt", ",")

	// could also potentially just use the median value
	// sort.Ints(positions)
	// fmt.Println(positions[int(len(positions)/2)])

	var max_position int
	for i := 0; i < len(positions); i++ {
		if positions[i] > max_position {
			max_position = positions[i]
		}
	}

	var best_run int
	var best_run_position int
	var current_run int

	for i := 0; i < max_position; i++ {
		current_run = 0
		for j := 0; j < len(positions); j++ {
			if positions[j] > i {
				current_run += positions[j] - i
			} else {
				current_run += i - positions[j]
			}
		}

		if best_run == 0 || current_run < best_run {
			best_run = current_run
			best_run_position = i
		}
	}

	fmt.Println(best_run)
	fmt.Println(best_run_position)
}
