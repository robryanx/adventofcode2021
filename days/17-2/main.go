package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/robryanx/adventofcode2021/modules/readinput"
)

func main() {
	target_info := readinput.ReadStrings("inputs/17/input.txt", "\n")

	regex := *regexp.MustCompile(`^target area: x=([0-9-]+)\.\.([0-9-]+), y=([0-9-]+)\.\.([0-9-]+)$`)

	target_parts := regex.FindStringSubmatch(target_info[0])

	var target [4]int
	for i := 1; i < 5; i++ {
		target[(i - 1)], _ = strconv.Atoi(target_parts[i])
	}

	start_x := target[0] / 15
	end_x := target[1]

	start_y := target[2]
	end_y := 300

	count := 0

	for x := start_x; x <= end_x; x++ {
		for y := start_y; y <= end_y; y++ {
			if do_run(target, x, y) {
				count++
			}
		}
	}

	fmt.Println(count)
}

func do_run(target [4]int, x_velocity int, y_velocity int) bool {
	x := 0
	y := 0

	for {
		x += x_velocity
		y += y_velocity

		if end_check(target, x, y) {
			return false
		}

		if rec_check(target, x, y) {
			return true
		}

		if x_velocity > 0 {
			x_velocity--
		}

		y_velocity--
	}
}

func end_check(target [4]int, x int, y int) bool {
	return y < target[2] || x > target[1]
}

func rec_check(target [4]int, x int, y int) bool {
	return x >= target[0] && y <= target[1] && y >= target[2] && y <= target[3]
}
