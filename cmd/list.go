package cmd

import (
	"fmt"
	"time"

	"github.com/cnnrznn/authenticator/store"
	"github.com/cnnrznn/authenticator/totp"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists apps and current tokens",

	RunE: func(cmd *cobra.Command, args []string) error {
		tokens, _ := store.Load()

		fmt.Printf("%16v %-16v\n", "Name", "Token")
		for _, t := range tokens {
			token, err := totp.Generate(t.Secret, time.Now())
			if err != nil {
				return err
			}

			fmt.Printf("%16v %-16v\n", t.Name, token)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
