package problems

import (
	"strconv"
	"strings"

	"github.com/MgShepherd/AdventOfCode2024/src/utils"
)

const numBlinks = 25

func SolveProblem11() (int, error) {
	data, err := utils.ReadProblemFile(11)
	if err != nil {
		return 0, err
	}

	stones := strings.Fields(data)
	for i := 0; i < numBlinks; i++ {
		stones = processBlink(stones)
	}

	return len(stones), nil
}

func processBlink(stones []string) []string {
	newStones := []string{}

	for _, stone := range stones {
		if stone == "0" {
			newStones = append(newStones, "1")
		} else if len(stone)%2 == 0 {
			midPoint := len(stone) / 2
			newStones = append(newStones, convertToStone(stone[:midPoint]), convertToStone(stone[midPoint:]))
		} else {
			intVal, _ := strconv.Atoi(stone)
			newStones = append(newStones, strconv.Itoa(intVal*2024))
		}
	}

	return newStones
}

func convertToStone(number string) string {
	intVal, _ := strconv.Atoi(number)
	return strconv.Itoa(intVal)
}
