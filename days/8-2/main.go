package main

import (
    "fmt"
    "strings"
    "adventofcode/2021/modules/readinput"
)

func main() {
    lines := readinput.ReadStrings("inputs/8/input.txt", "\n")

    // 0000
    // 5  1
    // 5  1
    // 5661
    // 4  2
    // 4  2
    // 3333

    var count int
    var segment_lookup [7][]string
    var segments [7]string

    for _, line := range lines {
        parts := strings.Split(line, " | ")

        input := strings.Fields(parts[0])
        output := strings.Fields(parts[1])

        for _, segment := range input {
            segment_lookup[len(segment)-1] = append(segment_lookup[len(segment)-1], segment)
        }

        one_segments := strings.Split(segment_lookup[1][0], "")

        for i:=0; i<len(segment_lookup[5]); i++ {
            six_segments := strings.Split(segment_lookup[5][i], "")
            first_found := false
            second_found := false
            for j:=0; j<6; j++ {
                if six_segments[j] == one_segments[0] {
                    first_found = true
                }

                if six_segments[j] == one_segments[1] {
                    second_found = true
                }
            }

            if first_found && !second_found {
                segments[1] = one_segments[1]
                segments[2] = one_segments[0]
            } else if !first_found && second_found {
                segments[1] = one_segments[0]
                segments[2] = one_segments[1]
            }
        }

        for _, segment := range strings.Split(segment_lookup[2][0], "") {
            if segment != segments[1] && segment != segments[2] {
                segments[0] = segment
            }
        }

        // 4
        var new_segments []string
        for _, segment := range strings.Split(segment_lookup[3][0], "") {
            if segment != segments[1] && segment != segments[2] {
                new_segments = append(new_segments, segment)
            }
        }

        // 5, 3, 2
        for i:=0; i<len(segment_lookup[4]); i++ {
            five_segments := strings.Split(segment_lookup[4][i], "")

            // of of these will contain both new segments
            contains_count := 0
            contains_segment := ""
            for j:=0; j<5; j++ {
                if five_segments[j] == new_segments[0] {
                    contains_count++;
                    contains_segment = new_segments[0]
                }

                if five_segments[j] == new_segments[1] {
                    contains_count++;
                    contains_segment = new_segments[1]
                }
            }

            if contains_count == 1 {
                segments[6] = contains_segment
                if new_segments[0] != contains_segment {
                    segments[5] = new_segments[0]
                } else {
                    segments[5] = new_segments[1]
                }

                break;
            }
        }

        // 9, 0, 6
        for i:=0; i<len(segment_lookup[5]); i++ {
            six_segments := strings.Split(segment_lookup[5][i], "")

            found_count := 0
            char_not_found := ""
            for j:=0; j<6; j++ {
                char_found := false
                for k:=0; k<7; k++ {
                    if six_segments[j] == segments[k] {
                        found_count++
                        char_found = true
                        break
                    }
                }
                
                if !char_found {
                    char_not_found = six_segments[j]
                }
            }

            if found_count == 5 {
                segments[3] = char_not_found
                break
            }
        }

        for r:='a'; r<'h'; r++ {
            char_found := false
            for k:=0; k<7; k++ {
                if segments[k] == string(r) {
                    char_found = true
                }
            }

            if !char_found {
                segments[4] = string(r)

                break
            }
        }

        var number_list [4]int
        for i, segment := range output {
            var number int
            if len(segment) == 2 {
                number = 1
            } else if len(segment) == 3 {
                number = 7
            } else if len(segment) == 4 {
                number = 4;
            } else if len(segment) == 5 {
                // 5, 3 ,2
                if strings.Contains(segment, segments[5]) {
                    number = 5
                } else if strings.Contains(segment, segments[4]) {
                    number = 2
                } else {
                    number = 3
                }
            } else if len(segment) == 6 {
                // 9, 0, 6
                if !strings.Contains(segment, segments[4]) {
                    number = 9
                } else if !strings.Contains(segment, segments[6]) {
                    number = 0
                } else {
                    number = 6
                }
            } else if len(segment) == 7 {
                number = 8
            }

            number_list[i] = number
        }

        count += list_to_int(number_list);

        for i:=0; i<len(segment_lookup); i++ {
            segment_lookup[i] = nil
        }

        for i:=0; i<7; i++ {
            segments[i] = ""
        }
    }

    fmt.Println(count)
}

func list_to_int(s [4]int) int {
    res := 0
    op := 1
    for i := len(s) - 1; i >= 0; i-- {
        res += s[i] * op
        op *= 10
    }
    return res
}
