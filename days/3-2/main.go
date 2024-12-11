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

	values := make([][]uint16, len(numbers))
	for i := range values {
		values[i] = make([]uint16, len(numbers[0]))
	}

	for row, number := range numbers {
		bits := strings.Split(number, "")

		for i, rawbit := range bits {
			bit, err := strconv.Atoi(rawbit)
			check(err)

			values[row][i] = uint16(bit)
		}
	}

	oxygen_generator := bit_array_to_int(get_value(1, values))
	co2_scrubber := bit_array_to_int(get_value(0, values))

	fmt.Println(oxygen_generator * co2_scrubber)
}

func get_value(value_type int, values [][]uint16) []uint16 {
	for i := 0; i < len(values[0]); i++ {
		common_bit := common_bits(uint16(i), values)
		if value_type == 0 {
			common_bit = 1 - common_bit
		}

		values = find_bits(uint16(i), common_bit, values)

		if len(values) == 1 {
			break
		}
	}

	return values[0]
}

func common_bits(postion uint16, values [][]uint16) uint16 {
	var count uint16
	for i := 0; i < len(values); i++ {
		count += values[i][postion]
	}

	if count > uint16((len(values)-1)/2) {
		return 1
	}

	return 0
}

func find_bits(postion uint16, value uint16, values [][]uint16) [][]uint16 {
	var valid_values [][]uint16

	for i := 0; i < len(values); i++ {
		if values[i][postion] == value {
			valid_values = append(valid_values, values[i])
		}
	}

	return valid_values
}

func return_row(values [][]uint16, index int) []uint16 {
	row := make([]uint16, len(values))

	for i := 0; i < len(values); i++ {
		row = append(row, values[i][index])
	}

	return row
}

func bit_array_to_int(bit_arr []uint16) int {
	build_string := ""
	for i := 0; i < len(bit_arr); i++ {
		build_string += strconv.Itoa(int(bit_arr[i]))
	}

	fmt.Println(bit_arr)
	fmt.Println(build_string)

	parsed_int, _ := strconv.ParseInt(build_string, 2, 64)

	return int(parsed_int)
}
