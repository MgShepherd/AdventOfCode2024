package problems

import (
	"fmt"

	"github.com/MgShepherd/AdventOfCode2024/src/utils"
)

func SolveProblem1() (int, error) {
	data, err := utils.ReadProblemFile(1)
	if err != nil {
		return -1, err
	}
	fmt.Printf("Successfully read file with contents:\n\n%s\n", data)
	return 1, nil
}
