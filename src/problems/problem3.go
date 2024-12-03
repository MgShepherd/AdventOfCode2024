package problems

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/MgShepherd/AdventOfCode2024/src/utils"
)

func SolveProblem3() (int, error) {
	data, err := utils.ReadProblemFile(3)
	if err != nil {
		return -1, err
	}

	totalSum := 0
	shouldProcessMul := true
	for i := 0; i < len(data); i++ {
		if data[i] == 'm' && shouldProcessMul {
			valid, result := processMul(data, i)
			if valid {
				totalSum += result
			}
		} else if data[i] == 'd' {
			valid, result := processDoDont(data, i)
			if valid {
				shouldProcessMul = result
			}
		}
	}

	return totalSum, nil
}

func processDoDont(data string, startIndex int) (bool, bool) {
	const do = "do()"
	const dont = "don't()"

	if data[startIndex:startIndex+len(do)] == do {
		return true, true
	} else if data[startIndex:startIndex+len(dont)] == dont {
		return true, false
	}

	return false, false
}

func processMul(data string, startIndex int) (bool, int) {
	const mulInstruction = "mul("
	for i := 0; i < len(mulInstruction); i++ {
		if data[startIndex+i] != mulInstruction[i] {
			return false, 0
		}
	}

	num1StartIndex := startIndex + len(mulInstruction)
	valid, num1, num1Len := processNumber(data, num1StartIndex, ',')
	if !valid {
		return false, 0
	}

	num2StartIndex := startIndex + len(mulInstruction) + num1Len
	valid, num2, _ := processNumber(data, num2StartIndex, ')')
	if !valid {
		return false, 0
	}

	return true, num1 * num2
}

func processNumber(data string, startIndex int, endChar byte) (bool, int, int) {
	var numString strings.Builder
	for i := 0; i < 4; i++ {
		if data[startIndex+i] == endChar {
			num, err := strconv.Atoi(numString.String())
			if err != nil {
				return false, 0, 0
			}
			return true, num, i + 1
		} else if unicode.IsDigit(rune(data[startIndex+i])) {
			numString.WriteByte(data[startIndex+i])
		} else {
			return false, 0, 0
		}
	}
	return false, 0, 0
}
