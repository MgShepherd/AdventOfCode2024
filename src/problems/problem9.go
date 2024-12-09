package problems

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/MgShepherd/AdventOfCode2024/src/utils"
)

type Range struct {
	value string
	start int
	size  int
}

func SolveProblem9() (int, error) {
	data, err := utils.ReadProblemFile(9)
	if err != nil {
		return 0, err
	}

	data = strings.TrimSpace(data)
	generatedMemory, freeSpaceRanges, elementRanges := generateMemory(strings.Split(data, ""))
	sortMemory(generatedMemory, freeSpaceRanges, elementRanges)
	return computeChecksum(generatedMemory), nil
}

func generateMemory(data []string) ([]string, []Range, []Range) {
	currentId := 0
	isFile := true
	decompressed := []string{}
	freeSpaceRanges := []Range{}
	elementRanges := []Range{}

	for _, item := range data {
		numRepeatedTimes, err := strconv.Atoi(item)
		if err != nil {
			fmt.Println("Unable to process item", item)
		}

		if isFile {
			stringId := strconv.Itoa(currentId)
			elementRange := Range{value: stringId, start: len(decompressed), size: numRepeatedTimes}
			elementRanges = append(elementRanges, elementRange)
			decompressed = appendItem(decompressed, stringId, numRepeatedTimes)
			currentId += 1
		} else {
			freeRange := Range{value: ".", start: len(decompressed), size: numRepeatedTimes}
			freeSpaceRanges = append(freeSpaceRanges, freeRange)
			decompressed = appendItem(decompressed, ".", numRepeatedTimes)
		}
		isFile = !isFile
	}

	return decompressed, freeSpaceRanges, elementRanges
}

func sortMemory(data []string, freeSpaceRanges, elementRanges []Range) {
	for i := len(elementRanges) - 1; i >= 0; i-- {
		for j := 0; j < len(freeSpaceRanges); j++ {
			if freeSpaceRanges[j].start < elementRanges[i].start && freeSpaceRanges[j].size >= elementRanges[i].size {
				insertValue(data, elementRanges[i].value, freeSpaceRanges[j].start, elementRanges[i].start, elementRanges[i].size)
				freeSpaceRanges[j].start = freeSpaceRanges[j].start + elementRanges[i].size
				freeSpaceRanges[j].size = freeSpaceRanges[j].size - elementRanges[i].size
				break
			}
		}
	}
}

func insertValue(data []string, element string, newStartPos, oldStartPos, size int) {
	for i := 0; i < size; i++ {
		data[newStartPos+i] = element
		data[oldStartPos+i] = "."
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

func appendItem(data []string, item string, numTimes int) []string {
	for i := 0; i < numTimes; i++ {
		data = append(data, item)
	}
	return data
}
