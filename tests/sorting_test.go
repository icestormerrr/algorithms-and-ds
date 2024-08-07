package tests

import (
	"go-training/algorithms/sorting"
	"reflect"
	"testing"
)

type SortTest struct {
	input    []int
	expected []int
}

var sortTests = []SortTest{
	{[]int{}, []int{}},
	{[]int{42}, []int{42}},
	{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
	{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
	{[]int{5, 2, 7, 2, 5, 3, 7}, []int{2, 2, 3, 5, 5, 7, 7}},
	{[]int{9, -3, 5, -7, 2}, []int{-7, -3, 2, 5, 9}},
}

func createSortTest(SortFunc func(arr []int) []int) func(t *testing.T) {
	return func(t *testing.T) {
		for _, test := range sortTests {
			result := SortFunc(test.input)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("For input %v, expected %v, but got %v", test.input, test.expected, result)
			}
		}
	}
}

func TestMergeSort(t *testing.T) {
	var TestFunc = createSortTest(sorting.MergeSort)
	TestFunc(t)
}

func TestQuickSort(t *testing.T) {
	var TestFunc = createSortTest(sorting.QuickSort)
	TestFunc(t)
}

func TestSelectionSort(t *testing.T) {
	var TestFunc = createSortTest(sorting.SelectionSort)
	TestFunc(t)
}
