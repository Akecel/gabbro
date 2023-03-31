package cmd

import (
	"strings"

	"github.com/akecel/gabbro/data"
	"github.com/akecel/gabbro/models"
	"github.com/akecel/gabbro/responses"
	"github.com/akecel/gabbro/utils"

	"github.com/spf13/cobra"
)

var CharacterCmd = &cobra.Command{
	Use:   "character",
	Short: "Get character informations",
	Long:  "Get several informations about a game character by using it name for searching in IGDB database.",
	Run:   GetCharacter,
}

func init() {
	RootCmd.AddCommand(CharacterCmd)
}

func GetCharacter(cmd *cobra.Command, args []string) {
	character := data.GetCharacterDataByName(strings.Join(args, " "), 1)[0]
	characterMugShot := data.GetMugShotsDataByIDs(character.MugShot, 1)
	characterMugShotURL := utils.ReconstructImgURL(characterMugShot.URL)

	var games []string
	characterHaveGame := len(character.Games) > 0
	if characterHaveGame {
		characterGames := data.GetGamesDataByIDs(character.Games, len(character.Games))
		for i := 0; i < len(characterGames); i++ {
			games = append(games, characterGames[i].Name)
		}
	}

	response := responses.CharacterResponse{
		Name:        character.Name,
		Description: character.Description,
		Gender:      models.CharacterGender(int(character.Gender)),
		Species:     models.CharacterSpecies(int(character.Species)),
		Games:       utils.JoinAndRemoveDuplicateStr(games),
		AKAS:        strings.Join(character.AKAS, ", "),
		CountryName: character.CountryName,
		URL:         character.URL,
	}

	responses.PrintImageResponse(characterMugShotURL)
	responses.PrintResponse(response)
}
