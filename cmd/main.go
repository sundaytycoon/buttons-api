package main

import (
	"github.com/spf13/cobra"

	cmdent "github.com/sundaytycoon/buttons-api/cmd/ent"
	cmdserver "github.com/sundaytycoon/buttons-api/cmd/server"
)

var rootCmd = &cobra.Command{
	Use:     "buttons",
	Aliases: []string{"btn"},
	Short:   "buttons server project",
	RunE: func(c *cobra.Command, _ []string) error {
		return c.Help()
	},
	CompletionOptions: struct {
		DisableDefaultCmd   bool
		DisableNoDescFlag   bool
		DisableDescriptions bool
	}{
		DisableDefaultCmd:   true,
		DisableNoDescFlag:   true,
		DisableDescriptions: true,
	},
}

func main() {
	rootCmd.AddCommand(
		cmdserver.ServerCommand(),
		cmdent.MigrationCommand(),
	)

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
