package tests

import (
	"testing"

	"github.com/akecel/gabbro/utils"
)

func TestCamelCaseToNormal(t *testing.T) {
	str := "HelloWorld"
	expected := "Hello world"
	result := utils.CamelCaseToNormal(str)

	if result != expected {
		t.Errorf("CamelCaseToNormal(%s) = %s; expected %s", str, result, expected)
	}

	str = "helloWorld"
	expected = "hello world"
	result = utils.CamelCaseToNormal(str)

	if result != expected {
		t.Errorf("CamelCaseToNormal(%s) = %s; expected %s", str, result, expected)
	}

	str = "already normal text"
	expected = "already normal text"
	result = utils.CamelCaseToNormal(str)

	if result != expected {
		t.Errorf("CamelCaseToNormal(%s) = %s; expected %s", str, result, expected)
	}
}

func TestContainsI(t *testing.T) {
	// Test with two strings that match
	strA := "Hello World"
	strB := "world"
	result := utils.ContainsI(strA, strB)
	if !result {
		t.Errorf("Expected ContainsI(%q, %q) to return true, but got false", strA, strB)
	}

	// Test with two strings that do not match
	strA = "Hello World"
	strB = "test"
	result = utils.ContainsI(strA, strB)
	if result {
		t.Errorf("Expected ContainsI(%q, %q) to return false, but got true", strA, strB)
	}

	// Test with empty strings
	strA = ""
	strB = ""
	result = utils.ContainsI(strA, strB)
	if !result {
		t.Errorf("Expected ContainsI(%q, %q) to return true, but got false", strA, strB)
	}
}
