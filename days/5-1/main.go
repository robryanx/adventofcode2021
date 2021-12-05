package main

import (
    "fmt"
    //"strings"
    "strconv"
    "regexp"
    "adventofcode/2021/modules/readinput"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

type path struct {
    start int
    end int
    direction int
    slope int
    constant int
}

func main() {
    lines := readinput.ReadStrings("inputs/5/input.txt", "\n")

    r_path, _ := regexp.Compile("([0-9]+),([0-9]+) -> ([0-9]+),([0-9]+)")

    var grid [1000][1000]int
    var valid_paths []path
    for _, line := range lines {
        path_matches := r_path.FindStringSubmatch(line)

        var path_int [4]int
        for i:=1; i<5; i++ {
            val, _ := strconv.Atoi(path_matches[i])
            path_int[i-1] = int(val)
        }

        start := [2]int{path_int[0], path_int[1]}
        end := [2]int{path_int[2], path_int[3]}

        if start[0] == end[0] || start[1] == end[1] {
            var slope int
            var make_path path

            if start[0] == end[0] {
                if start[1] > end[1] {
                    slope = -1
                } else {
                    slope = 1
                }

                make_path = path {
                    start: start[1],
                    end: end[1],
                    direction: 1,
                    slope: slope,
                    constant: start[0],
                }
            } else {
                if start[0] > end[0] {
                    slope = -1
                } else {
                    slope = 1
                }

                make_path = path {
                    start: start[0],
                    end: end[0],
                    direction: 0,
                    slope: slope,
                    constant: start[1],
                }
            }

            valid_paths = append(valid_paths, make_path)
        }
    }

    for _, path := range valid_paths {
        if path.slope == 1 {
            for i:=path.start; i<=path.end; i++ {
                if path.direction == 1 {
                    grid[path.constant][i]++
                } else {
                    grid[i][path.constant]++
                }
            }
        } else {
            for i:=path.start; i>=path.end; i-- {
                if path.direction == 1 {
                    grid[path.constant][i]++
                } else {
                    grid[i][path.constant]++
                }
            }
        }
    }

    print_grid(grid)

    fmt.Println(score_grid(grid))
}

func print_grid(grid [1000][1000]int) {
    max_y := len(grid)
    max_x := len(grid[0])

    for y := 0; y<max_y; y++ {
        for x := 0; x<max_x; x++ {
            fmt.Printf("%d ", grid[x][y]);
        }

        fmt.Printf("\n");
    }

    fmt.Printf("\n\n\n");
}

func score_grid(grid [1000][1000]int) int {
    score := 0

    max_y := len(grid)
    max_x := len(grid[0])

    for y := 0; y<max_y; y++ {
        for x := 0; x<max_x; x++ {
            if grid[x][y] > 1 {
                score++
            }
        }
    }

    return score
}