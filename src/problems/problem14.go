package problems

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/MgShepherd/AdventOfCode2024/src/utils"
)

type Robot struct {
	position Point
	velocity Point
}

const BATHROOM_WIDTH = 101
const BATHROOM_HEIGHT = 103

func SolveProblem14() (int, error) {
	data, err := utils.ReadProblemFile(14)
	if err != nil {
		return 0, err
	}

	robots := convertToRobots(data)
	numSteps := simulateRobotMovements(robots)
	return numSteps, nil
}

func convertToRobots(data string) []Robot {
	robots := []Robot{}
	lines := strings.Split(data, "\n")

	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if len(trimmedLine) > 0 {
			robots = append(robots, processRobot(line))
		}
	}

	return robots
}

func simulateRobotMovements(robots []Robot) int {
	numSeconds := 1
	grid := convertToGrid(robots)
	for true {
		for j := 0; j < len(robots); j++ {
			robots[j].position.x = (robots[j].position.x + robots[j].velocity.x) % BATHROOM_WIDTH
			if robots[j].position.x < 0 {
				robots[j].position.x = BATHROOM_WIDTH + robots[j].position.x
			}
			robots[j].position.y = (robots[j].position.y + robots[j].velocity.y) % BATHROOM_HEIGHT
			if robots[j].position.y < 0 {
				robots[j].position.y = BATHROOM_HEIGHT + robots[j].position.y
			}
		}
		grid = convertToGrid(robots)
		if numSeconds == 7138 {
			fmt.Println("Num seconds:", numSeconds)
			utils.PrintGridNoZero(grid)
			fmt.Println()
			break
		}
		numSeconds += 1
	}

	return numSeconds
}

func containsTree(grid [][]int) bool {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y])-5; x++ {
			if grid[y][x] > 0 && grid[y][x+1] > 0 && grid[y][x+2] > 0 && grid[y][x+3] > 0 && grid[y][x+4] > 0 {
				return true
			}
		}
	}

	return false
}

func convertToGrid(robots []Robot) [][]int {
	grid := make([][]int, BATHROOM_HEIGHT)
	for i := range grid {
		grid[i] = make([]int, BATHROOM_WIDTH)
	}

	for _, robot := range robots {
		grid[robot.position.y][robot.position.x] += 1
	}

	return grid
}

func processRobot(line string) Robot {
	sections := strings.Fields(line)
	return Robot{position: convertToPoint(sections[0]), velocity: convertToPoint(sections[1])}
}

func convertToPoint(section string) Point {
	elements := strings.Split(section, ",")

	xDim, _ := strconv.Atoi(elements[0][2:])
	yDim, _ := strconv.Atoi(elements[1])

	return Point{x: xDim, y: yDim}
}
