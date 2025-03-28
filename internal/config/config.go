package config

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

const (
	AppName        = "shopping-bag"
	DevelopmentEnv = "dev"
	ProductionEnv  = "prod"
)

type Configuration struct {
	WebConfig
	AppConfig
	AWSConfig
}

type AppConfig struct {
	Name        string
	Version     string
	Environment string
}

type WebConfig struct {
	Port                     int
	IdleTimeout              time.Duration
	ReadTimeout              time.Duration
	WriteTimeout             time.Duration
	ShutdownTimeout          time.Duration
	GracefulShutdownDisabled bool
}

type AWSConfig struct {
	Region string
	AWSCredentialsConfig
	DynamoDBConfig
}

type AWSCredentialsConfig struct {
	CredentialType  string
	AccessKeyID     string
	SecretAccessKey string
}

type DynamoDBConfig struct {
	Endpoint string
}

func LoadConfig(env string) Configuration {
	configs := loadFromFile()
	port, _ := strconv.Atoi(configs["PORT"])

	return Configuration{
		WebConfig: WebConfig{
			Port:            port,
			IdleTimeout:     time.Second * 10,
			ReadTimeout:     time.Second * 10,
			WriteTimeout:    time.Second * 10,
			ShutdownTimeout: time.Second * 20,
		},
		AppConfig: AppConfig{
			Name:        AppName,
			Version:     "1.0.0",
			Environment: env,
		},
		AWSConfig: AWSConfig{
			Region: configs["AWS_REGION"],
			DynamoDBConfig: DynamoDBConfig{
				Endpoint: configs["AWS_DYNAMODB_ENDPOINT"],
			},
		},
	}
}

func loadFromFile() map[string]string {
	path, _ := os.Getwd()
	configFilePath := fmt.Sprintf("%s/.env", path)

	configs, err := godotenv.Read(configFilePath)

	if err != nil {
		slog.Error("Error loading .env file.", slog.String("error", err.Error()))
		panic(err)
	}

	return configs
}
