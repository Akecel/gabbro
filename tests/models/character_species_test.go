package tests

import (
	"github.com/akecel/gabbro/models"
	"testing"
)

func TestCharacterSpecies_String(t *testing.T) {
	testCases := []struct {
		species  models.CharacterSpecies
		expected string
	}{
		{1, "Human"},
		{2, "Alien"},
		{3, "Animal"},
		{4, "Android"},
		{5, "Unknown"},
		{models.CharacterSpecies(99), "Unknown"},
	}

	for _, tc := range testCases {
		actual := tc.species.String()
		if actual != tc.expected {
			t.Errorf("Expected %s but got %s for species %d", tc.expected, actual, tc.species)
		}
	}
}
