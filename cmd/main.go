package main

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/dbiagi/shopping-bag/cmd/command"
)

func main() {
	rootCmd := cobra.Command{
		Use:   "shopping-bag",
		Short: "Shopping bag app CLI.",
	}

	rootCmd.AddCommand(command.NewServeCommand())
	err := rootCmd.Execute()

	if err != nil {
		log.Fatal("Error executing a command")
	}
}
