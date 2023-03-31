package tests

import (
	"testing"
	"github.com/akecel/gabbro/models"
)

func TestCharacterGender_String(t *testing.T) {
	testCases := []struct {
		gender   models.CharacterGender
		expected string
	}{
		{0, "Male"},
		{1, "Female"},
		{2, "Other"},
		{models.CharacterGender(99), "Unknown"},
	}

	for _, tc := range testCases {
		actual := tc.gender.String()
		if actual != tc.expected {
			t.Errorf("Expected %s but got %s for gender %d", tc.expected, actual, tc.gender)
		}
	}
}