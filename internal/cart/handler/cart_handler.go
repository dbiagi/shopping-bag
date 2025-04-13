package handler

import (
	"encoding/json"
	"net/http"

	"github.com/dbiagi/shopping-bag/internal/cart/domain"
	"github.com/dbiagi/shopping-bag/internal/cart/repository"
	"github.com/dbiagi/shopping-bag/pkg/httputil"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
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
	vars := mux.Vars(r)
	cartId, err := uuid.Parse(vars["cartId"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		httputil.NewJsonResponse().Response(w, r)
		return
	}

	cart, err := c.CartRepository.CartById(cartId)

	if err != nil && err == repository.ErrCartNotFound {
		w.WriteHeader(http.StatusNotFound)
		httputil.NewJsonResponse().Response(w, r)
		return
	}

	if err != nil {
		httputil.NewJsonResponse(httputil.WithStatusCode(http.StatusInternalServerError)).
			Response(w, r)
		return
	}

	httputil.NewJsonResponse(httputil.WithBody(cart)).
		Response(w, r)
}

func (c *CartHandler) CreateCart(w http.ResponseWriter, r *http.Request) {
	var request domain.CreateCartRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		httputil.NewJsonResponse(httputil.WithStatusCode(http.StatusBadRequest)).
			Response(w, r)
	}

}
