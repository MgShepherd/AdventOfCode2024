package problems

import (
	"fmt"
	"sort"
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

	return calcTotalDistance(locationLists), nil
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

	for i := 0; i < len(lists); i++ {
		sort.Ints(lists[i])
	}

	return lists, nil
}

func calcTotalDistance(locationLists [2][]int) int {
	var totalDistance int

	for i := 0; i < len(locationLists[0]); i++ {
		totalDistance += absInt(locationLists[0][i] - locationLists[1][i])
	}
	return totalDistance
}

func absInt(value int) int {
	if value < 0 {
		return -value
	}
	return value
}
