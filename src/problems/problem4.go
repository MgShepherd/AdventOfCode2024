package problems

import (
	"strings"

	"github.com/MgShepherd/AdventOfCode2024/src/utils"
)

func SolveProblem4() (int, error) {
	data, err := utils.ReadProblemFile(4)
	if err != nil {
		return 0, err
	}
	grid := convertToGrid(data)

	foundWords := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == "X" {
				foundWords += seachAllDirections("M", x, y, []string{"A", "S"}, grid)
			}
		}
	}

	return foundWords, nil
}

func convertToGrid(data string) [][]string {
	var grid [][]string
	lines := strings.Split(data, "\n")

	for _, line := range lines {
		if len(strings.TrimSpace(line)) > 0 {
			grid = append(grid, strings.Split(line, ""))
		}
	}

	return grid
}

func seachAllDirections(letter string, xPos, yPos int, remaingLetters []string, grid [][]string) int {
	numFound := 0
	for yDir := -1; yDir < 2; yDir++ {
		for xDir := -1; xDir < 2; xDir++ {
			if searchForString(letter, xPos, yPos, xDir, yDir, remaingLetters, grid) {
				numFound += 1
			}
		}
	}

	return numFound
}

func searchForString(letter string, xPos, yPos, xDir, yDir int, remaingLetters []string, grid [][]string) bool {
	x := xPos + xDir
	y := yPos + yDir
	if y >= 0 && x >= 0 && y < len(grid) && x < len(grid[y]) && grid[y][x] == letter {
		if len(remaingLetters) == 0 {
			return true
		} else {
			return searchForString(remaingLetters[0], x, y, xDir, yDir, remaingLetters[1:], grid)
		}
	}

	return false
}
