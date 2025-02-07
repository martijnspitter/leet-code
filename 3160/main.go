package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}

func queryResults(limit int, queries [][]int) []int {
	result := make([]int, 0, len(queries))
	balls := make(map[int]int)
	colorCount := make(map[int]int)
	distinctColors := 0

	for _, query := range queries {
		ball, newColor := query[0], query[1]

		// Check if the ball already has a color
		if oldColor, exists := balls[ball]; exists {
			if oldColor != newColor {
				// Decrease the count of the old color
				colorCount[oldColor]--
				// If the old color count drops to zero, it's no longer a distinct color
				if colorCount[oldColor] == 0 {
					distinctColors--
				}
			}
		}

		// Update the ball's color only if it changes
		if oldColor, exists := balls[ball]; !exists || oldColor != newColor {
			balls[ball] = newColor
			// Increase the count of the new color
			if colorCount[newColor] == 0 {
				distinctColors++
			}
			colorCount[newColor]++
		}

		// Append the number of distinct colors to the result
		result = append(result, distinctColors)
	}

	return result
}
