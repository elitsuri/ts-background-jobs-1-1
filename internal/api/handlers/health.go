package handlers

import (
	"net/http"
	"time"

	"github.com/example/ts-background-jobs-1/pkg/response"
)

var startTime = time.Now()

// HealthHandler handles /health endpoints.
type HealthHandler struct{}

func NewHealthHandler() *HealthHandler { return &HealthHandler{} }

func (h *HealthHandler) Health(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, map[string]any{
		"status":  "healthy",
		"service": "ts-background-jobs-1",
		"version": "1.0.0",
		"uptime":  time.Since(startTime).String(),
	})
}

func (h *HealthHandler) Ready(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, map[string]any{"status": "ready"})
}
