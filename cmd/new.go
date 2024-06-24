package cmd

import (
	"github.com/cnnrznn/authenticator/model"
	"github.com/cnnrznn/authenticator/store"
	"github.com/spf13/cobra"
)

var (
	name string
	key  string
	// TODO consider providing parameters:
	// - digits
	// - algorithm
	// - issuer
	// - period

	newCmd = &cobra.Command{
		Use:   "new",
		Short: "Add a new app",

		RunE: func(cmd *cobra.Command, args []string) error {
			tokens, _ := store.Load()

			tokens = append(tokens, model.Token{
				Name:   name,
				Secret: key,
			})

			if err := store.Save(tokens); err != nil {
				return err
			}

			return nil
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
