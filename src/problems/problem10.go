package problems

import (
	"strconv"

	"github.com/MgShepherd/AdventOfCode2024/src/utils"
)

func SolveProblem10() (int, error) {
	data, err := utils.ReadProblemFile(10)
	if err != nil {
		return 0, err
	}

	grid := utils.ConvertToGrid(data)
	totalScore := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == "0" {
				totalScore += getTrailheadScore(grid, x, y, 0)
			}
		}
	}
	return totalScore, nil
}

func getTrailheadScore(grid [][]string, x, y, currentValue int) int {
	score := 0

	if currentValue == 9 {
		return 1
	}

	nextValue := strconv.Itoa(currentValue + 1)
	if doesLocationHaveValue(grid, x-1, y, nextValue) {
		score += getTrailheadScore(grid, x-1, y, currentValue+1)
	}
	if doesLocationHaveValue(grid, x+1, y, nextValue) {
		score += getTrailheadScore(grid, x+1, y, currentValue+1)
	}
	if doesLocationHaveValue(grid, x, y-1, nextValue) {
		score += getTrailheadScore(grid, x, y-1, currentValue+1)
	}
	if doesLocationHaveValue(grid, x, y+1, nextValue) {
		score += getTrailheadScore(grid, x, y+1, currentValue+1)
	}

	return score
}

func doesLocationHaveValue(grid [][]string, x, y int, value string) bool {
	return x >= 0 && y >= 0 && x < len(grid[0]) && y < len(grid) && grid[y][x] == value
}
