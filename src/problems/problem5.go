package problems

import (
	"slices"
	"strconv"
	"strings"

	"github.com/MgShepherd/AdventOfCode2024/src/utils"
)

func SolveProblem5() (int, error) {
	data, err := utils.ReadProblemFile(5)
	if err != nil {
		return 0, err
	}

	ruleLines, pageLines := []string{}, []string{}
	addingRules := true
	lines := strings.Split(data, "\n")

	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if addingRules && len(trimmedLine) == 0 {
			addingRules = false
		} else if addingRules && len(trimmedLine) > 0 {
			ruleLines = append(ruleLines, trimmedLine)
		} else if len(trimmedLine) > 0 {
			pageLines = append(pageLines, trimmedLine)
		}
	}

	rules := getOrderingRules(ruleLines)
	return getInvalidPagesSum(pageLines, rules), nil
}

func getOrderingRules(lines []string) map[string][]string {
	var rules = make(map[string][]string)

	for _, line := range lines {
		elements := strings.Split(line, "|")

		if val, ok := rules[elements[0]]; ok {
			rules[elements[0]] = append(val, elements[1])
		} else {
			rules[elements[0]] = []string{elements[1]}
		}
	}

	return rules
}

func getInvalidPagesSum(pages []string, rules map[string][]string) int {
	validSum := 0
	for _, page := range pages {
		elements := strings.Split(page, ",")
		if !isPageValid(elements, rules) {
			validPage := makePageValid(elements, rules)
			midPoint := len(validPage) / 2
			intVal, _ := strconv.Atoi(validPage[midPoint])
			validSum += intVal
		}
	}
	return validSum
}

func isPageValid(elements []string, rules map[string][]string) bool {
	invalidElements := []string{}

	for i := len(elements) - 1; i >= 0; i-- {
		if slices.Contains(invalidElements, elements[i]) {
			return false
		}
		invalidElements = append(invalidElements, rules[elements[i]]...)
	}

	return true
}

func makePageValid(elements []string, rules map[string][]string) []string {
	validPage := []string{}
	elementsToAdd := elements

	nextElementIdx := 0
	for len(validPage) < len(elements) {
		for i := 0; i < len(elementsToAdd); i++ {
			if slices.Contains(rules[elementsToAdd[i]], elementsToAdd[nextElementIdx]) {
				nextElementIdx = i
				i = 0
			}
		}

		validPage = append(validPage, elementsToAdd[nextElementIdx])
		elementsToAdd = slices.Delete(elementsToAdd, nextElementIdx, nextElementIdx+1)
		nextElementIdx = 0
	}

	return validPage
}
