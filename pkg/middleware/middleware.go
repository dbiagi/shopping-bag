package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type (
	ContextKey string
)

const (
	TraceIDHeader                = "X-TRACE-ID"
	TraceIDContextKey ContextKey = "traceId"
)

func TraceIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		traceID := r.Header.Get(TraceIDHeader)
		if traceID == "" {
			traceID = uuid.New().String()
		}
		w.Header().Set(TraceIDHeader, traceID)

		ctx := r.Context()
		ctx = context.WithValue(ctx, TraceIDContextKey, traceID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
