package main

import (
	"testing"
)

var values = []int{-2, 17, 52, 5, 3, 47, 68, 76, 39, 59, 66, 34, 80, 27, 64, 46, 6, 88, 91, 84, 61, 14, 71, 94, 20, 83, 1, 57, 12, 93, 4, 11, 70, 55, 36, 44, 31, 90, 15, 8, 43, 21, 16, 35, 62, 49, 9, 38, 85, 33, 67, 56, 75, 30, 78, 28, 7, 51, 50, 19, 41, 37, 79, 60, 86, 45, 73, 89, 32, 95, 58, 69, 92, 2, 10, 25, 23, 26, 65, 82, 74, 29, 77, 24, 72, 22, 48, 81, 87, 40, 53, 18, 54, 63, 0, 42, 96, 99, 98, 97, 100}

func BenchmarkSelectionSort(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		selectionSort(values)
	}
}
func BenchmarkRecursiveSelectionSort(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		selectionSortRecursive(values, 0)
	}
}
