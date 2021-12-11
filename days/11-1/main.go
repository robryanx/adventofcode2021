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
    lines := readinput.ReadStrings("inputs/11/input.txt", "\n")

    grid := make([][]uint8, len(lines))
    for row, line := range lines {
        numbers := strings.Split(line, "")
        grid[row] = make([]uint8, len(numbers))

        for i:=0; i<len(numbers); i++ {
            number, err := strconv.Atoi(numbers[i])
            check(err)

            grid[row][i] = uint8(number)
        }
    }

    flashes := 0
    for step:=0; step<100; step++ {
        for y:=0; y<len(grid); y++ {
            for x:=0; x<len(grid[0]); x++ {
                grid[y][x]++
            }
        }

        for y:=0; y<len(grid); y++ {
            for x:=0; x<len(grid[0]); x++ {
                if grid[y][x] > 9 {
                    grid, flashes = flash(grid, y, x, flashes)
                }
            }
        }

        for y:=0; y<len(grid); y++ {
            for x:=0; x<len(grid[0]); x++ {
                if grid[y][x] > 100 {
                    grid[y][x] = 0
                }
            }
        }
    }

    fmt.Println(flashes)
}

func flash(grid [][]uint8, y int, x int, flashes int) ([][]uint8, int) {
    if grid[y][x] > 100 {
        return grid, flashes
    }

    flashes++
    grid[y][x] = 101

    if x-1 >= 0 {
        grid[y][x-1]++
        if grid[y][x-1] > 9 {
            grid, flashes = flash(grid, y, x-1, flashes)
        }

        if y-1 >= 0 {
            grid[y-1][x-1]++
            if grid[y-1][x-1] > 9 {
                grid, flashes = flash(grid, y-1, x-1, flashes)
            }
        }

        if y+1 < len(grid) {
            grid[y+1][x-1]++
            if grid[y+1][x-1] > 9 {
                grid, flashes = flash(grid, y+1, x-1, flashes)
            }
        }
    }

    if x+1 < len(grid[0]) {
        grid[y][x+1]++
        if grid[y][x+1] > 9 {
            grid, flashes = flash(grid, y, x+1, flashes)
        }

        if y-1 >= 0 {
            grid[y-1][x+1]++
            if grid[y-1][x+1] > 9 {
                grid, flashes = flash(grid, y-1, x+1, flashes)
            }
        }
        if y+1 < len(grid) {
            grid[y+1][x+1]++
            if grid[y+1][x+1] > 9 {
                grid, flashes = flash(grid, y+1, x+1, flashes)
            }
        }
    }

    if y-1 >= 0 {
        grid[y-1][x]++
        if grid[y-1][x] > 9 {
            grid, flashes = flash(grid, y-1, x, flashes)
        }
    }

    if y+1 < len(grid) {
        grid[y+1][x]++
        if grid[y+1][x] > 9 {
            grid, flashes = flash(grid, y+1, x, flashes)
        }
    }

    return grid, flashes
}

func print_grid(grid [][]uint8) {
    max_y := len(grid)
    max_x := len(grid[0])

    for y := 0; y<max_y; y++ {
        for x := 0; x<max_x; x++ {
            fmt.Printf("%d ", grid[y][x]);
        }

        fmt.Printf("\n");
    }

    fmt.Printf("\n\n\n");
}