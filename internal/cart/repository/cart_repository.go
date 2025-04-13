package repository

import (
	"fmt"
	"log/slog"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
)

//go:generate mockgen -source=cart_repository.go -destination=./mocks/cart_repository_mock.go -package=mocks

const (
	TableName    = "cart"
	PartitionKey = "id"
)

type (
	Cart struct {
		Id string `json:"id"`
	}

	CartRepositoryInterface interface {
		CartById(id uuid.UUID) Cart
	}

	CartRepository struct {
		*dynamodb.DynamoDB
	}
)

func NewCartRepository(db *dynamodb.DynamoDB) CartRepository {
	return CartRepository{
		DynamoDB: db,
	}
}

func (cr *CartRepository) CartById(id uuid.UUID) (*Cart, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(TableName),
		Key: map[string]*dynamodb.AttributeValue{
			PartitionKey: {
				S: aws.String(id.String()),
			},
		},
	}

	result, err := cr.DynamoDB.GetItem(input)
	if err != nil {
		slog.Error(fmt.Sprintf("Error fetching cart %s from dynamodb.", id.String()), slog.String("error", err.Error()))
		return nil, ErrFetchingCart
	}

	if result.Item == nil {
		return nil, ErrCartNotFound
	}

	var cart *Cart
	if err := dynamodbattribute.UnmarshalMap(result.Item, cart); err != nil {
		slog.Error(fmt.Sprintf("Error unmarshaling cart. Error: %s", err.Error()))
		return nil, ErrFetchingCart
	}

	return cart, nil
}
