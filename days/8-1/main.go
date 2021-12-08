package main

import (
    "fmt"
    "strings"
    "adventofcode/2021/modules/readinput"
)

func main() {
    lines := readinput.ReadStrings("inputs/8/input.txt", "\n")

    var count int
    for _, line := range lines {
        parts := strings.Split(line, " | ")
        output := strings.Fields(parts[1])

        for _, segment := range output {
            if len(segment) == 2 || len(segment) == 3 || len(segment) == 4 || len(segment) == 7 {
                count++
            }
        }
    }

    fmt.Println(count)
}
