package problems

import (
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
const NUM_SECONDS = 100

func SolveProblem14() (int, error) {
	data, err := utils.ReadProblemFile(14)
	if err != nil {
		return 0, err
	}

	robots := convertToRobots(data)
	simulateRobotMovements(robots)
	return getSafetyFactor(robots), nil
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

func simulateRobotMovements(robots []Robot) {
	for i := 0; i < NUM_SECONDS; i++ {
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
	}
}

func getSafetyFactor(robots []Robot) int {
	xSplit, ySplit := BATHROOM_WIDTH/2, BATHROOM_HEIGHT/2

	quadValues := [4]int{0, 0, 0, 0}

	for _, robot := range robots {
		if robot.position.y < ySplit {
			if robot.position.x < xSplit {
				quadValues[0] += 1
			} else if robot.position.x > xSplit {
				quadValues[1] += 1
			}
		} else if robot.position.y > ySplit {
			if robot.position.x < xSplit {
				quadValues[2] += 1
			} else if robot.position.x > xSplit {
				quadValues[3] += 1
			}
		}
	}

	total := 1
	for _, val := range quadValues {
		total *= val
	}
	return total
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
