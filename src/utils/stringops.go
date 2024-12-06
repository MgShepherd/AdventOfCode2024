package utils

import (
	"fmt"
	"strings"
)

func ConvertToGrid(data string) [][]string {
	var grid [][]string
	lines := strings.Split(data, "\n")

	for _, line := range lines {
		if len(strings.TrimSpace(line)) > 0 {
			grid = append(grid, strings.Split(line, ""))
		}
	}

	return grid
}

func PrintGrid(grid [][]string) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			fmt.Printf("%s", grid[y][x])
		}
		fmt.Println()
	}
}
