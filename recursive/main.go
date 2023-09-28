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

	maximumParts := int(bigSize / smallSize)
	newSize := bigSize - smallSize*maximumParts

	if newSize > smallSize {
		return findLargestSquareSize(newSize, smallSize)
	}

	return findLargestSquareSize(smallSize, newSize)
}
