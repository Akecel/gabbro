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
	Use:   "character [Character name]",
	Short: "Get character informations",
	Long:  "Get informations about a game character by using it name for searching in IGDB database.",
	Args:  cobra.MinimumNArgs(1),
	Run:   GetCharacter,
}

func init() {
	RootCmd.AddCommand(CharacterCmd)
}

func GetCharacter(cmd *cobra.Command, args []string) {
	charactersList, err := data.GetCharacterDataByName(strings.Join(args, " "), 1); if err != nil {
		responses.PrintErrorResponse("Character not found")
	}

	character := charactersList[0]

	var games []string
	characterHaveGame := len(character.Games) > 0
	if characterHaveGame {
		characterGames, _ := data.GetGamesDataByIDs(character.Games, len(character.Games))
		for i := 0; i < len(characterGames); i++ {
			games = append(games, characterGames[i].Name)
		}
	}

	var characterMugShotURL string
	if WithImage {
		characterMugShot, err := data.GetMugShotsDataByID(character.MugShot); if err == nil {
			characterMugShotURL = utils.ReconstructImgURL(characterMugShot.URL)
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

	if len(characterMugShotURL) > 0 {
		responses.PrintImageResponse(characterMugShotURL)
	}
	
	responses.PrintResponse(response)
}
