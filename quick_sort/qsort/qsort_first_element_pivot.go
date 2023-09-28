package qsort

func quickSortFirstElementPivot(values []int) []int {
	if len(values) <= 1 {
		return values
	}

	pivotIndex := 0
	pivotValue := values[pivotIndex]
	less := []int{}
	equal := []int{pivotValue}
	greater := []int{}

	for _, value := range values[1:] {
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
	result = append(result, quickSortFirstElementPivot(less)...)
	result = append(result, equal...)
	result = append(result, quickSortFirstElementPivot(greater)...)

	return result
}
