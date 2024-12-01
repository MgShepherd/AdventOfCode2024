package problems

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/MgShepherd/AdventOfCode2024/src/utils"
)

func SolveProblem1() (int, error) {
	data, err := utils.ReadProblemFile(1)
	if err != nil {
		return -1, err
	}

	locationLists, err := processFile(data)
	if err != nil {
		return -1, err
	}

	frequencies := getValueFrequencies(locationLists[1])

	return getTotalSimilarityScore(locationLists[0], frequencies), nil
}

func processFile(data string) ([2][]int, error) {
	lines := strings.Split(data, "\n")
	var lists [2][]int

	for _, line := range lines {
		values := strings.Fields(line)
		for i := 0; i < len(values); i++ {
			intVal, err := strconv.Atoi(values[i])
			if err != nil {
				fmt.Printf("Unable to process line %s\n", line)
				return lists, err
			}
			lists[i] = append(lists[i], intVal)
		}
	}

	return lists, nil
}

func getValueFrequencies(values []int) map[int]int {
	frequencies := make(map[int]int)

	for _, element := range values {
		currentFreq := frequencies[element]
		frequencies[element] = currentFreq + 1
	}

	return frequencies
}

func absInt(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func getTotalSimilarityScore(values []int, frequencies map[int]int) int {
	var totalSimilarity int

	for _, element := range values {
		totalSimilarity += element * frequencies[element]
	}

	return totalSimilarity
}
