package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}

func countServers(grid [][]int) int {
	count := 0
	rows := make(map[int]int)
	cols := make(map[int]int)

	for i, row := range grid {
		for j, val := range row {
			if val == 1 {
				if rows[i] == 1 {
					count = count + 2
				}
				if rows[i] > 1 {
					count++
				}
				rows[i] = rows[i] + 1

				if cols[j] == 1 {
					count = count + 2
				}
				if cols[j] > 1 {
					count++
				}
				cols[j] = rows[j] + 1
			}
		}
	}

	return count
}
