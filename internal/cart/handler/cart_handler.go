package handler

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/dbiagi/shopping-bag/internal/cart/domain"
	"github.com/dbiagi/shopping-bag/internal/cart/repository"
	"github.com/dbiagi/shopping-bag/pkg/httputil"
)

type CartHandler struct {
	CartRepository repository.CartRepositoryInterface
}

func NewCartHandler(cr repository.CartRepositoryInterface) CartHandler {
	return CartHandler{
		CartRepository: cr,
	}
}

func (c *CartHandler) Cart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cartID, err := uuid.Parse(vars["cartId"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		httputil.NewJSONResponse().Response(w, r)
		return
	}

	cart, err := c.CartRepository.CartByID(cartID)

	if err != nil && err == repository.ErrCartNotFound {
		w.WriteHeader(http.StatusNotFound)
		httputil.NewJSONResponse().Response(w, r)
		return
	}

	if err != nil {
		httputil.NewJSONResponse(httputil.WithStatusCode(http.StatusInternalServerError)).
			Response(w, r)
		return
	}

	httputil.NewJSONResponse(httputil.WithBody(cart)).
		Response(w, r)
}

func (c *CartHandler) CreateCart(w http.ResponseWriter, r *http.Request) {
	var request domain.CreateCartRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		httputil.NewJSONResponse(httputil.WithStatusCode(http.StatusBadRequest)).
			Response(w, r)
	}

}
