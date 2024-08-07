package utils

import "slices"

func DeleteByIndex(slice []int, targetIndex int) []int {
	//var newSlice = make([]int, len(slice)-1)
	//for i := 0; i < len(slice); i++ {
	//	if i < targetIndex {
	//		newSlice[i] = slice[i]
	//		continue
	//	}
	//	if i > targetIndex {
	//		newSlice[i-1] = slice[i]
	//	}
	//}
	//return newSlice

	return slices.Concat(slice[:targetIndex], slice[targetIndex+1:])
}
