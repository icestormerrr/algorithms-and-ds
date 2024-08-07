package sorting

func merge(arr1, arr2 []int) []int {
	result := make([]int, 0, len(arr1)+len(arr2))
	var i, j = 0, 0

	for i < len(arr1) || j < len(arr2) {
		if i == len(arr1) {
			result = append(result, arr2[j])
			j++
			continue
		}
		if j == len(arr2) {
			result = append(result, arr1[i])
			i++
			continue
		}
		if arr1[i] < arr2[j] {
			result = append(result, arr1[i])
			i++
		} else {
			result = append(result, arr2[j])
			j++
		}
	}
	return result
}

func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	return merge(MergeSort(arr[:mid]), MergeSort(arr[mid:]))
}
