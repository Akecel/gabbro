package cmd

import (
	"strings"

	"github.com/akecel/gabbro/data"
	"github.com/akecel/gabbro/responses"
	"github.com/akecel/gabbro/utils"

	"github.com/spf13/cobra"
)

var GameCmd = &cobra.Command{
	Use:   "game",
	Short: "Get game informations",
	Long:  `Get several informations about a game by using it name for searching in IGDB database.`,
	Run:   GetGame,
}

func init() {
	RootCmd.AddCommand(GameCmd)
}

func GetGame(cmd *cobra.Command, args []string) {
	game := data.GetGamesDataByName(strings.Join(args, " "), 1)[0]
	gameInvolvedCompanies := data.GetInvolvedCompaniesDataByIDs(game.InvolvedCompanies, len(game.InvolvedCompanies))
	gameCover := data.GetCoverDataByIDs(game.Cover, 1)
	gameCoverURL := utils.ReconstructCoverURL(gameCover.URL)

	var companies []string
	for i := 0; i < len(gameInvolvedCompanies); i++ {
		company := data.GetCompaniesDataByID(gameInvolvedCompanies[i].Company, 1)
		companies = append(companies, company.Name)
	}

	var dlcs []string
	gameHaveDLCs := len(game.DLCS) > 0
	if gameHaveDLCs {
		gameDLCs := data.GetGamesDataByIDs(game.DLCS, len(game.DLCS))
		for i := 0; i < len(gameDLCs); i++ {
			dlcs = append(dlcs, gameDLCs[i].Name)
		}
	}

	var genres []string
	gameHaveGenres := len(game.Genres) > 0
	if gameHaveGenres {
		gameGenres := data.GetGenresDataByIDs(game.Genres, len(game.Genres))
		for i := 0; i < len(gameGenres); i++ {
			genres = append(genres, gameGenres[i].Name)
		}
	}

	var themes []string
	gameHaveThemes := len(game.Themes) > 0
	if gameHaveThemes {
		gameThemes := data.GetThemesDataByIDs(game.Themes, len(game.Themes))
		for i := 0; i < len(gameThemes); i++ {
			themes = append(themes, gameThemes[i].Name)
		}
	}

	var platforms []string
	gameHavePlatforms := len(game.Platforms) > 0
	if gameHavePlatforms {
		gamePlatfroms := data.GetPlatformsDataByIDs(game.Platforms, len(game.Platforms))
		for i := 0; i < len(gamePlatfroms); i++ {
			platforms = append(platforms, gamePlatfroms[i].Name)
		}
	}

	response := responses.GameResponse{
		Name:        game.Name,
		Description: game.Summary,
		Genres:      utils.JoinAndRemoveDuplicateStr(genres),
		Themes:      utils.JoinAndRemoveDuplicateStr(themes),
		DLCs:        utils.JoinAndRemoveDuplicateStr(dlcs),
		Companies:   utils.JoinAndRemoveDuplicateStr(companies),
		Platforms:   utils.JoinAndRemoveDuplicateStr(platforms),
		ReleaseDate: utils.ParseTimeStampToString(game.FirstReleaseDate),
		URL:         game.URL,
		Rating:      game.Rating,
	}

	responses.PrintImageResponse(gameCoverURL)
	responses.PrintResponse(response)
}
