package main

import (
	"fmt"
	"go-training/algorithms/sorting"
)

func main() {
	//var testMatrix1 = [][]int{
	//	{-1, 5, -1, -1, -1, 2},
	//	{-1, -1, 4, -1, 2, -1},
	//	{-1, -1, -1, 3, 6, -1},
	//	{-1, -1, -1, -1, -1, -1},
	//	{-1, -1, -1, 1, -1, -1},
	//	{-1, 8, -1, -1, 7, -1},
	//}
	//
	//var testGraph = graph.NewGraph()
	//testGraph.Build(testMatrix1)

	var arr = []int{9, 3, 4, 2, 1, 8}

	sorting.MutableQuickSort(arr)

	fmt.Println(arr)
}
