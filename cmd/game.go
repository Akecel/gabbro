package cmd

import (
	"fmt"
	"strings"
	"log"

	"github.com/akecel/gabbro/config"
	"github.com/akecel/gabbro/utils"

	"github.com/Henry-Sarabia/igdb/v2"
	"github.com/spf13/cobra"
)

// gameCmd represents the game command
var gameCmd = &cobra.Command{
	Use:   "game",
	Short: "Get game information by name",
	Long: `Get several informations about a game by using it name for searching in IGDB database.`,
	Run: GetGame,
}

// Return the command.
func NewGameCmd() (*cobra.Command) {
    return gameCmd
}

func init() {
	rootCmd.AddCommand(gameCmd)
}

func GetGame(cmd *cobra.Command, args []string) {
	client := config.InitClient()
	game, err := client.Games.Search(
		strings.Join(args, " "),
		igdb.SetFields("*"),
		igdb.SetLimit(1),
	)
	
	if err != nil {
		log.Fatal(err)
	}

	title := utils.SetColor()
	fmt.Printf("%s %s\n", title("Name:"), game[0].Name)
	fmt.Printf("%s %s\n", title("Description:"), game[0].Summary)
	fmt.Printf("%s %s\n", title("URL:"), game[0].URL)
	fmt.Printf("%s %f\n", title("Rating:"), game[0].Rating)
}
