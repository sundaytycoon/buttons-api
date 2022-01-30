package main

import (
	"github.com/spf13/cobra"

	cmdentgo "github.com/sundaytycoon/buttons-api/cmd/entgo"
	cmdserver "github.com/sundaytycoon/buttons-api/cmd/server"
)

var rootCmd = &cobra.Command{
	Use:     "buttons",
	Aliases: []string{"btn"},
	Short:   "buttons server project",
	RunE: func(c *cobra.Command, _ []string) error {
		return c.Help()
	},
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd:   true,
		DisableNoDescFlag:   true,
		DisableDescriptions: true,
	},
}

func main() {
	rootCmd.AddCommand(
		cmdserver.ServerCommand(),
		cmdentgo.MigrationCommand(),
	)

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
