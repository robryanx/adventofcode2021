package main

import (
	"fmt"

	"github.com/robryanx/adventofcode2021/modules/readinput"
)

func main() {
	fish := readinput.ReadInts("inputs/6/input.txt", ",")

	var fish_list [9]int
	var new_fish int

	for i := 0; i < len(fish); i++ {
		fish_list[fish[i]]++
	}

	days := 256
	for current_day := 0; current_day < days; current_day++ {
		new_fish = fish_list[0]

		for j := 1; j < 9; j++ {
			fish_list[j-1] = fish_list[j]
		}

		fish_list[6] += new_fish
		fish_list[8] = new_fish
	}

	var total_fish int

	for i := 0; i < 9; i++ {
		total_fish += fish_list[i]
	}

	fmt.Println(total_fish)
}
