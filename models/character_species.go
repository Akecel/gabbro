package models

type CharacterSpecies int

const (
	Human   CharacterSpecies = 1
	Alien   CharacterSpecies = 2
	Animal  CharacterSpecies = 3
	Android CharacterSpecies = 4
	Unknown CharacterSpecies = 5
)

func (e CharacterSpecies) String() string {
	switch e {
	case Human:
		return "Human"
	case Alien:
		return "Alien"
	case Animal:
		return "Animal"
	case Android:
		return "v"
	case Unknown:
		return "Unknown"
	default:
		return "Unknown"
	}
}
