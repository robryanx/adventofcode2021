package main

import (
    "fmt"
    "strings"
    "sort"
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

    var sizes []int
    for y:=0; y<len(grid); y++ {
        for x:=0; x<len(grid[0]); x++ {
            if ((x-1)<0 || grid[y][x-1] > grid[y][x]) &&
               ((x+1)>=len(grid[0]) || grid[y][x+1] > grid[y][x]) && 
               ((y-1)<0 || grid[y-1][x] > grid[y][x]) &&
               ((y+1)>=len(grid) || grid[y+1][x] > grid[y][x]) {
                seen := make(map[string]bool)

                seen = calculate_basin(grid, y, x, seen)

                sizes = append(sizes, len(seen))
            }
        }
    }

    sort.Sort(sort.Reverse(sort.IntSlice(sizes)))

    fmt.Println(sizes[0] * sizes[1] * sizes[2])
}

func calculate_basin(grid [][]uint8, y int, x int, seen map[string]bool) map[string]bool {
    pos := strconv.Itoa(y) + "-" + strconv.Itoa(x)
    seen[pos] = true

    if (x-1)>=0 && grid[y][x-1] != 9 {
        check_pos := strconv.Itoa(y) + "-" + strconv.Itoa(x-1)
        if _, ok := seen[check_pos]; !ok {
            seen = calculate_basin(grid, y, x-1, seen)
        }
    }

    if (x+1)<len(grid[0]) && grid[y][x+1] != 9 {
        check_pos := strconv.Itoa(y) + "-" + strconv.Itoa(x+1)
        if _, ok := seen[check_pos]; !ok {
            seen = calculate_basin(grid, y, x+1, seen)
        }
    }

    if (y-1)>=0 && grid[y-1][x] != 9 {
        check_pos := strconv.Itoa(y-1) + "-" + strconv.Itoa(x)
        if _, ok := seen[check_pos]; !ok {
            seen = calculate_basin(grid, y-1, x, seen)
        }
    }

    if (y+1)<len(grid) && grid[y+1][x] != 9 {
        check_pos := strconv.Itoa(y+1) + "-" + strconv.Itoa(x)
        if _, ok := seen[check_pos]; !ok {
            seen = calculate_basin(grid, y+1, x, seen)
        }
    }

    return seen;
}