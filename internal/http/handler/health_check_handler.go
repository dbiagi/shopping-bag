package handler

import (
	"net/http"

	"github.com/dbiagi/shopping-bag/internal/util"
)

type (
	HealthStatus string

	HealthCheck struct {
		Status HealthStatus `json:"status"`
	}

	HealthCheckHandler struct {
	}
)

const (
	Ok   HealthStatus = "ok"
	Down HealthStatus = "down"
)

func NewHealthCheckHandler() HealthCheckHandler {
	return HealthCheckHandler{}
}

func (h *HealthCheckHandler) Health(w http.ResponseWriter, r *http.Request) {
	hc := HealthCheck{
		Status: Ok,
	}

	util.JsonResponse(w, r, hc)
}
