package handlers

import ("encoding/json"; "net/http")

type SearchHandler struct { svc SearchSvc }
type SearchSvc interface { Search(q, sType string, limit int) (map[string]interface{}, error) }

func (h *SearchHandler) Query(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	sType := r.URL.Query().Get("type")
	if sType == "" { sType = "all" }
	result, err := h.svc.Search(q, sType, 20)
	if err != nil { http.Error(w, err.Error(), 500); return }
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"data": result})
}
