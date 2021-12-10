package main

import (
    "fmt"
    "strings"
    "adventofcode/2021/modules/readinput"
)

func main() {
    lines := readinput.ReadStrings("inputs/10/input.txt", "\n")

    var chunks_open []string

    score := 0
    for _, line := range lines {
        chars := strings.Split(line, "")
        for i:=0; i<len(chars); i++ {
            if chars[i] == "(" || chars[i] == "[" || chars[i] == "{" || chars[i] == "<" {
                chunks_open = append(chunks_open, chars[i])
            } else if (chars[i] == ")" && chunks_open[len(chunks_open)-1] == "(") ||
                      (chars[i] == "]" && chunks_open[len(chunks_open)-1] == "[") ||
                      (chars[i] == "}" && chunks_open[len(chunks_open)-1] == "{") ||
                      (chars[i] == ">" && chunks_open[len(chunks_open)-1] == "<") {
                chunks_open = chunks_open[:len(chunks_open)-1]
            } else {
                switch chars[i] {
                case ")":
                    score += 3
                    break
                case "]":
                    score += 57
                    break
                case "}":
                    score += 1197
                    break
                case ">":
                    score += 25137
                    break
                }

                break;
            }
        }

        chunks_open = nil
    }

    fmt.Println(score)
}