package main

import (
	"github.com/spf13/cobra"

	"github.com/dbiagi/shopping-bag/cmd/command"
)

func main() {
	rootCmd := cobra.Command{
		Use:   "shopping-bag",
		Short: "Shopping bag app CLI.",
	}

	rootCmd.AddCommand(command.NewServeCommand())
	rootCmd.Execute()
}
