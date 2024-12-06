package problems

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/MgShepherd/AdventOfCode2024/src/utils"
)

type Direction int

const (
	DirUp Direction = iota
	DirDown
	DirLeft
	DirRight
)

var directionToStr = map[Direction]string{
	DirUp:    "^",
	DirDown:  "v",
	DirLeft:  ">",
	DirRight: "<",
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

	numSteps := getNumLoops(grid, startX, startY)
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

func getNumLoops(grid [][]string, startX int, startY int) int {
	currentDir := DirUp
	currentX, currentY := startX, startY
	numLoops := 0
	index := 0

	for isPositionInBounds(currentX, currentY, grid) {
		grid[currentY][currentX] = "V"
		nextX, nextY, nextDir := move(grid, currentX, currentY, currentDir)
		if isPositionInBounds(nextX, nextY, grid) && grid[nextY][nextX] != "#" && (currentX != nextX || currentY != nextY) && !strings.Contains(grid[nextY][nextX], "V") {
			grid[nextY][nextX] = "#"
			if isLoop(grid, currentX, currentY, currentDir, index) {
				numLoops += 1
			}
		}
		index += 1
		currentX, currentY, currentDir = nextX, nextY, nextDir
	}

	return numLoops
}

func isLoop(grid [][]string, startX, startY int, startDir Direction, index int) bool {
	currentX, currentY := startX, startY
	currentDir := startDir
	strIndex := strconv.Itoa(index)
	for isPositionInBounds(currentX, currentY, grid) {
		alreadyVisited := strings.Contains(grid[currentY][currentX], "V")
		if !strings.Contains(grid[currentY][currentX], strIndex) && !alreadyVisited {
			grid[currentY][currentX] = strIndex + directionToStr[currentDir]
		} else if !strings.Contains(grid[currentY][currentX], strIndex) {
			grid[currentY][currentX] = "V" + strIndex + directionToStr[currentDir]
		} else {
			grid[currentY][currentX] += directionToStr[currentDir]
		}

		nextX, nextY, nextDir := move(grid, currentX, currentY, currentDir)
		if !isPositionInBounds(nextX, nextY, grid) {
			return false
		} else if strings.Contains(grid[nextY][nextX], strIndex) && strings.Contains(grid[nextY][nextX], directionToStr[nextDir]) {
			return true
		}
		currentX, currentY, currentDir = nextX, nextY, nextDir
	}

	return false
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
