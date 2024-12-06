package problems

import (
	"github.com/MgShepherd/AdventOfCode2024/src/utils"
)

func SolveProblem4() (int, error) {
	data, err := utils.ReadProblemFile(4)
	if err != nil {
		return 0, err
	}
	grid := utils.ConvertToGrid(data)

	foundWords := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == "A" {
				if findMasCross(x, y, grid) {
					foundWords += 1
				}
			}
		}
	}

	return foundWords, nil
}

func findMasCross(x, y int, grid [][]string) bool {
	if (searchForString("M", x, y, -1, -1, grid) && searchForString("S", x, y, 1, 1, grid)) ||
		(searchForString("M", x, y, 1, 1, grid) && searchForString("S", x, y, -1, -1, grid)) {
		if (searchForString("M", x, y, 1, -1, grid) && searchForString("S", x, y, -1, 1, grid)) ||
			(searchForString("M", x, y, -1, 1, grid) && searchForString("S", x, y, 1, -1, grid)) {
			return true
		}

	}
	return false
}

func searchForString(letter string, xPos, yPos, xDir, yDir int, grid [][]string) bool {
	x := xPos + xDir
	y := yPos + yDir
	if y >= 0 && x >= 0 && y < len(grid) && x < len(grid[y]) && grid[y][x] == letter {
		return true
	}

	return false
}
