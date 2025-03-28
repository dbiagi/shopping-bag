package domain

import "github.com/google/uuid"

type (
	Cart struct {
		ID             uuid.UUID `json:"id"`
		OrganizationID uuid.UUID `json:"organizationId"`
		Bundles        []Bundle  `json:"bundles"`
	}

	Bundle struct {
		ID                 uuid.UUID  `json:"id"`
		SellerID           uuid.UUID  `json:"sellerId"`
		CartItems          []CartItem `json:"cartItems"`
		CommitmentLocation string     `json:"commitmentLocation"`
	}

	CartItem struct {
		Product  Product `json:"product"`
		Quantity int     `json:"quantity"`
	}
)
