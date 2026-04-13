package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/example/ts-background-jobs-1/internal/api/middleware"
	"github.com/example/ts-background-jobs-1/internal/domain"
	"github.com/example/ts-background-jobs-1/internal/service"
	"github.com/example/ts-background-jobs-1/pkg/response"
)

type ItemHandler struct{ svc *service.ItemService }

func NewItemHandler(svc *service.ItemService) *ItemHandler { return &ItemHandler{svc: svc} }

func (h *ItemHandler) List(w http.ResponseWriter, r *http.Request) {
	userID, _ := middleware.GetUserID(r)
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page < 1 { page = 1 }
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if limit < 1 || limit > 100 { limit = 20 }
	search := r.URL.Query().Get("q")
	items, total, err := h.svc.List(r.Context(), userID, page, limit, search)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to list items")
		return
	}
	response.JSON(w, http.StatusOK, map[string]any{
		"data": items, "total": total, "page": page, "limit": limit,
	})
}

func (h *ItemHandler) Create(w http.ResponseWriter, r *http.Request) {
	userID, _ := middleware.GetUserID(r)
	var req domain.ItemCreate
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}
	item, err := h.svc.Create(r.Context(), userID, req)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to create item")
		return
	}
	response.JSON(w, http.StatusCreated, item)
}

func (h *ItemHandler) Get(w http.ResponseWriter, r *http.Request) {
	userID, _ := middleware.GetUserID(r)
	id, err := extractID(r.URL.Path)
	if err != nil { response.Error(w, http.StatusBadRequest, "invalid id"); return }
	item, err := h.svc.Get(r.Context(), id, userID)
	if err != nil { response.Error(w, http.StatusNotFound, "item not found"); return }
	response.JSON(w, http.StatusOK, item)
}

func (h *ItemHandler) Update(w http.ResponseWriter, r *http.Request) {
	userID, _ := middleware.GetUserID(r)
	id, err := extractID(r.URL.Path)
	if err != nil { response.Error(w, http.StatusBadRequest, "invalid id"); return }
	var req domain.ItemUpdate
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}
	item, err := h.svc.Update(r.Context(), id, userID, req)
	if err != nil { response.Error(w, http.StatusNotFound, "item not found"); return }
	response.JSON(w, http.StatusOK, item)
}

func (h *ItemHandler) Delete(w http.ResponseWriter, r *http.Request) {
	userID, _ := middleware.GetUserID(r)
	id, err := extractID(r.URL.Path)
	if err != nil { response.Error(w, http.StatusBadRequest, "invalid id"); return }
	if err := h.svc.Delete(r.Context(), id, userID); err != nil {
		response.Error(w, http.StatusNotFound, "item not found")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func extractID(path string) (int64, error) {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	return strconv.ParseInt(parts[len(parts)-1], 10, 64)
}
