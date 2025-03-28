package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:        "spcat",
	Args:       cobra.ExactArgs(1),
	ArgAliases: []string{"file"},
	Short:      "Special cat command",
	Long:       `Special cat command that can show images and pdf and every other file type.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
