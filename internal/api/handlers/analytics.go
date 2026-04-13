package handlers

import ("encoding/json"; "net/http")

type AnalyticsHandler struct { repo AnalyticsRepo }
type AnalyticsRepo interface {
	Overview() (map[string]int64, error)
	Timeseries(days int) ([]map[string]interface{}, error)
}

func (h *AnalyticsHandler) Overview(w http.ResponseWriter, r *http.Request) {
	stats, err := h.repo.Overview()
	if err != nil { http.Error(w, err.Error(), 500); return }
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"data": stats})
}

func (h *AnalyticsHandler) Timeseries(w http.ResponseWriter, r *http.Request) {
	data, err := h.repo.Timeseries(30)
	if err != nil { http.Error(w, err.Error(), 500); return }
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"data": data})
}
