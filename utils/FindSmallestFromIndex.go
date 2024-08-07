package utils

func FindSmallestFromIndex(arr []int, fromIndex int) int {
	var smallestElementIndex = fromIndex
	for i := fromIndex; i < len(arr); i++ {
		if arr[i] < arr[smallestElementIndex] {
			smallestElementIndex = i
		}
	}
	return smallestElementIndex
}
