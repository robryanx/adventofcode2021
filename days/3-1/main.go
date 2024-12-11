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
	numbers := readinput.ReadStrings("inputs/3/input.txt", "\n")

	values := make([][]uint16, len(numbers[0]))
	for i := range values {
		values[i] = make([]uint16, len(numbers))
	}

	for row, number := range numbers {
		bits := strings.Split(number, "")

		for i, rawbit := range bits {
			bit, err := strconv.Atoi(rawbit)
			check(err)

			values[i][row] = uint16(bit)
		}
	}

	gamma := ""
	epsilon := ""
	count := uint16(0)
	for i := 0; i < len(numbers[0]); i++ {
		count = 0
		for j := 0; j < len(numbers); j++ {
			count += values[i][j]
		}

		if count > (uint16(len(numbers) / 2)) {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	gamma_converted, _ := strconv.ParseInt(gamma, 2, 64)
	epsilon_coverted, _ := strconv.ParseInt(epsilon, 2, 64)

	fmt.Printf("%d\n", (gamma_converted * epsilon_coverted))
}
