package problems

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/MgShepherd/AdventOfCode2024/src/utils"
)

func SolveProblem2() (int, error) {
	data, err := utils.ReadProblemFile(2)
	if err != nil {
		return -1, err
	}

	lines := strings.Split(data, "\n")

	totalSafeReports := 0
	for _, line := range lines {
		safe, err := isReportSafe(line)

		if err != nil {
			return 0, err
		}

		if safe {
			totalSafeReports += 1
		}
	}

	return totalSafeReports, nil
}

func isReportSafe(report string) (bool, error) {
	if len(strings.TrimSpace(report)) == 0 {
		return false, nil
	}
	levels, err := convertToIntSlice(strings.Fields(report))

	if err != nil {
		fmt.Printf("Unable to process report: %s\n", report)
		return false, err
	}

	var increasing = levels[0] < levels[1]

	for i := 1; i < len(levels); i++ {
		if !isLevelValid(levels[i-1], levels[i], increasing) {
			return false, nil
		}
	}

	return true, nil
}

func convertToIntSlice(elements []string) ([]int, error) {
	var intElements []int

	for _, element := range elements {
		intVal, err := strconv.Atoi(element)

		if err != nil {
			return intElements, err
		}

		intElements = append(intElements, intVal)
	}

	return intElements, nil
}

func isLevelValid(prev, current int, increasing bool) bool {
	if increasing {
		return current > prev && current <= prev+3
	}
	return current < prev && current >= prev-3
}
