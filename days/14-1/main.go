package main

import (
	"fmt"
	"strings"

	"github.com/robryanx/adventofcode2021/modules/readinput"
)

func main() {
	lines := readinput.ReadStrings("inputs/14/input.txt", "\n")

	pattern := lines[0]

	replacements := make(map[string]string)
	for _, line := range lines[2:] {
		line_parts := strings.Split(line, " ")

		replacements[line_parts[0]] = line_parts[2]
	}

	pattern_build := pattern
	var build_pos int
	for step := 0; step < 10; step++ {
		build_pos = 0

		for i := 0; i < len(pattern)-1; i++ {
			check := string(pattern[i]) + string(pattern[i+1])

			if _, ok := replacements[check]; ok {
				pattern_build = pattern_build[:build_pos+1] + replacements[check] + pattern_build[build_pos+1:]
				build_pos++
			}

			build_pos++
		}

		pattern = pattern_build
	}

	frequency := make(map[string]int)

	for i := 0; i < len(pattern); i++ {
		frequency[string(pattern[i])]++
	}

	smallest_count := -1
	largest_count := 0

	for _, count := range frequency {
		if smallest_count == -1 || smallest_count > count {
			smallest_count = count
		}

		if largest_count < count {
			largest_count = count
		}
	}

	fmt.Println(largest_count - smallest_count)
}
