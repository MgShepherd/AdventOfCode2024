package utils

import "strconv"

func ConvertToIntSlice(elements []string) ([]int, error) {
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
