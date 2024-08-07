package sorting

import "slices"

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	var pivotIndex = len(arr) / 2
	var left []int
	var right []int

	for i, value := range arr {
		if i == pivotIndex {
			continue
		}
		if value >= arr[pivotIndex] {
			right = append(right, value)
		} else if value < arr[pivotIndex] {
			left = append(left, value)
		}
	}

	return slices.Concat(QuickSort(left), []int{arr[pivotIndex]}, QuickSort(right))
}

func MutableQuickSort(arr []int) {
	mutableQuickSortRecursive(arr, 0, len(arr)-1)
}

func mutableQuickSortRecursive(arr []int, low, high int) {
	if low >= high {
		return
	}

	var pivotIndex = partition(arr, low, high)
	mutableQuickSortRecursive(arr, low, pivotIndex-1)
	mutableQuickSortRecursive(arr, pivotIndex+1, high)
}

func partition(arr []int, low, high int) int {
	var pivot = arr[low]
	var count int
	for i := low + 1; i <= high; i++ {
		if arr[i] <= pivot {
			count++
		}
	}
	var pivotIndex = low + count
	arr[low], arr[pivotIndex] = arr[pivotIndex], arr[low]

	var i = low
	var j = high
	for i < pivotIndex && j > pivotIndex {
		for arr[i] <= pivot {
			i++
		}
		for arr[j] > pivot {
			j--
		}
		if i < pivotIndex && j > pivotIndex {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
	}
	return pivotIndex
}
