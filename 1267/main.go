package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}

func countServers(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	rowCount := make([]int, m)
	colCount := make([]int, n)

	// Count servers in each row and column
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				rowCount[i]++
				colCount[j]++
			}
		}
	}

	// Count servers that can communicate
	result := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 && (rowCount[i] > 1 || colCount[j] > 1) {
				result++
			}
		}
	}

	return result
}
