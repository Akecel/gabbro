package cmd

import (
	"fmt"
	"strings"
	"log"

	"github.com/akecel/gabbro/config"

	"github.com/Henry-Sarabia/igdb/v2"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// getGameByNameCmd represents the getGameByName command
var getGameByNameCmd = &cobra.Command{
	Use:   "game",
	Short: "Get game information by name",
	Long: `Get several informations about a game by using it name for searching in IGDB database.`,
	Run: func(cmd *cobra.Command, args []string) {
		getGame(args)
	},
}

// Return the command.
func NewGetGameByNameCmd() *cobra.Command {
    return getGameByNameCmd
}

func init() {
	rootCmd.AddCommand(getGameByNameCmd)
}

func getGame(args []string) {
	client := config.InitClient()
	game, err := client.Games.Search(
		strings.Join(args, " "),
		igdb.SetFields("*"),
		igdb.SetLimit(1),
	)
	
	if err != nil {
		log.Fatal(err)
	}

	title := color.New(color.FgCyan, color.Bold, color.Underline).SprintFunc()

	fmt.Printf("%s %s\n", title("Name:"), game[0].Name)
	fmt.Printf("%s %s\n", title("Description:"), game[0].Summary)
	fmt.Printf("%s %s\n", title("URL:"), game[0].URL)
	fmt.Printf("%s %f\n", title("Rating:"), game[0].Rating)
}
