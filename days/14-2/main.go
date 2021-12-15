package main

import (
	//"fmt"
	"adventofcode/2021/modules/readinput"
	"fmt"
	"strings"
)

func main() {
    lines := readinput.ReadStrings("inputs/14/input.txt", "\n")

    pattern := []byte(lines[0])

    replacements := make(map[string]byte);
    for _, line := range lines[2:] {
        line_parts := strings.Split(line, " ")

        replacements[line_parts[0]] = line_parts[2][0]
    }

    var build_pos int
    var check string
    for step:=0; step<40; step++ {
        fmt.Printf("Step %d\n", (step+1));

        build_pos = 0
        inserts := make(map[int]byte);
        for i:=0; i<len(pattern)-1; i++ {
            check = string([]byte{pattern[i], pattern[i+1]})

            if _, ok := replacements[check]; ok {
                inserts[build_pos+1] = replacements[check]

                build_pos++
            }

            build_pos++;
        }

        fmt.Println("inserts done")

        build_string := make([]byte, build_pos+1)
        org_pos := 0
        for i:=0; i<build_pos; i++ {
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

    for i:=0; i<len(pattern); i++ {
        frequency[string(pattern[i])]++
    }

    fmt.Println(frequency)

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