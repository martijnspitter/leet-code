package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}

const (
	UNVISITED = 0
	VISITING  = 1
	SAFE      = 2
)

func eventualSafeNodes(graph [][]int) []int {
	n := len(graph)
	// state array to track status of each node
	state := make([]int, n)
	result := []int{}

	// Helper function for DFS
	var isSafe func(node int) bool
	isSafe = func(node int) bool {
		// If node is being visited, we found a cycle
		if state[node] == VISITING {
			return false
		}
		// If node has been fully processed, return its safety status
		if state[node] == SAFE {
			return true
		}

		// Mark node as being visited
		state[node] = VISITING

		// Check all neighbors
		for _, neighbor := range graph[node] {
			if !isSafe(neighbor) {
				return false
			}
		}

		// If we get here, all paths from this node are safe
		state[node] = SAFE
		return true
	}

	// Process all nodes
	for node := 0; node < n; node++ {
		if isSafe(node) {
			result = append(result, node)
		}
	}

	return result
}
