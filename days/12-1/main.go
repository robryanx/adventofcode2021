package main

import (
    "fmt"
    "strings"
    "unicode"
    "adventofcode/2021/modules/readinput"
)

func main() {
    connections := readinput.ReadStrings("inputs/12/input.txt", "\n")

    connections_lookup := make(map[string][]string)

    for _, connection := range connections {
        connection_parts := strings.Split(connection, "-")

        connections_lookup[connection_parts[0]] = append(connections_lookup[connection_parts[0]], connection_parts[1])
        connections_lookup[connection_parts[1]] = append(connections_lookup[connection_parts[1]], connection_parts[0])
    }

    visited := make(map[string]bool)

    var current_path []string
    var completed_paths [][]string

    completed_paths = visit(connections_lookup, "start", visited, current_path, completed_paths)

    fmt.Println(len(completed_paths))
}

func visit(connections_lookup map[string][]string, node string, visited map[string]bool, current_path []string, completed_paths [][]string) [][]string {
    current_path = append(current_path, node)
    visited[node] = true

    if node == "end" {
        completed_paths = append(completed_paths, current_path)
    } else if _, ok := connections_lookup[node]; ok {
        for _, node_connection := range connections_lookup[node] {
            if _, ok := visited[node_connection]; IsUpper(node_connection) || !ok {
                copy_visited := CopyMap(visited)

                completed_paths = visit(connections_lookup, node_connection, copy_visited, current_path, completed_paths)
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

func CopyMap(m map[string]bool) map[string]bool {
    cp := make(map[string]bool)
    for k, v := range m {
        cp[k] = v
    }

    return cp
}