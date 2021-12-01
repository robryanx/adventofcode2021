package main

import (
    "fmt"
    "adventofcode/2020/modules/readinput"
)

func main() {
    numbers := readinput.ReadInts("inputs/1/input.txt", "\n")

    previous_sum := -1
    increased := 0
    var current_sum int;
    for i:=0; i<len(numbers)-3; i++ {
        current_sum = sum(numbers, i)
        if previous_sum != -1 && current_sum > previous_sum {
            increased++
        }

        previous_sum = current_sum
    }

    fmt.Printf("%d\n", increased);
}

func sum(numbers []int, start int) int {
    sum := 0
    for i:=start; i<(start+3); i++ {
        sum += numbers[i]
    }

    return sum
}