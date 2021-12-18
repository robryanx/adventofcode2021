package main

import (
    "fmt"
    "container/heap"
    "adventofcode/2021/modules/grid"
    "adventofcode/2021/modules/priority_queue"
)

type position struct {
    x int
    y int
}

func main() {
    build_grid := grid.Grid("inputs/15/input.txt")

    full_grid := make([][]uint8, len(build_grid)*5)
    for i:=0; i<len(full_grid); i++ {
        full_grid[i] = make([]uint8, len(build_grid[0])*5)
    }

    for y_g:=0; y_g<5; y_g++ {
        for x_g:=0; x_g<5; x_g++ {
            for y:=0; y<len(build_grid); y++ {
                for x:=0; x<len(build_grid[0]); x++ {
                    y_pos := (len(build_grid)*y_g) + y
                    x_pos := (len(build_grid[0])*x_g) + x

                    val := (int(build_grid[y][x]) + y_g + x_g)
                    if val > 9 {
                        val -= 9
                    }

                    full_grid[y_pos][x_pos] = uint8(val)
                }
            }   
        }
    }

    weights := make([][]uint16, len(full_grid))
    for i:=0; i<len(weights); i++ {
        weights[i] = make([]uint16, len(full_grid[0]))
    }

    // create the priority queue
    pq := make(priority_queue.PriorityQueue, 1)

    pq[0] = &priority_queue.Item {
        Data: &position{
            y: 0,
            x: 0,
        },
        Priority: 0,
    };

    heap.Init(&pq)

    pairs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

    for pq.Len() > 0 {
        item := heap.Pop(&pq).(*priority_queue.Item)
        pos := item.Data.(*position)

        // check neighbours
        for _, pair := range pairs {
            if pos.y+pair[0] >= 0 && pos.y+pair[0] < len(full_grid) && pos.x+pair[1] >= 0 && pos.x+pair[1] < len(full_grid[0]) {
                if weights[pos.y+pair[0]][pos.x+pair[1]] == 0 || weights[pos.y+pair[0]][pos.x+pair[1]] > (item.Priority + uint16(full_grid[pos.y+pair[0]][pos.x+pair[1]])) {
                    weights[pos.y+pair[0]][pos.x+pair[1]] = item.Priority + uint16(full_grid[pos.y+pair[0]][pos.x+pair[1]])

                    pq.Push(&priority_queue.Item {
                        Data: &position{
                            y: pos.y+pair[0],
                            x: pos.x+pair[1],
                        },
                        Priority: weights[pos.y+pair[0]][pos.x+pair[1]],
                    });
                }
            }
        }
    }

    //grid.Print_grid_16(weights)

    fmt.Println(weights[len(full_grid)-1][len(full_grid[0])-1])
}
