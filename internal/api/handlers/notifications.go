package handlers

import ("encoding/json"; "net/http"; "strconv")

type NotificationsHandler struct { svc NotificationSvc }
type NotificationSvc interface {
	List(userID int64, limit, offset int) ([]map[string]interface{}, error)
	Create(userID int64, title, body, nType string) error
	MarkRead(id, userID int64) error
}

func (h *NotificationsHandler) List(w http.ResponseWriter, r *http.Request) {
	userID, _ := r.Context().Value(ctxKeyUserID).(int64)
	notifs, err := h.svc.List(userID, 50, 0)
	if err != nil { http.Error(w, err.Error(), 500); return }
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"data": notifs})
}

func (h *NotificationsHandler) Create(w http.ResponseWriter, r *http.Request) {
	userID, _ := r.Context().Value(ctxKeyUserID).(int64)
	var req struct{ Title, Body, Type string }
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil { http.Error(w, err.Error(), 400); return }
	if err := h.svc.Create(userID, req.Title, req.Body, req.Type); err != nil { http.Error(w, err.Error(), 500); return }
	w.WriteHeader(http.StatusCreated)
}

func (h *NotificationsHandler) MarkRead(w http.ResponseWriter, r *http.Request) {
	userID, _ := r.Context().Value(ctxKeyUserID).(int64)
	id, _ := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err := h.svc.MarkRead(id, userID); err != nil { http.Error(w, err.Error(), 500); return }
	w.WriteHeader(http.StatusNoContent)
}
