package main

import (
	"github.com/dbiagi/shopping-bag/cmd/command"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := cobra.Command{
		Use:   "shopping-bag",
		Short: "Shopping bag app CLI.",
	}

	rootCmd.AddCommand(command.NewServeCommand())
	rootCmd.Execute()
}
