package cmd

import (
	"fmt"
	"os"
	"spcat/pkg/content"
	"spcat/pkg/renderer"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:        "spcat",
	Args:       cobra.ExactArgs(1),
	ArgAliases: []string{"file"},
	Short:      "Special cat command",
	Long:       `Special cat command that can show images and pdf and every other file type.`,
	Run: func(cmd *cobra.Command, args []string) {
		adress := args[0]
		content := content.GetContent(adress)
		img := renderer.GenerateKittyPic(content, renderer.KittyOptions)
		os.Stdout.Write(img)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
