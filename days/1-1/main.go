package main

import (
    "fmt"
    "adventofcode/2021/modules/readinput"
)

func main() {
    numbers := readinput.ReadInts("inputs/1/input.txt", "\n")

    increased := 0
    for i:=1; i<len(numbers); i++ {
        if(numbers[i] > numbers[i-1]) {
            increased++
        }
    }

    fmt.Printf("%d\n", increased);
}