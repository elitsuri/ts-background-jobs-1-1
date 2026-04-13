package handlers

import ("encoding/json"; "net/http"; "strconv")

type AdminHandler struct { repo AdminRepo }
type AdminRepo interface {
	ListUsers(limit, offset int) ([]map[string]interface{}, int64, error)
	UpdateRole(id int64, role string) error
	DeleteUser(id int64) error
}

func (h *AdminHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, total, err := h.repo.ListUsers(50, 0)
	if err != nil { http.Error(w, err.Error(), 500); return }
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"data":users,"total":total})
}

func (h *AdminHandler) UpdateRole(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(r.PathValue("id"), 10, 64)
	var req struct{ Role string }
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil { http.Error(w, err.Error(), 400); return }
	if err := h.repo.UpdateRole(id, req.Role); err != nil { http.Error(w, err.Error(), 500); return }
	w.WriteHeader(http.StatusNoContent)
}
