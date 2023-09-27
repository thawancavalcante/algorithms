package main

import (
	"fmt"
)

func main() {
	values := []int{-3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100}
	if found := binarySearch(values, 6); !found {
		fmt.Println("Value not found")
		return
	}

	fmt.Println("Value found!")
}

func binarySearch(values []int, searchValue int) bool {
	lowerIndex := 0
	highIndex := len(values) - 1

	for {
		middleIndex := (lowerIndex + highIndex) >> 1

		if lowerIndex > highIndex {
			return false
		}

		middleItem := values[middleIndex]

		if middleItem == searchValue {
			return true
		}

		if middleItem > searchValue {
			highIndex = middleIndex - 1
			continue
		}

		lowerIndex = middleIndex + 1
	}
}

func recursiveBinarySearch(values []int, searchValue int, lowerIndex int, highIndex int) bool {
	if lowerIndex > highIndex {
		return false
	}

	middleIndex := (lowerIndex + highIndex) >> 1
	middleItem := values[middleIndex]

	if middleItem == searchValue {
		return true
	}

	if middleItem > searchValue {
		return recursiveBinarySearch(values, searchValue, lowerIndex, middleIndex-1)
	}

	return recursiveBinarySearch(values, searchValue, middleIndex+1, highIndex)
}
