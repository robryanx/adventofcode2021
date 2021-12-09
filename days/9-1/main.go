package main

import (
    "fmt"
    "strings"
    "strconv"
    "adventofcode/2021/modules/readinput"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    lines := readinput.ReadStrings("inputs/9/input.txt", "\n")

    grid := make([][]uint8, len(lines))
    for i := range lines {
        grid[i] = make([]uint8, len(lines[0]))
    }

    for y, line := range lines {
        numbers := strings.Split(line, "")

        for x, s_number := range numbers {
            number, err := strconv.Atoi(s_number)
            check(err)

            grid[y][x] = uint8(number)
        }
    }

    count := 0
    for y:=0; y<len(grid); y++ {
        for x:=0; x<len(grid[0]); x++ {
            if ((x-1)<0 || grid[y][x-1] > grid[y][x]) &&
               ((x+1)>=len(grid[0]) || grid[y][x+1] > grid[y][x]) && 
               ((y-1)<0 || grid[y-1][x] > grid[y][x]) &&
               ((y+1)>=len(grid) || grid[y+1][x] > grid[y][x]) {
                count += 1 + int(grid[y][x])
            }
        }
    }

    fmt.Println(count)
}