package main

type Position struct {
	row, col int
}

func Main(arr []int, mat [][]int) int {
	positions := make(map[int]Position, len(mat)*len(mat[0]))
	rowCounts := make([]int, len(mat))
	colCounts := make([]int, len(mat[0]))

	// Build positions map once
	for i := range mat {
		for j := range mat[i] {
			positions[mat[i][j]] = Position{i, j}
		}
	}

	for i, val := range arr {
		pos, exists := positions[val]
		if !exists {
			return -1
		}

		rowCounts[pos.row]++
		colCounts[pos.col]++

		// Check winning condition
		if rowCounts[pos.row] == len(mat[0]) ||
			colCounts[pos.col] == len(mat) {
			return i
		}
	}

	return -1
}
