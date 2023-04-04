package cmd

import (
	"github.com/spf13/cobra"
)

var (
	RootCmd = &cobra.Command{
		Use:   "gabbro",
		Short: "Simple CLI to interact with IGDB API",
		Long:  `Gabbro is a simple CLI written in Go allowing interaction with the IGDB API. It is possible to search for information about specific video game and many other things.`,
	}
)

var WithImage bool

func Execute() error {
	return RootCmd.Execute()
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&WithImage, "image", "i", false, "print covers, logos and mugshots in terminal")
}
