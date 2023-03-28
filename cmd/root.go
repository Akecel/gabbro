package cmd

import (
	"github.com/akecel/gabbro/config"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "gabbro",
		Short: "Simple CLI to interact with IGDB API",
		Long: `Gabbro is a simple CLI written in Go allowing a simplified interaction with the IGDB API. It is possible to search for information about specific video game and many other things.`,
	}
)

// Return the command.
func NewRootCmd() (*cobra.Command) {
    return rootCmd
}

// Execute executes the root command.
func Execute() (error) {
	return rootCmd.Execute()
}

// Init the command
func init() {
	cobra.OnInitialize(config.InitConfig)
}