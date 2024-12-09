package main

import (
	"fmt"

	"github.com/MgShepherd/AdventOfCode2024/src/problems"
)

func main() {
	result, err := problems.SolveProblem9()
	if err != nil {
		fmt.Printf("Unable to solve problem due to following error:\n%s", err)
		return
	}

	fmt.Println("Solution to Problem is:", result)
}
