package models

type CharacterGender int

const (
	Male   CharacterGender = 0
	Female CharacterGender = 1
	Other  CharacterGender = 2
)

func (e CharacterGender) String() string {
	switch e {
	case Male:
		return "Male"
	case Female:
		return "Female"
	case Other:
		return "Other"
	default:
		return "Unknown"
	}
}
