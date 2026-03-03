package main

import (
    "context"
    "log"

    "github.com/dbiagi/shopping-bag/internal/config"
    "github.com/dbiagi/shopping-bag/pkg/contextutil"
    "github.com/dbiagi/shopping-bag/pkg/logging"
    "github.com/kelseyhightower/envconfig"
    "github.com/spf13/cobra"

    "github.com/dbiagi/shopping-bag/cmd/command"
)

var (
    buildVersion = "unknown"
    goVersion    = "unknown"
    buildTime    = "unknown"
)

func main() {
    cfg := &config.Configuration{}
    err := envconfig.Process("", cfg)
    if err != nil {
        log.Fatalf("error loading configuration, error: %+v", err)
    }

    ctx := contextutil.WithServiceConfig(context.Background(), cfg)
    logLevel := logging.LogLevelFromString(cfg.LogLevel)
    logging.NewLogger(logLevel, map[string]string{
        "service":   "shopping-bag",
        "version":   buildVersion,
        "goVersion": goVersion,
        "buildTime": buildTime,
    })

    rootCmd := cobra.Command{
        Use:   cfg.AppConfig.Name,
        Short: cfg.AppConfig.Description,
    }
    rootCmd.AddCommand(command.NewServeCommand())
    rootCmd.SetContext(ctx)

    if cmdErr := rootCmd.Execute(); cmdErr != nil {
        log.Fatalf("error executing a command, error: %+v", cmdErr)
    }
}
