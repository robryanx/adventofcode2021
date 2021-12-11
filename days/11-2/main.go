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

    permutations := permutations()
    
    step:=1
    for ;; {
        flashes := 0
        for y:=0; y<len(grid); y++ {
            for x:=0; x<len(grid[0]); x++ {
                grid[y][x]++
            }
        }

        for y:=0; y<len(grid); y++ {
            for x:=0; x<len(grid[0]); x++ {
                if grid[y][x] > 9 {
                    grid, flashes = flash(permutations, grid, y, x, flashes)
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

        if flashes == len(grid)*len(grid[0]) {
            fmt.Println(step)
            break
        }

        step++
    }
}

func permutations() [][2]int {
    var perms [][2]int

    x_d := [3]int{-1, 0, 1}
    y_d := [3]int{-1, 0, 1}

    for _, x := range x_d {
        for _, y := range y_d {
            if x != 0 || y != 0 {
                perms = append(perms, [2]int{y, x})
            }
        }
    }

    return perms
}

func flash(permutations [][2]int, grid [][]uint8, y int, x int, flashes int) ([][]uint8, int) {
    if grid[y][x] > 100 {
        return grid, flashes
    }

    flashes++
    grid[y][x] = 101

    for _, delta := range permutations {
        if y + delta[0] >= 0 && 
           y + delta[0] < len(grid) &&
           x + delta[1] >= 0 &&
           x + delta[1] < len(grid[0]) {
            grid[y+delta[0]][x+delta[1]]++;
            if grid[y+delta[0]][x+delta[1]] > 9 {
                grid, flashes = flash(permutations, grid, y+delta[0], x+delta[1], flashes)
            }
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