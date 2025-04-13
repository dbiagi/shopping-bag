package domain

import "github.com/google/uuid"

type Product struct {
	ID               uuid.UUID `json:"id"`
	SKU              string    `json:"sku"`
	Title            string    `json:"title"`
	ShortDescription string    `json:"shortDescription"`
	LongDescription  string    `json:"longDescription"`
	Category         string    `json:"category"`
}
