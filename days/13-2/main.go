package main

import (
	"adventofcode/2021/modules/readinput"
    "adventofcode/2021/modules/grid"
	"strings"
    "strconv"
)

func main() {
    lines := readinput.ReadStrings("inputs/13/input.txt", "\n")

    var positions [][2]int
    var folds [][2]int

    max_x := 0
    max_y := 0

    mode := "pos"
    for _, line := range lines {
        if line == "" {
            mode = "fold"
        } else if mode == "pos" {
            pos := strings.Split(line, ",")

            x, _ := strconv.Atoi(pos[0])
            y, _ := strconv.Atoi(pos[1])

            if x > max_x {
                max_x = x
            }

            if y > max_y {
                max_y = y
            }

            positions = append(positions, [2]int{x, y})
        } else {
            fold := strings.Replace(line, "fold along ", "", -1)
            fold_parts := strings.Split(fold, "=")

            direction := 0
            if fold_parts[0] == "y" {
                direction = 1
            }

            fold_pos, _ := strconv.Atoi(fold_parts[1])

            folds = append(folds, [2]int{direction, fold_pos})
        }
    }

    max_x++
    max_y++

    make_grid := make([][]uint8, max_y)
    for y:=0; y<max_y; y++ {
        make_grid[y] = make([]uint8, max_x)
    }

    for _, position := range positions {
        make_grid[position[1]][position[0]] = 1
    }

    for _, fold := range folds {
        if fold[0] == 1 {
            for y:=fold[1]+1; y<max_y; y++ {
                paste_y := fold[1] - (y - fold[1]);
                if paste_y >= 0 {
                    for x:=0; x<max_x; x++ {
                        if make_grid[y][x] == 1 {
                            make_grid[paste_y][x] = 1
                        }
                    }
                }  
            }

            max_y = fold[1]
        } else {
            for x:=fold[1]+1; x<max_x; x++ {
                paste_x := fold[1] - (x - fold[1]);
                if paste_x >= 0 {
                    for y:=0; y<max_y; y++ {
                        if make_grid[y][x] == 1 {
                            make_grid[y][paste_x] = 1
                        }
                    }
                }  
            }

            max_x = fold[1]
        }
    }

    // count the dots
    count := 0
    for y:=0; y<max_y; y++ {
        for x:=0; x<max_x; x++ {
            if make_grid[y][x] == 1 {
                count++
            }
        }
    }

    grid.Print_grid_part(make_grid, max_y, max_x)

    // PFKLKCFP
}
