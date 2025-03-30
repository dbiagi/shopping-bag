package repository

import "errors"

var (
	ErrFetchingCart = errors.New("error fetching the cart from db")
	ErrCartNotFound = errors.New("cart now found")
)
