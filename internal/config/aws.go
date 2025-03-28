package config

import (
	"log/slog"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func CreateDynamoDBConnection(awsConfig AWSConfig) (*dynamodb.DynamoDB, *error) {
	config := aws.Config{
		Region:   aws.String(awsConfig.Region),
		Endpoint: aws.String(awsConfig.Endpoint),
	}
	sess := session.Must(session.NewSession(&config))
	dynamoDB := dynamodb.New(sess)
	_, err := dynamoDB.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		slog.Error("Error creating session on AWS", slog.String("error", err.Error()))
		return nil, &ErrCreatingAWSSession
	}

	return dynamoDB, nil
}
