package service

import "github.com/example/ts-background-jobs-1/internal/repository"

type AnalyticsService struct{ repo *repository.AnalyticsRepository }

func NewAnalyticsService(repo *repository.AnalyticsRepository) *AnalyticsService {
	return &AnalyticsService{repo: repo}
}

func (s *AnalyticsService) GetOverview() (map[string]int64, error) {
	return s.repo.Overview()
}

func (s *AnalyticsService) GetTimeseries(days int) ([]map[string]interface{}, error) {
	if days <= 0 || days > 365 { days = 30 }
	return s.repo.Timeseries(days)
}
