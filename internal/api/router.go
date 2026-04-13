package api

import ("github.com/example/ts-background-jobs-1/internal/api/handlers"; "github.com/example/ts-background-jobs-1/internal/api/middleware"; "net/http")

func NewRouter(h *handlers.Handlers) http.Handler {
	mux := http.NewServeMux()
	// Static files
	mux.Handle("/", http.FileServer(http.Dir("web/static")))
	// Health
	mux.HandleFunc("GET /api/v1/health", h.Health.Status)
	// Auth
	mux.HandleFunc("POST /api/v1/auth/register", h.Auth.Register)
	mux.HandleFunc("POST /api/v1/auth/login", h.Auth.Login)
	mux.HandleFunc("POST /api/v1/auth/refresh", h.Auth.Refresh)
	// Items (auth required)
	mux.Handle("GET /api/v1/items", middleware.Auth(http.HandlerFunc(h.Items.List)))
	mux.Handle("POST /api/v1/items", middleware.Auth(http.HandlerFunc(h.Items.Create)))
	mux.Handle("GET /api/v1/items/{id}", middleware.Auth(http.HandlerFunc(h.Items.Get)))
	mux.Handle("PUT /api/v1/items/{id}", middleware.Auth(http.HandlerFunc(h.Items.Update)))
	mux.Handle("DELETE /api/v1/items/{id}", middleware.Auth(http.HandlerFunc(h.Items.Delete)))
	// Analytics
	mux.Handle("GET /api/v1/analytics/overview", middleware.Auth(http.HandlerFunc(h.Analytics.Overview)))
	mux.Handle("GET /api/v1/analytics/timeseries", middleware.Auth(http.HandlerFunc(h.Analytics.Timeseries)))
	// Notifications
	mux.Handle("GET /api/v1/notifications", middleware.Auth(http.HandlerFunc(h.Notifications.List)))
	mux.Handle("POST /api/v1/notifications", middleware.Auth(http.HandlerFunc(h.Notifications.Create)))
	mux.Handle("PUT /api/v1/notifications/{id}/read", middleware.Auth(http.HandlerFunc(h.Notifications.MarkRead)))
	// Admin
	mux.Handle("GET /api/v1/admin/users", middleware.Admin(http.HandlerFunc(h.Admin.ListUsers)))
	mux.Handle("PUT /api/v1/admin/users/{id}/role", middleware.Admin(http.HandlerFunc(h.Admin.UpdateRole)))
	// Search
	mux.Handle("GET /api/v1/search", middleware.Auth(http.HandlerFunc(h.Search.Query)))
	// Upload
	mux.Handle("POST /api/v1/upload/file", middleware.Auth(http.HandlerFunc(h.Upload.File)))
	// WebSocket
	mux.Handle("GET /ws", middleware.Auth(http.HandlerFunc(h.WS.Connect)))
	return middleware.Chain(mux, middleware.CORS, middleware.RequestID, middleware.Logging, middleware.RateLimit)
}
