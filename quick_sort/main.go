package main

import (
	"fmt"
	"math/rand"
	"quick_sort/qsort"
)

func main() {
	values := []int{3, -2, 5, 4, 1, 2, 10, 20, 8}
	// values := generateRandomArray(30000)
	fmt.Println(qsort.QuickSort(values, qsort.RandomElementPivot))
}

func generateRandomArray(size int) []int {
	arr := make([]int, size)
	usedNumbers := make(map[int]bool)

	for i := 0; i < size; i++ {
		var num int
		for {
			num = rand.Intn(size) - size/2
			if !usedNumbers[num] {
				usedNumbers[num] = true
				break
			}
		}
		arr[i] = num
	}

	return arr
}
