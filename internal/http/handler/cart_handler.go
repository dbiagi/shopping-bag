package handler

import (
	"net/http"

	"github.com/dbiagi/shopping-bag/internal/util"
)

type CartHandler struct {
}

func NewCartHandler() CartHandler {
	return CartHandler{}
}

func (c *CartHandler) Cart(w http.ResponseWriter, r *http.Request) {
	cart := map[string]string{"id": "1"}
	util.JsonResponse(w, r, cart)
}
