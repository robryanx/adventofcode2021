package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/robryanx/adventofcode2021/modules/readinput"
)

func main() {
	connections := readinput.ReadStrings("inputs/12/input.txt", "\n")

	connections_lookup := make(map[string][]string)

	for _, connection := range connections {
		connection_parts := strings.Split(connection, "-")

		connections_lookup[connection_parts[0]] = append(connections_lookup[connection_parts[0]], connection_parts[1])
		connections_lookup[connection_parts[1]] = append(connections_lookup[connection_parts[1]], connection_parts[0])
	}

	visited := make(map[string]int)

	var current_path []string
	var completed_paths [][]string

	completed_paths = visit(connections_lookup, "start", visited, current_path, completed_paths)

	//fmt.Println(completed_paths)
	fmt.Println(len(completed_paths))
}

func visit(connections_lookup map[string][]string, node string, visited map[string]int, current_path []string, completed_paths [][]string) [][]string {
	current_path = append(current_path, node)

	if !IsUpper(node) {
		visited[node]++
	}

	if node == "end" {
		completed_paths = append(completed_paths, current_path)
	} else if _, ok := connections_lookup[node]; ok {
		for _, node_connection := range connections_lookup[node] {
			if node_connection != "start" {
				if IsUpper(node_connection) || node_connection == "end" {
					copy_visited := CopyMap(visited)

					completed_paths = visit(connections_lookup, node_connection, copy_visited, current_path, completed_paths)
				} else {
					max := 0
					for _, v := range visited {
						if v > max {
							max = v
						}
					}

					count, _ := visited[node_connection]

					if count < 1 || max < 2 {
						copy_visited := CopyMap(visited)

						completed_paths = visit(connections_lookup, node_connection, copy_visited, current_path, completed_paths)
					}
				}
			}
		}
	}

	return completed_paths
}

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func CopyMap(m map[string]int) map[string]int {
	cp := make(map[string]int)
	for k, v := range m {
		cp[k] = v
	}

	return cp
}
