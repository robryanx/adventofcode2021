package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/robryanx/adventofcode2021/modules/readinput"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	commands := readinput.ReadStrings("inputs/2/input.txt", "\n")

	aim := 0
	x := 0
	y := 0
	for _, command := range commands {
		parts := strings.Split(command, " ")

		distance, err := strconv.Atoi(parts[1])
		check(err)

		switch parts[0] {
		case "up":
			aim -= distance
		case "down":
			aim += distance
		case "forward":
			x += distance
			y += (aim * distance)
		}
	}

	fmt.Printf("%d\n", x*y)
}
