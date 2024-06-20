package cmd

import (
	"github.com/spf13/cobra"
)

var (
	name string
	key  string

	newCmd = &cobra.Command{
		Use:   "new",
		Short: "Add a new app",

		Run: func(cmd *cobra.Command, args []string) {
			// Add name, key to encrypted storage
		},
	}
)

func init() {
	rootCmd.AddCommand(newCmd)

	newCmd.Flags().StringVar(&name, "name", "", "Name of the new app")
	newCmd.Flags().StringVar(&key, "key", "", "Setup key provided by app")

	newCmd.MarkFlagRequired("name")
	newCmd.MarkFlagRequired("key")
}
