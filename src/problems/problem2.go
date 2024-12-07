package problems

import (
	"fmt"
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
	levels, err := utils.ConvertToIntSlice(strings.Fields(report))
	removedInvalid := false
	increasing := isIncreasing(levels)

	if err != nil {
		fmt.Printf("Unable to process report: %s\n", report)
		return false, err
	}

	for i := 0; i < len(levels)-1; i++ {
		if !isLevelValid(levels[i], levels[i+1], increasing) {
			if !removedInvalid {
				levels = append(levels[:i], levels[i+1:]...)
				i -= 1
				removedInvalid = true
			} else {
				return false, nil
			}
		}
	}

	return true, nil
}

func isIncreasing(levels []int) bool {
	return levels[0] < levels[1] && levels[1] < levels[2]
}

func isLevelValid(prev, current int, increasing bool) bool {
	if increasing {
		return current > prev && current <= prev+3
	}
	return current < prev && current >= prev-3
}
