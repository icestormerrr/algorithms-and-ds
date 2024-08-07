package tests

import (
	"go-training/algorithms/search"
	"testing"
)

type SearchTest struct {
	inputArr      []int
	inputVal      int
	expectedIndex int
}

const badResult = -1

var unsortedSearchTests = []SearchTest{
	{[]int{}, 1, badResult},
	{[]int{42}, 42, 0},
	{[]int{1, 2, 3, 4, 5}, 6, badResult},
	{[]int{5, 2, 7, 2, 5, 3, 7}, 7, 2},
}

var sortedSearchTests = []SearchTest{
	{[]int{}, 1, badResult},
	{[]int{42}, 42, 0},
	{[]int{1, 2, 3, 4, 5}, 6, badResult},
	{[]int{2, 2, 3, 5, 5, 7, 7}, 7, 5},
}

func createSearchTestForUnsortedArray(SearchFunc func(arr []int, target int) int) func(t *testing.T) {
	return func(t *testing.T) {
		for _, test := range unsortedSearchTests {
			result := SearchFunc(test.inputArr, test.inputVal)
			if result != test.expectedIndex {
				t.Errorf("For input %v, expected %d, but got %d", test.inputArr, test.expectedIndex, result)
			}
		}
	}
}

func createSearchTestForSortedArray(SearchFunc func(arr []int, target int) int) func(t *testing.T) {
	return func(t *testing.T) {
		for _, test := range sortedSearchTests {
			result := SearchFunc(test.inputArr, test.inputVal)
			if result != test.expectedIndex {
				t.Errorf("For input %v, expected %d, but got %d", test.inputArr, test.expectedIndex, result)
			}
		}
	}
}

func TestBinarySearch(t *testing.T) {
	var TestFunc = createSearchTestForSortedArray(search.BinarySearch)
	TestFunc(t)
}
