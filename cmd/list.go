package cmd

import (
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists apps and current tokens",

	Run: func(cmd *cobra.Command, args []string) {
		// Read apps from storage
		// Compute current keys
		// Print
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
