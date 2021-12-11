package main

import (
    "fmt"
    "adventofcode/2021/modules/grid"
)

func main() {
    grid := grid.Grid("inputs/11/input.txt")

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
