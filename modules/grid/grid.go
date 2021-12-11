package grid

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

func Grid(filename string) [][]uint8 {
	lines := readinput.ReadStrings(filename, "\n")

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

	return grid
}

func Print_grid(grid [][]uint8) {
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