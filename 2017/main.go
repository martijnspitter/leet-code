package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

func gridGame(grid [][]int) int {
	n := len(grid[0])
	secondScore := 0
	grid1 := make([]int, n)
	grid2 := make([]int, n)
	movedDown := false

	for i := 0; i < n; i++ {
		overLap := 0
		if !movedDown {
			overLap = grid[0][i]
		}
		grid1Sum := sumArr(&grid[0], i)
		grid2Sum := sumArr(&grid[1], i) + overLap

		if grid1Sum >= grid2Sum && !movedDown {
			grid1[i] = 0
			grid2[i] = grid[1][i]
			continue
		}

		if movedDown {
			grid1[i] = grid[0][i]
			grid2[i] = 0
		} else {
			movedDown = true
			grid1[i] = 0
			grid2[i] = 0
		}
	}

	movedDown = false

	for i := 0; i < n; i++ {
		overLap := 0
		if !movedDown {
			overLap = grid1[i]
		}
		grid1Sum := sumArr(&grid1, i)
		grid2Sum := sumArr(&grid2, i) + overLap

		if grid1Sum >= grid2Sum && !movedDown {
			secondScore = secondScore + grid1[i]
			continue
		}

		if movedDown {
			secondScore = secondScore + grid2[i]
		} else {
			movedDown = true
			secondScore = secondScore + grid2[i] + grid1[i]
		}
	}

	return secondScore
}

func sumArr(arr *[]int, index int) int {
	sum := 0
	for i, val := range *arr {
		if i >= index {
			sum = sum + val
		}
	}

	return sum
}
