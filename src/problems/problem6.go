package problems

import (
	"fmt"

	"github.com/MgShepherd/AdventOfCode2024/src/utils"
)

type Direction int

const (
	DirUp Direction = iota
	DirDown
	DirLeft
	DirRight
)

var directionFromStr = map[string]Direction{
	"^": DirUp,
	"v": DirDown,
	">": DirLeft,
	"<": DirRight,
}

func (dir Direction) turnRight() Direction {
	switch dir {
	case DirUp:
		return DirRight
	case DirDown:
		return DirLeft
	case DirLeft:
		return DirUp
	default:
		return DirDown
	}
}

func getNextPosition(x, y int, direction Direction) (int, int) {
	switch direction {
	case DirUp:
		return x, y - 1
	case DirDown:
		return x, y + 1
	case DirLeft:
		return x - 1, y
	default:
		return x + 1, y
	}
}

func SolveProblem6() (int, error) {
	data, err := utils.ReadProblemFile(6)
	if err != nil {
		return 0, err
	}

	grid := utils.ConvertToGrid(data)
	startX, startY, err := findStartingPosition(grid)
	if err != nil {
		return 0, err
	}

	numSteps := getStepsTaken(grid, startX, startY)
	return numSteps, nil
}

func findStartingPosition(grid [][]string) (int, int, error) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == "^" {
				return x, y, nil
			}
		}
	}

	return 0, 0, fmt.Errorf("Unable to find starting position")
}

func getStepsTaken(grid [][]string, startX int, startY int) int {
	currentDir := DirUp
	currentX, currentY := startX, startY
	numVisitedLocations := 0

	for isPositionInBounds(currentX, currentY, grid) {
		if grid[currentY][currentX] != "X" {
			numVisitedLocations += 1
		}
		grid[currentY][currentX] = "X"
		currentX, currentY, currentDir = move(grid, currentX, currentY, currentDir)
	}

	return numVisitedLocations
}

func move(grid [][]string, currentX, currentY int, currentDir Direction) (int, int, Direction) {
	nextX, nextY := getNextPosition(currentX, currentY, currentDir)

	if isPositionInBounds(nextX, nextY, grid) && grid[nextY][nextX] == "#" {
		return currentX, currentY, currentDir.turnRight()
	}
	return nextX, nextY, currentDir
}

func isPositionInBounds(x, y int, grid [][]string) bool {
	return x >= 0 && y >= 0 && x < len(grid[0]) && y < len(grid)
}
