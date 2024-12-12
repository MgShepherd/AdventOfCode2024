package problems

import (
	"slices"

	"github.com/MgShepherd/AdventOfCode2024/src/utils"
)

type Position struct {
	x, y int
}

type Side struct {
	constDimension    int
	changingDimension []int
	horizontal        bool
	above             bool
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
				area, numSides := getRegionInformation(grid, x, y)
				total += area * numSides
			}
		}
	}

	return total, nil
}

func getRegionInformation(grid [][]string, x, y int) (int, int) {
	positionsToCheck := []Position{{x, y}}

	sides := []Side{}

	area := 0
	regionVal := grid[y][x]

	for len(positionsToCheck) > 0 {
		position := positionsToCheck[0]
		if len(grid[position.y][position.x]) > 1 {
			positionsToCheck = positionsToCheck[1:]
			continue
		}
		area += 1

		positionsToCheck, sides = checkPosition(grid, Position{position.x - 1, position.y}, sides, positionsToCheck, regionVal, false, false)
		positionsToCheck, sides = checkPosition(grid, Position{position.x + 1, position.y}, sides, positionsToCheck, regionVal, false, true)
		positionsToCheck, sides = checkPosition(grid, Position{position.x, position.y - 1}, sides, positionsToCheck, regionVal, true, false)
		positionsToCheck, sides = checkPosition(grid, Position{position.x, position.y + 1}, sides, positionsToCheck, regionVal, true, true)

		grid[position.y][position.x] += "V"
		positionsToCheck = positionsToCheck[1:]
	}

	return area, len(sides)
}

func checkPosition(grid [][]string, position Position, sides []Side, positionsToCheck []Position, regionVal string, horizontal bool, above bool) ([]Position, []Side) {
	if partOfReigion(grid, position.x, position.y, regionVal) {
		positionsToCheck = append(positionsToCheck, position)
	} else if bordersOtherRegion(grid, position.x, position.y, regionVal) {
		sides = updateSides(sides, position, horizontal, above)
	}

	return positionsToCheck, sides
}

func partOfReigion(grid [][]string, x, y int, regionVal string) bool {
	return x >= 0 && y >= 0 && x < len(grid[0]) && y < len(grid) && string(grid[y][x][0]) == regionVal
}

func bordersOtherRegion(grid [][]string, x, y int, regionVal string) bool {
	return x < 0 || y < 0 || x >= len(grid[0]) || y >= len(grid) || string(grid[y][x][0]) != regionVal
}

func updateSides(sides []Side, newPos Position, horizontal bool, above bool) []Side {
	for i := 0; i < len(sides); i++ {
		if horizontal && sides[i].horizontal && ((above && sides[i].above) || (!above && !sides[i].above)) {
			if sides[i].constDimension == newPos.y &&
				(slices.Contains(sides[i].changingDimension, newPos.x-1) || slices.Contains(sides[i].changingDimension, newPos.x+1)) {
				if !(slices.Contains(sides[i].changingDimension, newPos.x)) {
					sides[i].changingDimension = append(sides[i].changingDimension, newPos.x)
				}
				return sides
			}
		} else if !horizontal && !sides[i].horizontal && ((above && sides[i].above) || (!above && !sides[i].above)) {
			if sides[i].constDimension == newPos.x &&
				(slices.Contains(sides[i].changingDimension, newPos.y-1) || slices.Contains(sides[i].changingDimension, newPos.y+1)) {
				sides[i].changingDimension = append(sides[i].changingDimension, newPos.y)
				return sides
			}
		}

	}

	if horizontal {
		return append(sides, Side{constDimension: newPos.y, changingDimension: []int{newPos.x}, horizontal: true, above: above})
	}
	return append(sides, Side{constDimension: newPos.x, changingDimension: []int{newPos.y}, horizontal: false, above: above})
}
