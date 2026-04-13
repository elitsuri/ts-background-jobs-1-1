package worker

import "database/sql"

type Scheduler struct {
	cleanup *CleanupWorker
	email   *EmailWorker
}

func NewScheduler(db *sql.DB) *Scheduler {
	return &Scheduler{
		cleanup: NewCleanupWorker(db),
		email:   NewEmailWorker(db),
	}
}

func (s *Scheduler) Start() {
	s.cleanup.Start()
	s.email.Start()
}

func (s *Scheduler) Stop() {
	s.cleanup.Stop()
	s.email.Stop()
}
