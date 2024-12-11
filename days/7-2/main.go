package main

import (
	"fmt"

	"github.com/robryanx/adventofcode2021/modules/readinput"
)

func main() {
	positions := readinput.ReadInts("inputs/7/input.txt", ",")

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
			var gap int

			if positions[j] > i {
				gap = positions[j] - i
			} else {
				gap = i - positions[j]
			}

			var gap_total int
			for k := 0; k < gap; k++ {
				gap_total += k + 1
			}

			// triangle numbers
			current_run += int((gap * (gap + 1)) / 2)
		}

		if best_run == 0 || current_run < best_run {
			best_run = current_run
			best_run_position = i
		}
	}

	fmt.Println(best_run)
	fmt.Println(best_run_position)
}
