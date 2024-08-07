package tests

import (
	"go-training/data-structures/linkedlist"
	"reflect"
	"slices"
	"testing"
)

func TestPushingNewValuesToList(t *testing.T) {
	var expected = []int{1, 2, 3, 4, 5}
	var list = linkedlist.NewLinkedList()
	for _, v := range expected {
		list.Push(v)
	}

	var result = list.ToArray()
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected: %v, but got: %v", expected, result)
	}
}

func TestGetLengthAfterPushingElements(t *testing.T) {
	var expected = []int{1, 2, 3, 4, 5}
	var list = linkedlist.NewLinkedList()
	for _, v := range expected {
		list.Push(v)
	}

	if list.GetLength() != len(expected) {
		t.Errorf("Expected: %v, but got: %v", len(expected), list.GetLength())
	}
}

func TestGettingElementsByIndex(t *testing.T) {
	var expected = []int{1, 2, 3, 4, 5}
	var list = linkedlist.NewLinkedList()
	for _, v := range expected {
		list.Push(v)
	}

	for i, v := range expected {
		var result = list.Get(i).Value
		if result != v {
			t.Errorf("Expected: %d, but got: %d", v, result)
		}
	}
}

func TestGettingElementsByNonExistentIndex(t *testing.T) {
	var list = linkedlist.NewLinkedList()
	list.Push(9)
	list.Push(99)

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but no panic occurred")
		}
	}()
	list.Get(3)
}

func TestGettingElementsByNegativeIndex(t *testing.T) {
	var list = linkedlist.NewLinkedList()
	list.Push(9)
	list.Push(99)

	var value = list.Get(-10).Value
	if value != 9 {
		t.Errorf("Expected: %d (first element), but got: %d", 9, value)
	}
}

func TestInsertingNewValuesToList(t *testing.T) {
	var source = []int{1, 3, 5}
	var list = linkedlist.NewLinkedList()
	for _, v := range source {
		list.Push(v)
	}

	list.Insert(2, 0)
	list.Insert(4, 2)

	var expected = []int{1, 2, 3, 4, 5}
	var result = list.ToArray()
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected: %v, but got: %v", expected, result)
	}
}

func TestInsertingAfterNonExistentIndex(t *testing.T) {
	var list = linkedlist.NewLinkedList()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but no panic occurred")
		}
	}()
	list.Insert(1, 0)
}

func TestGetLengthAfterInsertingElements(t *testing.T) {
	var list = linkedlist.NewLinkedList()
	list.Push(1)
	list.Insert(2, 0)
	list.Insert(3, 1)
	list.Insert(4, 2)
	list.Insert(5, 3)

	var expected = []int{1, 2, 3, 4, 5}
	if list.GetLength() != len(expected) {
		t.Errorf("Expected: %v, but got: %v", len(expected), list.GetLength())
	}
}

func TestDeletingMiddleElement(t *testing.T) {
	var source = []int{1, 2, 3, 4, 5}
	var list = linkedlist.NewLinkedList()
	for _, v := range source {
		list.Push(v)
	}

	list.Delete(2)

	var expected = []int{1, 2, 4, 5}
	if !reflect.DeepEqual(expected, list.ToArray()) {
		t.Errorf("Expected: %v, but got: %v", expected, list.ToArray())
	}

	slices.Reverse(expected)
	if !reflect.DeepEqual(expected, list.ToReverseArray()) {
		t.Errorf("Expected: %v, but got: %v", expected, list.ToReverseArray())
	}
}

func TestDeletingFirstElement(t *testing.T) {
	var source = []int{1, 2, 3, 4, 5}
	var list = linkedlist.NewLinkedList()
	for _, v := range source {
		list.Push(v)
	}

	list.Delete(0)

	var expected = []int{2, 3, 4, 5}
	if !reflect.DeepEqual(expected, list.ToArray()) {
		t.Errorf("Expected: %v, but got: %v", expected, list.ToArray())
	}

	slices.Reverse(expected)
	if !reflect.DeepEqual(expected, list.ToReverseArray()) {
		t.Errorf("Expected: %v, but got: %v", expected, list.ToReverseArray())
	}
}

func TestDeletingLastElement(t *testing.T) {
	var source = []int{1, 2, 3, 4, 5}
	var list = linkedlist.NewLinkedList()
	for _, v := range source {
		list.Push(v)
	}

	list.Delete(4)

	var expected = []int{1, 2, 3, 4}
	if !reflect.DeepEqual(expected, list.ToArray()) {
		t.Errorf("Expected: %v, but got: %v", expected, list.ToArray())
	}

	slices.Reverse(expected)
	if !reflect.DeepEqual(expected, list.ToReverseArray()) {
		t.Errorf("Expected: %v, but got: %v", expected, list.ToReverseArray())
	}
}

func TestDeletingNonExistentIndex(t *testing.T) {
	var list = linkedlist.NewLinkedList()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but no panic occurred")
		}
	}()
	list.Delete(0)
}
