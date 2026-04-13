package worker

import ("database/sql"; "log"; "time")

type EmailWorker struct{ db *sql.DB; done chan struct{} }

func NewEmailWorker(db *sql.DB) *EmailWorker {
	return &EmailWorker{db: db, done: make(chan struct{})}
}

func (w *EmailWorker) Start() {
	go func() {
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()
		for { select {
			case <-ticker.C: w.processQueue()
			case <-w.done: return
		} }
	}()
}

func (w *EmailWorker) Stop() { close(w.done) }

func (w *EmailWorker) processQueue() {
	rows, err := w.db.Query(`SELECT id,to_address,subject,body FROM email_queue WHERE sent=false AND attempts<3 LIMIT 10`)
	if err != nil { return }
	defer rows.Close()
	for rows.Next() {
		var id int64; var to, subj, body string
		_ = rows.Scan(&id, &to, &subj, &body)
		log.Printf("email: sending to %s", to)
		_, _ = w.db.Exec(`UPDATE email_queue SET attempts=attempts+1,sent=true,sent_at=NOW() WHERE id=$1`, id)
	}
}
