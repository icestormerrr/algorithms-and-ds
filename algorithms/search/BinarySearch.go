package search

func BinarySearch(arr []int, target int) int {
	var left, right = 0, len(arr) - 1
	var mid int
	for left <= right {
		mid = (right + left) / 2
		if arr[mid] < target {
			left = mid + 1
			continue
		}
		if arr[mid] > target {
			right = mid - 1
			continue
		}
		return mid
	}
	return -1
}
