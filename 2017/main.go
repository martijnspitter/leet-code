package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

func gridGame(grid [][]int) int {
	n := len(grid[0])

	// Use prefix sums to avoid repeated summations
	topSum := make([]int64, n+1) // Using int64 to prevent potential overflow
	bottomSum := make([]int64, n+1)

	// Calculate prefix sums
	for i := 0; i < n; i++ {
		topSum[i+1] = topSum[i] + int64(grid[0][i])
		bottomSum[i+1] = bottomSum[i] + int64(grid[1][i])
	}

	fmt.Println("topSum", topSum)
	fmt.Println("bottomSum", bottomSum)

	result := int64(1<<63 - 1) // MaxInt64
	for i := 0; i < n; i++ {
		// For each possible turning point:
		// - Top right sum: topSum[n] - topSum[i+1]
		// - Bottom left sum: bottomSum[i]
		secondRobotScore := max(
			topSum[n]-topSum[i+1], // Remaining top path
			bottomSum[i],          // Used bottom path
		)
		result = min(result, secondRobotScore)
	}

	return int(result)
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
