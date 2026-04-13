package repository

import ("database/sql"; "time")
"github.com/example/ts-background-jobs-1/internal/model"

type AuditRepository struct{ db *sql.DB }

func (r *AuditRepository) Create(a *model.AuditLog) error {
	a.CreatedAt = time.Now()
	return r.db.QueryRow(`INSERT INTO audit_logs(user_id,action,resource_type,resource_id,ip_address,user_agent,created_at) VALUES($1,$2,$3,$4,$5,$6,$7) RETURNING id`
		, a.UserID, a.Action, a.ResourceType, a.ResourceID, a.IPAddress, a.UserAgent, a.CreatedAt).Scan(&a.ID)
}

func (r *AuditRepository) Cleanup(olderThanDays int) (int64, error) {
	res, err := r.db.Exec(`DELETE FROM audit_logs WHERE created_at < NOW() - ($1 || ' days')::INTERVAL`, olderThanDays)
	if err != nil { return 0, err }
	return res.RowsAffected()
}
