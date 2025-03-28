// this command is for interactive mode feature for later
package cmd

import "github.com/spf13/cobra"

var interactiveCmd = &cobra.Command{
	Use:   "interactive",
	Short: "Start spcat in interactive mode",
	Long:  "Start spcat in interactive mode which means that you can switch files in the directory",
	Run: func(cmd *cobra.Command, args []string) {
		// todo
	},
}
