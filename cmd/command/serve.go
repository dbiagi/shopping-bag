package command

import (
	"github.com/spf13/cobra"

	internalhttp "github.com/dbiagi/shopping-bag/internal/http"
)

func NewServeCommand() *cobra.Command {
	return &cobra.Command{
		Use: "serve",
		Run: runServe,
	}
}

func runServe(cmd *cobra.Command, args []string) {
	server := internalhttp.NewServer()
	server.Start()
}
