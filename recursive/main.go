package main

import (
	"fmt"
)

func main() {
	fmt.Println(findLargestSquareSizeHandle(25, 40))
	fmt.Println(findLargestSquareSizeHandle(20, 50))
	fmt.Println(findLargestSquareSizeHandle(33, 30))
	fmt.Println(findLargestSquareSizeHandle(200, 250))
	fmt.Println(findLargestSquareSizeHandle(200, 500))
	fmt.Println(findLargestSquareSizeHandle(300, 250))

	values := []int{1, 4, 10, 100, 2, 3, -1, -50, 5, 99, 10, 2}
	fmt.Println(sum(values, 0))
}

func findLargestSquareSizeHandle(horizontalSize int, verticalSize int) (int, int) {
	if horizontalSize > verticalSize {
		return findLargestSquareSize(horizontalSize, verticalSize)
	}

	return findLargestSquareSize(verticalSize, horizontalSize)
}

func findLargestSquareSize(bigSize int, smallSize int) (int, int) {
	if bigSize%smallSize == 0 {
		return smallSize, smallSize
	}

	maximumParts := bigSize / smallSize
	newSize := bigSize - smallSize*maximumParts

	if newSize > smallSize {
		return findLargestSquareSize(newSize, smallSize)
	}

	return findLargestSquareSize(smallSize, newSize)
}

func sum(values []int, total int) int {
	if len(values) == 0 {
		return total
	}

	return sum(values[1:], values[0]+total)
}
