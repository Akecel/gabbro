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
		t.Errorf("Result len=%s; Expected=%s", result, expected)
	}
}
