package main

import (
	"quick_sort/qsort"
	"testing"
)

var data = generateRandomArray(100000)

func benchmarkQuickSort(b *testing.B, pivotType qsort.PivotType) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		qsort.QuickSort(data, pivotType)
	}
}

func BenchmarkQuickSortFirstElementPivot(b *testing.B) {
	benchmarkQuickSort(b, qsort.FirstElementPivot)
}

func BenchmarkQuickSortLastElementPivot(b *testing.B) {
	benchmarkQuickSort(b, qsort.LastElementPivot)
}

func BenchmarkQuickSortMiddleElementPivot(b *testing.B) {
	benchmarkQuickSort(b, qsort.MiddleElementPivot)
}

func BenchmarkQuickSortRandomElementPivot(b *testing.B) {
	benchmarkQuickSort(b, qsort.RandomElementPivot)
}
