package handler

import (
	"net/http"

	"github.com/dbiagi/shopping-bag/internal/repository"
	"github.com/dbiagi/shopping-bag/internal/util"
)

type CartHandler struct {
	repository.CartRepository
}

func NewCartHandler(cr repository.CartRepository) CartHandler {
	return CartHandler{
		CartRepository: cr,
	}
}

func (c *CartHandler) Cart(w http.ResponseWriter, r *http.Request) {
	cart := map[string]string{"id": "1"}
	util.JsonResponse(w, r, cart)
}
