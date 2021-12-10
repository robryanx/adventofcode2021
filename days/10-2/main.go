package main

import (
    "fmt"
    "strings"
    "sort"
    "adventofcode/2021/modules/readinput"
)

func main() {
    lines := readinput.ReadStrings("inputs/10/input.txt", "\n")

    var chunks_open []string
    var line_scores []int

    for _, line := range lines {
        discard := false
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
                discard = true

                break;
            }
        }

        if !discard {
            line_score := 0
            for i:=len(chunks_open)-1; i>=0; i-- {
                line_score *= 5
                
                switch chunks_open[i] {
                case "(":
                    line_score += 1
                    break
                case "[":
                    line_score += 2
                    break
                case "{":
                    line_score += 3
                    break
                case "<":
                    line_score += 4
                    break
                }
            }

            line_scores = append(line_scores, line_score)
        }

        chunks_open = nil
    }

    sort.Ints(line_scores)

    fmt.Println(line_scores[len(line_scores)/2])
}

