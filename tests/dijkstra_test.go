package tests

import (
	"go-training/data-structures/graph"
	"testing"
)

func TestDijkstraAlgorithm(t *testing.T) {
	var testMatrix1 = [][]int{
		{-1, 5, -1, -1, -1, 2},
		{-1, -1, 4, -1, 2, -1},
		{-1, -1, -1, 3, 6, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, 1, -1, -1},
		{-1, 8, -1, -1, 7, -1},
	}
	var testMatrix2 = [][]int{
		{0, 1, 1},
		{4, 0, 1},
		{2, 1, 0},
	}

	var testGraph = graph.NewGraph()
	testGraph.Build(testMatrix1)
	var costsMap = testGraph.FindCostsToVerticesFrom(0)
	if costsMap[3] != 8 {
		t.Error("Expected 8, got ", costsMap[3])
	}

	testGraph.Build(testMatrix2)
	costsMap = testGraph.FindCostsToVerticesFrom(1)
	if costsMap[0] != 3 {
		t.Error("Expected 3, got ", costsMap[0])
	}
}
