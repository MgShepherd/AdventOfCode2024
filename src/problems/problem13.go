package problems

import (
	"strconv"
	"strings"

	"github.com/MgShepherd/AdventOfCode2024/src/utils"
)

func SolveProblem13() (int, error) {
	data, err := utils.ReadProblemFile(13)
	if err != nil {
		return 0, err
	}

	lines := strings.Split(data, "\n")
	minTokens := 0
	for i := 0; i < len(lines)-3; i += 4 {
		cheapestWin := getCheapestWayToWin(lines[i : i+3])
		if cheapestWin != -1 {
			minTokens += cheapestWin
		}
	}

	return minTokens, nil
}

func getCheapestWayToWin(lines []string) int {
	aMovement := processLine(lines[0], 2)
	bMovement := processLine(lines[1], 2)
	prizePos := processLine(lines[2], 1)

	cost := findPrize(aMovement, bMovement, prizePos)
	return cost
}

func processLine(line string, startLocation int) Position {
	elements := strings.Fields(line)

	xElement, _ := strconv.Atoi(elements[startLocation][2 : len(elements[startLocation])-1])
	yElement, _ := strconv.Atoi(elements[startLocation+1][2:])

	return Position{x: xElement, y: yElement}
}

func findPrize(aMovement, bMovement, prizePosition Position) int {
	maxAPresses := getMaxButtonPresses(aMovement, prizePosition)
	maxBPresses := getMaxButtonPresses(bMovement, prizePosition)

	minCost := -1
	for i := 0; i <= maxAPresses; i++ {
		for j := 0; j <= maxBPresses; j++ {
			position := Position{x: (aMovement.x * i) + (bMovement.x * j), y: (aMovement.y * i) + (bMovement.y * j)}
			if position == prizePosition {
				cost := i*3 + j
				if minCost == -1 || cost < minCost {
					minCost = cost
				}
				break
			} else if position.x > prizePosition.x || position.y > prizePosition.y {
				break
			}
		}
	}

	return minCost
}

func getMaxButtonPresses(buttonMovement, prizePosition Position) int {
	xVal := prizePosition.x / buttonMovement.x
	yVal := prizePosition.y / buttonMovement.y

	if xVal > 100 && yVal > 100 {
		return 100
	}
	if xVal < yVal {
		return xVal
	}
	return yVal
}
