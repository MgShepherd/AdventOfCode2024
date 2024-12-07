package problems

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/MgShepherd/AdventOfCode2024/src/utils"
)

type Operation int

const (
	OpMul Operation = iota
	OpAdd
	OpCat
)

func (o Operation) perform(val1, val2 int) int {
	switch o {
	case OpMul:
		return val1 * val2
	case OpAdd:
		return val1 + val2
	default:
		return concatValues(val1, val2)
	}
}

func concatValues(val1, val2 int) int {
	combined := strconv.Itoa(val1) + strconv.Itoa(val2)
	result, err := strconv.Atoi(combined)
	if err != nil {
		fmt.Printf("Unable to concat values %d and %d\n", val1, val2)
	}

	return result
}

func SolveProblem7() (int, error) {
	data, err := utils.ReadProblemFile(7)
	if err != nil {
		return 0, err
	}

	validEquationSum := 0
	lines := strings.Split(data, "\n")
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if len(trimmedLine) > 0 {
			target, operands, err := getTargetAndOperands(trimmedLine)
			if err != nil {
				return 0, fmt.Errorf("Unable to process line %s\n", trimmedLine)
			}

			if checkOperation(target, operands, OpAdd) || checkOperation(target, operands, OpMul) || checkOperation(target, operands, OpCat) {
				validEquationSum += target
			}
		}
	}

	return validEquationSum, nil
}

func getTargetAndOperands(equation string) (int, []int, error) {
	items := strings.Split(equation, ":")
	target, err := strconv.Atoi(items[0])
	if err != nil {
		return 0, []int{}, err
	}
	operands, err := utils.ConvertToIntSlice(strings.Fields(items[1]))
	if err != nil {
		return 0, []int{}, err
	}

	return target, operands, nil
}

func checkOperation(target int, operands []int, operation Operation) bool {
	result := operation.perform(operands[0], operands[1])

	if result > target {
		return false
	} else if len(operands) > 2 {
		newOperands := append([]int{result}, operands[2:]...)
		return checkOperation(target, newOperands, OpAdd) || checkOperation(target, newOperands, OpMul) || checkOperation(target, newOperands, OpCat)
	}

	return result == target
}
