package command

import (
	"github.com/dbiagi/shopping-bag/internal/config"
	internalhttp "github.com/dbiagi/shopping-bag/internal/http"

	"github.com/spf13/cobra"
)

func NewServeCommand() *cobra.Command {
	return &cobra.Command{
		Use: "serve",
		Run: runServe,
	}
}

func runServe(cmd *cobra.Command, args []string) {
	env, _ := cmd.Flags().GetString("env")
	if env == "" {
		env = config.DevelopmentEnv
	}
	c := config.LoadConfig(env)
	server := internalhttp.NewServer(c)
	server.Start()
}
