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

func PrintGrid[T string | int](grid [][]T) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			fmt.Printf("%v", grid[y][x])
		}
		fmt.Println()
	}
}

func PrintGridNoZero(grid [][]int) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] != 0 {
				fmt.Printf("%v", grid[y][x])
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
}
