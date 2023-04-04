package tests

import (
	"reflect"
	"testing"

	"github.com/akecel/gabbro/utils"
)

func TestJoinAndRemoveDuplicateStr(t *testing.T) {
	// Test case 1: Test JoinAndRemoveDuplicateStr with an empty slice
	result := utils.JoinAndRemoveDuplicateStr([]string{})
	expectedResult := ""
	if result != expectedResult {
		t.Errorf("Test case 1 failed - expected: %s, got: %s", expectedResult, result)
	}

	// Test case 2: Test JoinAndRemoveDuplicateStr with a slice containing duplicates
	result = utils.JoinAndRemoveDuplicateStr([]string{"foo", "bar", "baz", "foo"})
	expectedResult = "foo, bar, baz"
	if result != expectedResult {
		t.Errorf("Test case 2 failed - expected: %s, got: %s", expectedResult, result)
	}

	// Test case 3: Test JoinAndRemoveDuplicateStr with a slice containing no duplicates
	result = utils.JoinAndRemoveDuplicateStr([]string{"foo", "bar", "baz"})
	expectedResult = "foo, bar, baz"
	if result != expectedResult {
		t.Errorf("Test case 3 failed - expected: %s, got: %s", expectedResult, result)
	}
}

func TestRemoveDuplicateStr(t *testing.T) {
	// Test case 1: Test RemoveDuplicateStr with an empty slice
	result := utils.RemoveDuplicateStr([]string{})
	expectedResult := []string{}
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Test case 1 failed - expected: %v, got: %v", expectedResult, result)
	}

	// Test case 2: Test RemoveDuplicateStr with a slice containing duplicates
	result = utils.RemoveDuplicateStr([]string{"foo", "bar", "baz", "foo"})
	expectedResult = []string{"foo", "bar", "baz"}
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Test case 2 failed - expected: %v, got: %v", expectedResult, result)
	}

	// Test case 3: Test RemoveDuplicateStr with a slice containing no duplicates
	result = utils.RemoveDuplicateStr([]string{"foo", "bar", "baz"})
	expectedResult = []string{"foo", "bar", "baz"}
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Test case 3 failed - expected: %v, got: %v", expectedResult, result)
	}
}

func TestArrayContains(t *testing.T) {
	// Test with an array that contains the element
	slice := []string{"apple", "banana", "cherry"}
	result := utils.ArrayContains(slice, "banana")
	if !result {
		t.Errorf("Expected ArrayContains(slice, 'banana') to return true, but got false")
	}

	// Test with an array that does not contain the element
	result = utils.ArrayContains(slice, "orange")
	if result {
		t.Errorf("Expected ArrayContains(slice, 'orange') to return false, but got true")
	}

	// Test with an empty array
	emptySlice := []string{}
	result = utils.ArrayContains(emptySlice, "test")
	if result {
		t.Errorf("Expected ArrayContains(emptySlice, 'test') to return false, but got true")
	}
}
