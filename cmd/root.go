package cmd

import (
	"github.com/spf13/cobra"
)

var (
	RootCmd = &cobra.Command{
		Use:   "gabbro",
		Short: "Simple CLI to interact with IGDB API",
		Long: `
	 _____       _     _               
	|  __ \     | |   | |              
	| |  \/ __ _| |__ | |__  _ __ ___  
	| | __ / _' | '_ \| '_ \| '__/ _ \ 
	| |_\ \ (_| | |_) | |_) | | | (_) |
	\_____/\__,_|_.__/|_.__/|_|  \___/`,
	}
)

var WithImage bool

func Execute() error {
	return RootCmd.Execute()
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&WithImage, "image", "i", false, "print covers, logos and mugshots in terminal")
}
