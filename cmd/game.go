package cmd

import (
	"fmt"
	"strings"

	"github.com/akecel/gabbro/utils"
	"github.com/akecel/gabbro/data"

	"github.com/spf13/cobra"
)

// gameCmd represents the game command
var GameCmd = &cobra.Command{
	Use:   "game",
	Short: "Get game information by name",
	Long: `Get several informations about a game by using it name for searching in IGDB database.`,
	Run: GetGame,
}

func init() {
	RootCmd.AddCommand(GameCmd)
}

func GetGame(cmd *cobra.Command, args []string) {
	game := data.GetGameData(strings.Join(args, " "), 1)

	title := utils.SetColor()
	fmt.Printf("%s %s\n", title("Name:"), game[0].Name)
	fmt.Printf("%s %s\n", title("Description:"), game[0].Summary)
	fmt.Printf("%s %s\n", title("Release date:"), utils.ParseTimeStampToString(game[0].FirstReleaseDate))
	fmt.Printf("%s %s\n", title("URL:"), game[0].URL)
	fmt.Printf("%s %f\n", title("Rating:"), game[0].Rating)
}
