package worker

import ("database/sql"; "log"; "time")

type CleanupWorker struct{ db *sql.DB; done chan struct{} }

func NewCleanupWorker(db *sql.DB) *CleanupWorker {
	return &CleanupWorker{db: db, done: make(chan struct{})}
}

func (w *CleanupWorker) Start() {
	go func() {
		ticker := time.NewTicker(24 * time.Hour)
		defer ticker.Stop()
		for { select {
			case <-ticker.C: w.run()
			case <-w.done: return
		} }
	}()
}

func (w *CleanupWorker) Stop() { close(w.done) }

func (w *CleanupWorker) run() {
	res, _ := w.db.Exec(`DELETE FROM audit_logs WHERE created_at < NOW() - INTERVAL '90 days'`)
	n, _ := res.RowsAffected()
	log.Printf("cleanup: removed %d old audit logs", n)
}
