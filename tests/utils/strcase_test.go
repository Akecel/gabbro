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
