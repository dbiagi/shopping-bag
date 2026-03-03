package config

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Configuration struct {
	WebConfig
	AppConfig
	AWSConfig
	LogLevel string `envconfig:"APP_LOG_LEVEL"`
}

type AppConfig struct {
	Name        string `envconfig:"APP_NAME"`
	Description string `envconfig:"APP_DESCRIPTION"`
	Version     string `envconfig:"APP_VERSION"`
	Environment string `envconfig:"APP_ENVIRONMENT"`
}

type WebConfig struct {
	Port                     int           `envconfig:"HTTP_PORT"`
	IdleTimeout              time.Duration `envconfig:"HTTP_IDLE_TIMEOUT"`
	ReadTimeout              time.Duration `envconfig:"HTTP_READ_TIMEOUT"`
	WriteTimeout             time.Duration `envconfig:"HTTP_WRITE_TIMEOUT"`
	ShutdownTimeout          time.Duration `envconfig:"HTTP_SHUTDOWN_TIMEOUT"`
	GracefulShutdownDisabled bool          `envconfig:"HTTP_GRACEFUL_SHUTDOWN_DISABLED"`
}

type AWSConfig struct {
	Region string `envconfig:"AWS_REGION"`
	AWSCredentialsConfig
	DynamoDBConfig
}

type AWSCredentialsConfig struct {
	CredentialType  string `envconfig:"AWS_CREDENTIAL_TYPE"`
	AccessKeyID     string `envconfig:"AWS_ACCESS_KEY_ID"`
	SecretAccessKey string `envconfig:"AWS_SECRET_ACCESS_KEY"`
}

type DynamoDBConfig struct {
	Endpoint string `envconfig:"AWS_DYNAMODB_ENDPOINT"`
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
			Name:        "shopping-bag",
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
