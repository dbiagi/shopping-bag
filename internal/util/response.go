package util

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

func JsonResponse(w http.ResponseWriter, r *http.Request, body any) {
	w.Header().Set("Content-Type", "application/json")

	bytes, err := json.Marshal(body)

	if err != nil {
		slog.Error(
			fmt.Sprintf("Error marshalling response: %v", err.Error()),
			slog.String("error", err.Error()),
		)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.Write(bytes)
}
