package problems

import (
	"math"
	"strconv"
	"strings"

	"github.com/MgShepherd/AdventOfCode2024/src/utils"
	"gonum.org/v1/gonum/mat"
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
	aMovement := processLine(lines[0], 2, false)
	bMovement := processLine(lines[1], 2, false)
	prizePos := processLine(lines[2], 1, true)

	aPresses, bPresses := solveSimultaneousEquations(aMovement, bMovement, prizePos)
	if aPresses < 0 || bPresses < 0 {
		return -1
	}
	aPressInt, bPressInt := int(math.Round(aPresses)), int(math.Round(bPresses))
	if aPressInt*aMovement.x+bPressInt*bMovement.x == prizePos.x && aPressInt*aMovement.y+bPressInt*bMovement.y == prizePos.y {
		return int(math.Round(aPresses))*3 + int(math.Round(bPresses))
	}
	return -1
}

func solveSimultaneousEquations(aVals, bVals, prizeVals Position) (float64, float64) {
	eq := mat.NewDense(2, 2, []float64{float64(aVals.x), float64(bVals.x), float64(aVals.y), float64(bVals.y)})
	results := mat.NewVecDense(2, []float64{float64(prizeVals.x), float64(prizeVals.y)})

	var result mat.VecDense
	_ = result.SolveVec(eq, results)
	return result.AtVec(0), result.AtVec(1)
}

func processLine(line string, startLocation int, prizeLine bool) Position {
	elements := strings.Fields(line)

	xElement, _ := strconv.Atoi(elements[startLocation][2 : len(elements[startLocation])-1])
	yElement, _ := strconv.Atoi(elements[startLocation+1][2:])

	if prizeLine {
		return Position{x: xElement + 10000000000000, y: yElement + 10000000000000}
	}
	return Position{x: xElement, y: yElement}
}
