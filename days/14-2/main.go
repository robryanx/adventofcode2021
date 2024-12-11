package main

import (
	"fmt"
	"strings"

	"github.com/robryanx/adventofcode2021/modules/readinput"
)

type pattern struct {
	chars string
	count int
}

func main() {
	lines := readinput.ReadStrings("inputs/14/input.txt", "\n")

	pattern := []byte(lines[0])

	replacements := make(map[string]byte)
	for _, line := range lines[2:] {
		line_parts := strings.Split(line, " ")

		replacements[line_parts[0]] = line_parts[2][0]
	}

	frequency := list_version(replacements, pattern, 40)

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

func add_pattern(pattern_list []*pattern, chars string, count int) []*pattern {
	for i := 0; i < len(pattern_list); i++ {
		if pattern_list[i].chars == chars {
			pattern_list[i].count += count

			return pattern_list
		}
	}

	pattern_list = append(pattern_list, &pattern{chars: chars, count: count})

	return pattern_list
}

func list_version(replacements map[string]byte, start_list []byte, steps int) map[string]int {
	var pattern_list []*pattern

	var check string
	for i := 0; i < len(start_list)-1; i++ {
		check = string([]byte{start_list[i], start_list[i+1]})

		pattern_list = append(pattern_list, &pattern{chars: check, count: 1})
	}

	for step := 0; step < steps; step++ {
		new_patterns := []*pattern{}

		for _, pattern_i := range pattern_list {
			if _, ok := replacements[pattern_i.chars]; ok {
				check = string([]byte{pattern_i.chars[0], replacements[pattern_i.chars]})

				new_patterns = add_pattern(new_patterns, check, pattern_i.count)

				check = string([]byte{replacements[pattern_i.chars], pattern_i.chars[1]})

				new_patterns = add_pattern(new_patterns, check, pattern_i.count)
			} else {
				new_patterns = add_pattern(new_patterns, pattern_i.chars, pattern_i.count)
			}
		}

		pattern_list = new_patterns
	}

	frequency := make(map[string]int)
	for i := 0; i < len(pattern_list); i++ {
		frequency[string(pattern_list[i].chars[0])] += pattern_list[i].count
	}

	frequency[string(start_list[len(start_list)-1])]++

	return frequency
}

func concat_version(replacements map[string]byte, pattern []byte, steps int) map[string]int {
	var build_pos int
	var check string
	for step := 0; step < steps; step++ {
		fmt.Printf("Step %d\n", (step + 1))

		build_pos = 0
		inserts := make(map[int]byte)
		for i := 0; i < len(pattern)-1; i++ {
			check = string([]byte{pattern[i], pattern[i+1]})

			if _, ok := replacements[check]; ok {
				inserts[build_pos+1] = replacements[check]

				build_pos++
			}

			build_pos++
		}

		fmt.Println("inserts done")

		build_string := make([]byte, build_pos+1)
		org_pos := 0
		for i := 0; i < build_pos; i++ {
			if _, ok := inserts[i]; ok {
				build_string[i] = inserts[i]
				i++
			}

			build_string[i] = pattern[org_pos]
			org_pos++
		}

		pattern = build_string
	}

	frequency := make(map[string]int)

	for i := 0; i < len(pattern); i++ {
		frequency[string(pattern[i])]++
	}

	return frequency
}
