package httputil

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

type (
	JSONResponse struct {
		statusCode *int
		body       *any
	}

	Option func(*JSONResponse)
)

func NewJSONResponse(opts ...Option) *JSONResponse {
	jr := &JSONResponse{}

	for _, opt := range opts {
		opt(jr)
	}

	return jr
}

func (jr *JSONResponse) Response(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if jr.statusCode != nil {
		w.WriteHeader(*jr.statusCode)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	if jr.body != nil {
		writeBoby(*jr.body, w)
	}
}

func WithStatusCode(statusCode int) Option {
	return func(jr *JSONResponse) {
		jr.statusCode = &statusCode
	}
}

func WithBody(body any) Option {
	return func(jr *JSONResponse) {
		jr.body = &body
	}
}

func writeBoby(body any, w http.ResponseWriter) {
	bytes, err := json.Marshal(body)
	if err != nil {
		slog.Error(
			fmt.Sprintf("Error marshalling response: %v", err.Error()),
			slog.String("error", err.Error()),
		)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	_, err = w.Write(bytes)
	if err != nil {
		slog.Error("Error writing response body", slog.String("error", err.Error()))
	}
}
