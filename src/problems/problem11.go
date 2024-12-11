package problems

import (
	"strconv"
	"strings"

	"github.com/MgShepherd/AdventOfCode2024/src/utils"
)

const numBlinks = 75

func SolveProblem11() (int, error) {
	data, err := utils.ReadProblemFile(11)
	if err != nil {
		return 0, err
	}

	stones := strings.Fields(data)
	stoneMap := make(map[string]int)
	for _, stone := range stones {
		stoneMap[stone] += 1
	}
	numStones := 0
	for i := 0; i < numBlinks; i++ {
		stoneMap, numStones = processBlink(stoneMap)
	}

	return numStones, nil
}

func processBlink(stones map[string]int) (map[string]int, int) {
	newStones := make(map[string]int)

	for stone, num := range stones {
		if stone == "0" {
			newStones["1"] += num
		} else if len(stone)%2 == 0 {
			midPoint := len(stone) / 2
			newStones[convertToStone(stone[:midPoint])] += num
			newStones[convertToStone(stone[midPoint:])] += num
		} else {
			intVal, _ := strconv.Atoi(stone)
			newStones[strconv.Itoa(intVal*2024)] += num
		}
	}

	return newStones, getNumStones(newStones)
}

func getNumStones(stones map[string]int) int {
	total := 0
	for _, v := range stones {
		total += v
	}
	return total
}

func convertToStone(number string) string {
	intVal, _ := strconv.Atoi(number)
	return strconv.Itoa(intVal)
}
