package repository

import "database/sql"

type AnalyticsRepository struct{ db *sql.DB }

func (r *AnalyticsRepository) Overview() (map[string]int64, error) {
	m := map[string]int64{}
	_ = r.db.QueryRow(`SELECT COUNT(*) FROM users`).Scan(&m["users"])
	_ = r.db.QueryRow(`SELECT COUNT(*) FROM items`).Scan(&m["items"])
	_ = r.db.QueryRow(`SELECT COUNT(*) FROM notifications`).Scan(&m["notifications"])
	_ = r.db.QueryRow(`SELECT COUNT(*) FROM items WHERE created_at > NOW()-INTERVAL'7 days'`).Scan(&m["items_7d"])
	return m, nil
}

func (r *AnalyticsRepository) Timeseries(days int) ([]map[string]interface{}, error) {
	rows, err := r.db.Query(`SELECT DATE(created_at) d, COUNT(*) c FROM items WHERE created_at > NOW()-($1 || ' days')::INTERVAL GROUP BY d ORDER BY d`, days)
	if err != nil { return nil, err }
	defer rows.Close()
	var result []map[string]interface{}
	for rows.Next() {
		var d string; var c int64
		_ = rows.Scan(&d, &c)
		result = append(result, map[string]interface{}{"date": d, "count": c})
	}
	return result, nil
}
