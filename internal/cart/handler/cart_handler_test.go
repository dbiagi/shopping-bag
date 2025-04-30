package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/dbiagi/shopping-bag/internal/cart/repository"
	"github.com/dbiagi/shopping-bag/internal/cart/repository/mocks"
)

func TestGetCart(t *testing.T) {
	type testCase struct {
		assertions func(w *httptest.ResponseRecorder)
		setup      func(h CartHandler) *mux.Router
		setupMocks func(cr *mocks.MockCartRepositoryInterface, cartId uuid.UUID)
	}

	tc := []testCase{
		{
			assertions: func(w *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, w.Code)
			},
			setup: func(h CartHandler) *mux.Router {
				router := mux.NewRouter()
				router.HandleFunc("/carts/{cartId}", h.Cart).Methods("GET")

				return router
			},
			setupMocks: func(cr *mocks.MockCartRepositoryInterface, cartID uuid.UUID) {
				cart := &repository.Cart{
					ID: cartID.String(),
				}

				cr.EXPECT().CartById(cartID).Return(cart, nil)
			},
		},
	}

	for _, tc := range tc {
		ctrl := gomock.NewController(t)
		cartRepository := mocks.NewMockCartRepositoryInterface(ctrl)
		handler := NewCartHandler(cartRepository)
		cartID := uuid.Must(uuid.NewRandom())

		tc.setupMocks(cartRepository, cartID)

		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", fmt.Sprintf("/carts/%s", cartID.String()), nil)

		router := tc.setup(handler)
		router.ServeHTTP(w, r)

		handler.Cart(w, r)

		tc.assertions(w)
	}
}
