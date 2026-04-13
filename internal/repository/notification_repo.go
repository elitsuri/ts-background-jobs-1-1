package repository

import ("database/sql"; "time")
"github.com/example/ts-background-jobs-1/internal/model"

type NotificationRepository struct{ db *sql.DB }

func (r *NotificationRepository) Create(n *model.Notification) error {
	return r.db.QueryRow(`INSERT INTO notifications(user_id,title,body,type,created_at) VALUES($1,$2,$3,$4,$5) RETURNING id`
		, n.UserID, n.Title, n.Body, n.Type, time.Now()).Scan(&n.ID)
}

func (r *NotificationRepository) FindByUserID(userID int64, limit, offset int) ([]model.Notification, error) {
	rows, err := r.db.Query(`SELECT id,user_id,title,body,type,read,created_at FROM notifications WHERE user_id=$1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`, userID, limit, offset)
	if err != nil { return nil, err }
	defer rows.Close()
	var result []model.Notification
	for rows.Next() {
		var n model.Notification
		_ = rows.Scan(&n.ID, &n.UserID, &n.Title, &n.Body, &n.Type, &n.Read, &n.CreatedAt)
		result = append(result, n)
	}
	return result, nil
}

func (r *NotificationRepository) MarkRead(id, userID int64) error {
	_, err := r.db.Exec(`UPDATE notifications SET read=true WHERE id=$1 AND user_id=$2`, id, userID)
	return err
}
