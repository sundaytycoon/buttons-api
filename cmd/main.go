package main

import (
	"github.com/spf13/cobra"
	// _ "go.uber.org/automaxprocs"

	cmdserver "github.com/sundaytycoon/buttons-api/cmd/server"
)

var rootCmd = &cobra.Command{
	Use:     "profileme",
	Aliases: []string{"pm"},
	Short:   "profile.me server project",
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
	)

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
