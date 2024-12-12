package problems

import (
	"github.com/MgShepherd/AdventOfCode2024/src/utils"
)

type Position struct {
	x, y int
}

func SolveProblem12() (int, error) {
	data, err := utils.ReadProblemFile(12)
	if err != nil {
		return 0, err
	}

	grid := utils.ConvertToGrid(data)

	total := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if len(grid[y][x]) == 1 {
				area, perimeter := getRegionInformation(grid, x, y)
				total += area * perimeter
			}
		}
	}

	return total, nil
}

func getRegionInformation(grid [][]string, x, y int) (int, int) {
	positionsToCheck := []Point{{x, y}}

	area, perimeter := 0, 0
	regionVal := grid[y][x]

	for len(positionsToCheck) > 0 {
		position := positionsToCheck[0]
		if len(grid[position.y][position.x]) > 1 {
			positionsToCheck = positionsToCheck[1:]
			continue
		}
		area += 1

		if partOfReigion(grid, position.x-1, position.y, regionVal) {
			positionsToCheck = append(positionsToCheck, Point{position.x - 1, position.y})
		} else if bordersOtherRegion(grid, position.x-1, position.y, regionVal) {
			perimeter += 1
		}

		if partOfReigion(grid, position.x+1, position.y, regionVal) {
			positionsToCheck = append(positionsToCheck, Point{position.x + 1, position.y})
		} else if bordersOtherRegion(grid, position.x+1, position.y, regionVal) {
			perimeter += 1
		}

		if partOfReigion(grid, position.x, position.y-1, regionVal) {
			positionsToCheck = append(positionsToCheck, Point{position.x, position.y - 1})
		} else if bordersOtherRegion(grid, position.x, position.y-1, regionVal) {
			perimeter += 1
		}

		if partOfReigion(grid, position.x, position.y+1, regionVal) {
			positionsToCheck = append(positionsToCheck, Point{position.x, position.y + 1})
		} else if bordersOtherRegion(grid, position.x, position.y+1, regionVal) {
			perimeter += 1
		}

		grid[position.y][position.x] += "V"
		positionsToCheck = positionsToCheck[1:]
	}

	return area, perimeter
}

func partOfReigion(grid [][]string, x, y int, regionVal string) bool {
	return x >= 0 && y >= 0 && x < len(grid[0]) && y < len(grid) && grid[y][x] == regionVal
}

func bordersOtherRegion(grid [][]string, x, y int, regionVal string) bool {
	return x < 0 || y < 0 || x >= len(grid[0]) || y >= len(grid) || string(grid[y][x][0]) != regionVal
}
