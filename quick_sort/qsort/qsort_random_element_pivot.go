package qsort

import (
	"math/rand"
)

func quickSortRandomElementPivot(values []int) []int {
	if len(values) <= 1 {
		return values
	}

	pivotIndex := rand.Intn(len(values) - 1)
	pivotValue := values[pivotIndex]

	less := []int{}
	equal := []int{pivotValue}
	greater := []int{}

	valuesWithoutPivot := values[:pivotIndex]
	valuesWithoutPivot = append(valuesWithoutPivot, values[pivotIndex+1:]...)

	for _, value := range valuesWithoutPivot {
		switch {
		case value == pivotValue:
			equal = append(equal, value)
		case value > pivotValue:
			greater = append(greater, value)
		case value < pivotValue:
			less = append(less, value)
		}
	}

	result := []int{}
	result = append(result, quickSortRandomElementPivot(less)...)
	result = append(result, equal...)
	result = append(result, quickSortRandomElementPivot(greater)...)

	return result
}
