package domain

import (
	"github.com/google/uuid"

	productdomain "github.com/dbiagi/shopping-bag/internal/product/domain"
)

type (
	Cart struct {
		ID             uuid.UUID  `json:"id"`
		OrganizationID uuid.UUID  `json:"organizationId"`
		Items          []CartItem `json:"bundles"`
	}

	// Bundle struct {
	// 	ID                 uuid.UUID  `json:"id"`
	// 	SellerID           uuid.UUID  `json:"sellerId"`
	// 	CartItems          []CartItem `json:"cartItems"`
	// 	CommitmentLocation string     `json:"commitmentLocation"`
	// }

	CartItem struct {
		Product  productdomain.Product `json:"product"`
		Quantity int                   `json:"quantity"`
		SellerID uuid.UUID             `json:"sellerId"`
	}

	CreateCartRequest struct {
		OrganizationID uuid.UUID  `json:"organizationId"`
		Items          []CartItem `json:"items"`
	}
)
