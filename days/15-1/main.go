package main

import (
	"fmt"

	"github.com/robryanx/adventofcode2021/modules/grid"
)

func main() {
	build_grid := grid.Grid("inputs/15/input.txt")

	weights := make([][]uint16, len(build_grid))
	for i := 0; i < len(weights); i++ {
		weights[i] = make([]uint16, len(build_grid[0]))
	}

	for y := 0; y < len(build_grid); y++ {
		for x := 0; x < len(build_grid[0]); x++ {
			if x != 0 || y != 0 {
				test_weight := uint16(0)
				if x-1 >= 0 && weights[y][x-1] != 0 && (test_weight == 0 || weights[y][x-1] < test_weight) {
					test_weight = weights[y][x-1]
				}

				if x+1 < len(build_grid[0]) && weights[y][x+1] != 0 && (test_weight == 0 || weights[y][x+1] < test_weight) {
					test_weight = weights[y][x+1]
				}

				if y-1 >= 0 && weights[y-1][x] != 0 && (test_weight == 0 || weights[y-1][x] < test_weight) {
					test_weight = weights[y-1][x]
				}

				if y+1 < len(build_grid) && weights[y+1][x] != 0 && (test_weight == 0 || weights[y+1][x] < test_weight) {
					test_weight = weights[y+1][x]
				}

				weights[y][x] = uint16(build_grid[y][x]) + test_weight
			}
		}
	}

	fmt.Println(weights[len(build_grid)-1][len(build_grid[0])-1])
}
