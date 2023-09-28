package qsort

type PivotType int32

const (
	FirstElementPivot PivotType = iota
	LastElementPivot
	MiddleElementPivot
	RandomElementPivot
)

func QuickSort(values []int, pivotType PivotType) []int {
	switch {
	case pivotType == FirstElementPivot:
		return quickSortFirstElementPivot(values)
	case pivotType == LastElementPivot:
		return quickSortLastElementPivot(values)
	case pivotType == MiddleElementPivot:
		return quickSortMiddleElementPivot(values)
	case pivotType == RandomElementPivot:
		return quickSortRandomElementPivot(values)
	default:
		return quickSortRandomElementPivot(values)
	}
}
