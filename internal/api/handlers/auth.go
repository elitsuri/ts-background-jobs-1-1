package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/example/ts-background-jobs-1/internal/domain"
	"github.com/example/ts-background-jobs-1/internal/service"
	"github.com/example/ts-background-jobs-1/pkg/response"
)

type AuthHandler struct { svc *service.AuthService }

func NewAuthHandler(svc *service.AuthService) *AuthHandler { return &AuthHandler{svc: svc} }

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req domain.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}
	if req.Email == "" || req.Password == "" || req.FullName == "" {
		response.Error(w, http.StatusBadRequest, "email, password, and full_name are required")
		return
	}
	user, err := h.svc.Register(r.Context(), req)
	if err != nil {
		if err == service.ErrEmailTaken {
			response.Error(w, http.StatusConflict, "email already registered")
			return
		}
		response.Error(w, http.StatusInternalServerError, "registration failed")
		return
	}
	response.JSON(w, http.StatusCreated, user)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req domain.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}
	tokens, err := h.svc.Login(r.Context(), req)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, "invalid email or password")
		return
	}
	response.JSON(w, http.StatusOK, tokens)
}
