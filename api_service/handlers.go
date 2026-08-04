package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

// healthHandler is a liveness probe: it should only fail if the process
// itself is unable to serve requests.
func healthHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

// readyHandler is a readiness probe. Extend it to check downstream
// dependencies (database, cache, upstream services) before reporting ready,
// so orchestrators stop routing traffic here during startup or an outage.
func readyHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ready"})
}

type exampleResponse struct {
	Message   string    `json:"message"`
	RequestID string    `json:"request_id"`
	Time      time.Time `json:"time"`
}

// exampleHandler demonstrates a protected route. Replace with real
// business-logic handlers as this service grows.
func exampleHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := r.Context().Value(requestIDKey).(string)
	writeJSON(w, http.StatusOK, exampleResponse{
		Message:   "authenticated request succeeded",
		RequestID: id,
		Time:      time.Now().UTC(),
	})
}
