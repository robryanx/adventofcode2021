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
    lines := readinput.ReadStrings("inputs/4/input.txt", "\n")

    var draw_order []int

    for _, draw_number_str := range strings.Split(lines[0], ",") {
        draw_number, err := strconv.Atoi(draw_number_str)
        check(err)

        draw_order = append(draw_order, draw_number)
    }

    var boards [][5][5]int
    var current_board [5][5]int
    for i:=1; i<len(lines); i+=6 {
        for j:=i; j<(i+5); j++ {
            for index, board_number_str := range strings.Fields(lines[j+1]) {
                board_number, err := strconv.Atoi(board_number_str)
                check(err)

                current_board[j-i][index] = board_number
            }
        }

        boards = append(boards, current_board)
        current_board = [5][5]int{}
    }

    for i:=0; i<len(draw_order); i++ {
        for j:=0; j<len(boards); j++ {
            board, is_winner := check_board(boards[j], draw_order[0:i])

            if is_winner {
                fmt.Println(score_board(board) * draw_order[i-1])

                goto done;
            }
        }
    }

    done:
}

func check_board(board [5][5]int, values []int) ([5][5]int, bool) {
    for _, value := range values {
        for i:=0; i<5; i++ {
            for j:=0; j<5; j++ {
                if board[i][j] == value {
                    board[i][j] = -1;
                }
            }
        }
    }

    var max int

    // check rows for match
    for i:=0; i<5; i++ {
        max = -1
        for j:=0; j<5; j++ {
            if board[i][j] > -1 {
                max = board[i][j];
                break;
            }
        }

        if max == -1 {
            return board, true;
        }
    }

    // check columns for match
    for j:=0; j<5; j++ {
        max = -1
        for i:=0; i<5; i++ {
            if board[i][j] > -1 {
                max = board[i][j];
                break;
            }
        }

        if max == -1 {
            return board, true;
        }
    }

    return board, false
}

func score_board(board [5][5]int) int {
    score := 0
    for i:=0; i<5; i++ {
        for j:=0; j<5; j++ {
            if board[i][j] != -1 {
                score += board[i][j]
            }
        }
    }

    return score
}
