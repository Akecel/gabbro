package responses

import (
	"github.com/akecel/gabbro/models"
)

type CharacterResponse struct {
	Name        string
	Description string
	Gender      models.CharacterGender
	Species     models.CharacterSpecies
	Games       string
	AKAS        string
	CountryName string
	URL         string
}
