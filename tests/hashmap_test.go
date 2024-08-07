package tests

import (
	"go-training/data-structures/hashmap"
	"testing"
)

func TestGettingElementsAfterAddingElementsToHashMapMoreThanItsCapacity(t *testing.T) {
	var testKeys = []string{"sdjklfkjsdklf", "45645fgbx", "35345", "3", " "}
	var testMap = hashmap.NewHashMap(2)
	for i, key := range testKeys {
		testMap.Set(key, i)
	}

	for i, key := range testKeys {
		var value = testMap.Get(key)
		if value != i {
			t.Errorf("For key %s, expected %d , but got %d ", key, i, value)
		}
	}
}

func TestGettingValueByNonExistentKeyInHashMap(t *testing.T) {
	var testMap = hashmap.NewHashMap(5)
	var nonExistentKey = "nonExistentKey"

	var value = testMap.Get(nonExistentKey)
	if value != 0 {
		t.Errorf("For key %s, expected %d , but got %d ", nonExistentKey, -1, value)
	}
}

func TestGettingValueByEmptyHashMap(t *testing.T) {
	var testMap = hashmap.NewHashMap(0)
	var someKey = "someKey"

	var value = testMap.Get(someKey)
	if value != 0 {
		t.Errorf("For key %s, expected %d , but got %d ", someKey, -1, value)
	}
}

func TestGetLengthWhenSizeEqualsLength(t *testing.T) {
	var testKeys = []string{"sdjklfkjsdklf", "45645fgbx", "35345", "3", " "}
	var length = len(testKeys)
	var size = length
	var testMap = hashmap.NewHashMap(size)
	for i, key := range testKeys {
		testMap.Set(key, i)
	}

	var testLength = testMap.GetLength()
	if testLength != length {
		t.Errorf("Expected length = %d , but got %d ", length, testLength)
	}
}

func TestGetLengthWhenSizeNotEqualsLength(t *testing.T) {
	var testKeys = []string{"sdjklfkjsdklf", "45645fgbx", "35345", "3", " "}
	var size = 1
	var length = len(testKeys)
	var testMap = hashmap.NewHashMap(size)
	for i, key := range testKeys {
		testMap.Set(key, i)
	}

	var testLength = testMap.GetLength()
	if testLength != length {
		t.Errorf("Expected length = %d , but got %d ", length, testLength)
	}
}

func TestAddingEmptyStringAsKeyInHashMap(t *testing.T) {
	var testMap = hashmap.NewHashMap(0)
	testMap.Set("", 999)
	if testMap.GetLength() != 0 {
		t.Error("Added empty string key, but length increased")
	}
}

// mb should throw err
func TestGettingEmptyStringAsKeyInHashMap(t *testing.T) {
	var testMap = hashmap.NewHashMap(0)

	var testValue = testMap.Get("")
	if testValue != 0 {
		t.Errorf("Expected %d , but got %d ", 0, testValue)
	}
}

// mb should throw err
func TestSettingAlreadyExistingKeyInHashMap(t *testing.T) {
	var testMap = hashmap.NewHashMap(2)
	var someKey = "someKey"
	var firstValue = 1
	var secondValue = 1

	testMap.Set(someKey, firstValue)
	testMap.Set(someKey, secondValue)
	var value = testMap.Get(someKey)
	if value != secondValue {
		t.Errorf("For key %s, expected %d , but got %d ", someKey, secondValue, value)
	}
}

func TestSettingValuesWithZeroSizeHashMap(t *testing.T) {
	var testMap = hashmap.NewHashMap(0)
	var someKey = "someKey"
	var someValue = 1

	testMap.Set(someKey, someValue)
}
