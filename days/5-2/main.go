package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/robryanx/adventofcode2021/modules/readinput"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type path struct {
	start [2]int
	end   [2]int
}

func main() {
	lines := readinput.ReadStrings("inputs/5/input.txt", "\n")

	r_path, _ := regexp.Compile("([0-9]+),([0-9]+) -> ([0-9]+),([0-9]+)")

	var grid [1000][1000]int
	var valid_paths []path
	for _, line := range lines {
		path_matches := r_path.FindStringSubmatch(line)

		var path_int [4]int
		for i := 1; i < 5; i++ {
			val, _ := strconv.Atoi(path_matches[i])
			path_int[i-1] = int(val)
		}

		make_path := path{
			start: [2]int{path_int[0], path_int[1]},
			end:   [2]int{path_int[2], path_int[3]},
		}

		valid_paths = append(valid_paths, make_path)
	}

	for _, path := range valid_paths {
		var values int
		var x_gradient int
		var y_gradient int

		if path.end[0] > path.start[0] {
			values = path.end[0] - path.start[0]
			x_gradient = 1
		} else if path.start[0] > path.end[0] {
			values = path.start[0] - path.end[0]
			x_gradient = -1
		} else {
			x_gradient = 0
		}

		if path.end[1] > path.start[1] {
			values = path.end[1] - path.start[1]
			y_gradient = 1
		} else if path.start[1] > path.end[1] {
			values = path.start[1] - path.end[1]
			y_gradient = -1
		} else {
			y_gradient = 0
		}

		current_x := path.start[0]
		current_y := path.start[1]

		for i := 0; i <= values; i++ {
			grid[current_x][current_y]++

			current_x += x_gradient
			current_y += y_gradient
		}
	}

	//print_grid(grid)

	fmt.Println(score_grid(grid))
}

func print_grid(grid [1000][1000]int) {
	max_y := len(grid)
	max_x := len(grid[0])

	for y := 0; y < max_y; y++ {
		for x := 0; x < max_x; x++ {
			fmt.Printf("%d ", grid[x][y])
		}

		fmt.Printf("\n")
	}

	fmt.Printf("\n\n\n")
}

func score_grid(grid [1000][1000]int) int {
	score := 0

	max_y := len(grid)
	max_x := len(grid[0])

	for y := 0; y < max_y; y++ {
		for x := 0; x < max_x; x++ {
			if grid[x][y] > 1 {
				score++
			}
		}
	}

	return score
}
