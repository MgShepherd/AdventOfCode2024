package problems

import (
	"strings"

	"github.com/MgShepherd/AdventOfCode2024/src/utils"
)

var directionFromStr = map[string]Direction{
	"^": DirUp,
	"v": DirDown,
	">": DirRight,
	"<": DirLeft,
}

func SolveProblem15() (int, error) {
	data, err := utils.ReadProblemFile(15)
	if err != nil {
		return 0, err
	}

	grid, moves := splitFileIntoSections(data)
	startPos := findRobotStartingPosition(grid)
	processMoves(grid, moves, startPos)
	return getSumGPSCoords(grid), nil
}

func splitFileIntoSections(data string) ([][]string, string) {
	lines := strings.Split(data, "\n")
	grid := [][]string{}

	i := 0
	for len(lines[i]) > 0 {
		elements := strings.Split(lines[i], "")
		grid = append(grid, elements)
		i += 1
	}
	i += 1

	moves := ""
	for i < len(lines) {
		moves += lines[i]
		i += 1
	}
	return grid, moves
}

func findRobotStartingPosition(grid [][]string) Point {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == "@" {
				return Point{x, y}
			}
		}
	}
	return Point{x: -1, y: -1}
}

func processMoves(grid [][]string, moves string, startPos Point) int {
	elements := strings.Split(moves, "")
	currentPos := startPos
	for _, element := range elements {
		currentPos = processMove(grid, directionFromStr[element], currentPos)
	}
	return 0
}

func getSumGPSCoords(grid [][]string) int {
	sumCoords := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == "O" {
				sumCoords += 100*y + x
			}
		}
	}

	return sumCoords
}

func processMove(grid [][]string, direction Direction, startPos Point) Point {
	moveItems := []string{}
	nextPos := startPos
	for grid[nextPos.y][nextPos.x] != "." {
		if grid[nextPos.y][nextPos.x] == "#" {
			return startPos
		}
		moveItems = append(moveItems, grid[nextPos.y][nextPos.x])
		nextPos = moveInDirection(direction, nextPos)
		if !isPosInBounds(grid, nextPos) {
			return startPos
		}
	}

	backwardsDir := flipDirection(direction)
	for i := len(moveItems) - 1; i >= 0; i-- {
		grid[nextPos.y][nextPos.x] = moveItems[i]
		nextPos = moveInDirection(backwardsDir, nextPos)
	}

	grid[nextPos.y][nextPos.x] = "."
	return moveInDirection(direction, nextPos)
}

func moveInDirection(direction Direction, currentPos Point) Point {
	switch direction {
	case DirUp:
		return Point{x: currentPos.x, y: currentPos.y - 1}
	case DirDown:
		return Point{x: currentPos.x, y: currentPos.y + 1}
	case DirLeft:
		return Point{x: currentPos.x - 1, y: currentPos.y}
	case DirRight:
		return Point{x: currentPos.x + 1, y: currentPos.y}
	}
	return Point{x: -1, y: -1}
}

func flipDirection(direction Direction) Direction {
	switch direction {
	case DirUp:
		return DirDown
	case DirDown:
		return DirUp
	case DirLeft:
		return DirRight
	case DirRight:
		return DirLeft
	}
	return DirUp
}

func isPosInBounds(grid [][]string, pos Point) bool {
	return pos.x >= 0 && pos.y >= 0 && pos.x < len(grid[0]) && pos.y < len(grid)
}
