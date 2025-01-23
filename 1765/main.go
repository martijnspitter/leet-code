package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}

type Point struct {
	x, y int
}

var directions = []Point{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func highestPeak(isWater [][]int) [][]int {
	m, n := len(isWater), len(isWater[0])
	heights := make([][]int, m)
	queue := make([]Point, 0, m*n)

	// Initialize heights array with -1 (unvisited)
	// and collect water cells for BFS starting points
	for i := range heights {
		heights[i] = make([]int, n)
		for j := range heights[i] {
			if isWater[i][j] == 1 {
				heights[i][j] = 0
				queue = append(queue, Point{i, j})
			} else {
				heights[i][j] = -1
			}
		}
	}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		// Check all four directions
		for _, dir := range directions {
			newX, newY := curr.x+dir.x, curr.y+dir.y

			// Check bounds and if cell is unvisited
			if newX >= 0 && newX < m && newY >= 0 && newY < n && heights[newX][newY] == -1 {
				heights[newX][newY] = heights[curr.x][curr.y] + 1
				queue = append(queue, Point{newX, newY})
			}
		}
	}

	return heights
}
