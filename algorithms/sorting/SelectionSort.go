package sorting

import (
	"go-training/utils"
)

func SelectionSort(arr []int) []int {
	var sortedArr = make([]int, len(arr))
	copy(sortedArr, arr)
	for i := range sortedArr {
		var smallestElementIndex = utils.FindSmallestFromIndex(sortedArr, i)
		if smallestElementIndex != i {
			sortedArr[i], sortedArr[smallestElementIndex] = sortedArr[smallestElementIndex], sortedArr[i]
		}
	}
	return sortedArr
}
