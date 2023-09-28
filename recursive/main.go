package main

import (
	"fmt"
)

func main() {
	//Result: 25x25
	fmt.Println(findLargestSquareSizeHandle(25, 40))

	//Result: 20x20
	fmt.Println(findLargestSquareSizeHandle(20, 50))

	//Result: 3x3
	fmt.Println(findLargestSquareSizeHandle(33, 30))

	//Result: 200x200
	fmt.Println(findLargestSquareSizeHandle(200, 250))

	//Result: 200x200
	fmt.Println(findLargestSquareSizeHandle(200, 500))

	//Result: 250x250
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
