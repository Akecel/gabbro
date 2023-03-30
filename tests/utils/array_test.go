package tests

import (
	"testing"

	"github.com/akecel/gabbro/utils"
)

func TestJoinAndRemoveDuplicateStr(t *testing.T) {
	list := []string{"test", "test", "test2"}
	expected := "test, test2"
	result := utils.JoinAndRemoveDuplicateStr(list)

	if result != expected {
		t.Errorf("Result len=%s; Expected=%s", result, expected)
	}
}

func TestRemoveDuplicateStr(t *testing.T) {
	list := []string{"test", "test"}
	result := utils.RemoveDuplicateStr(list)

	if len(result) > 1 {
		t.Errorf("Result len=%d; Expected=%d", len(result), 1)
	}
}
