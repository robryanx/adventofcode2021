package main

import (
    "fmt"
    "strings"
    "strconv"
    "adventofcode/2020/modules/readinput"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    commands := readinput.ReadStrings("inputs/2/input.txt", "\n")

    x := 0
    y := 0
    for _, command := range commands {
        parts := strings.Split(command, " ")

        distance, err := strconv.Atoi(parts[1])
        check(err)

        switch parts[0] {
        case "up":
            y -= distance
        case "down":
            y += distance
        case "forward":
            x += distance
        }
    }

    fmt.Printf("%d\n", x*y);
}