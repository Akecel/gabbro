package tests

import (
	"testing"

	"github.com/akecel/gabbro/utils"
)

func TestReadableFloatNb(t *testing.T) {
	input := 90.22764927682539
	expectedOutput := 90.2
	result := utils.ReadableFloatNb(input)
	if result != expectedOutput {
		t.Errorf("Expected %f, but got %f", expectedOutput, result)
	}
}
