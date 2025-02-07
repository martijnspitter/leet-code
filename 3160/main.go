package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}

func queryResults(limit int, queries [][]int) []int {
	result := []int{}

	balls := make(map[int]int)

	for _, query := range queries {
		balls[query[0]] = query[1]

		colors := make(map[int]int)

		for key := range balls {
			colors[balls[key]] = 1
		}

		result = append(result, len(colors))
	}

	return result
}
