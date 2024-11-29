package utils

import (
	"fmt"
	"os"
)

func ReadProblemFile(problemNumber int) (string, error) {
	fileName := fmt.Sprintf("problems/problem%d.txt", problemNumber)

	contents, err := os.ReadFile(fileName)

	if err != nil {
		return "", err
	}

	return string(contents), nil
}
