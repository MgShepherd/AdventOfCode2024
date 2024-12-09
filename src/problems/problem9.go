package problems

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/MgShepherd/AdventOfCode2024/src/utils"
)

func SolveProblem9() (int, error) {
	data, err := utils.ReadProblemFile(9)
	if err != nil {
		return 0, err
	}

	data = strings.TrimSpace(data)
	generatedMemory := generateMemory(strings.Split(data, ""))
	sortMemory(generatedMemory)
	return computeChecksum(generatedMemory), nil
}

func generateMemory(data []string) []string {
	currentId := 0
	isFile := true
	decompressed := []string{}

	for _, item := range data {
		numRepeatedTimes, err := strconv.Atoi(item)
		if err != nil {
			fmt.Println("Unable to process item", item)
		}

		if isFile {
			decompressed = appendItem(decompressed, strconv.Itoa(currentId), numRepeatedTimes)
			currentId += 1
		} else {
			decompressed = appendItem(decompressed, ".", numRepeatedTimes)
		}
		isFile = !isFile
	}

	return decompressed
}

func sortMemory(data []string) {
	endPointer := len(data) - 1

	for i := 0; i < len(data); i++ {
		if data[i] == "." {
			endPointer = getNextEndPointer(data, endPointer)
			if i > endPointer {
				return
			}
			swapItems(data, i, endPointer)
			endPointer -= 1
		}
	}
}

func computeChecksum(data []string) int {
	checksum := 0

	for i := 0; i < len(data); i++ {
		if data[i] != "." {
			intElement, err := strconv.Atoi(data[i])
			if err != nil {
				fmt.Println("Unable to process element", data[i])
			}
			checksum += intElement * i
		}
	}

	return checksum
}

func getNextEndPointer(data []string, currentEndPointer int) int {
	for i := currentEndPointer; i >= 0; i-- {
		if data[i] != "." {
			return i
		}
	}
	return -1
}

func swapItems(data []string, index1 int, index2 int) {
	data[index1] = data[index2]
	data[index2] = "."
}

func appendItem(data []string, item string, numTimes int) []string {
	for i := 0; i < numTimes; i++ {
		data = append(data, item)
	}
	return data
}
