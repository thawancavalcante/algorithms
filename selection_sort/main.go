package main

import (
	"fmt"
)

func main() {
	values := []int{-2, 17, 52, 5, 3, 47, 68, 76, 39, 59, 66, 34, 80, 27, 64, 46, 6, 88, 91, 84, 61, 14, 71, 94, 20, 83, 1, 57, 12, 93, 4, 11, 70, 55, 36, 44, 31, 90, 15, 8, 43, 21, 16, 35, 62, 49, 9, 38, 85, 33, 67, 56, 75, 30, 78, 28, 7, 51, 50, 19, 41, 37, 79, 60, 86, 45, 73, 89, 32, 95, 58, 69, 92, 2, 10, 25, 23, 26, 65, 82, 74, 29, 77, 24, 72, 22, 48, 81, 87, 40, 53, 18, 54, 63, 0, 42, 96, 99, 98, 97, 100}
	result := selectionSort(values)

	fmt.Println(result)
}

func selectionSort(values []int) []int {
	valuesLength := len(values)
	newValues := make([]int, valuesLength)
	for index := 0; index < valuesLength; index++ {
		lowestNumberIndex := getLowestNumberIndex(values)
		newValues[index] = values[lowestNumberIndex]
		values = removeItemFromArray(values, lowestNumberIndex)
	}

	return newValues
}

func selectionSortRecursive(values []int, currentIndex int) []int {
	if currentIndex == len(values)-1 {
		return values
	}

	lowestIndex := currentIndex
	for index := currentIndex + 1; index < len(values); index++ {
		if values[currentIndex] < values[lowestIndex] {
			lowestIndex = currentIndex
		}
	}

	currentItem := values[currentIndex]
	lowestItem := values[lowestIndex]
	values[currentIndex] = lowestItem
	values[lowestIndex] = currentItem

	return selectionSortRecursive(values, currentIndex+1)
}

func removeItemFromArray(values []int, indexToRemove int) []int {
	values[indexToRemove] = values[len(values)-1]
	return values[:len(values)-1]
}

func getLowestNumberIndex(values []int) int {
	lowestNumberIndex := 0
	lowestNumber := values[lowestNumberIndex]

	for index, value := range values {
		if value < lowestNumber {
			lowestNumber = value
			lowestNumberIndex = index
		}
	}

	return lowestNumberIndex
}
