package main

import (
    "fmt"
    "adventofcode/2021/modules/readinput"
)

func main() {
    fish := readinput.ReadInts("inputs/6/input.txt", ",")

    var new_fish []int

    days := 80
    for current_day:=0; current_day<days; current_day++ {
        for i:=0; i<len(fish); i++ {
            if fish[i] == 0 {
                fish[i] = 6
                new_fish = append(new_fish, 8)
            } else {
                fish[i]--
            }
        }

        fish = append(fish, new_fish...)
        new_fish = nil
    }

    fmt.Println(len(fish))
}