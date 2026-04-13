package service

import ("github.com/example/ts-background-jobs-1/internal/model"; "github.com/example/ts-background-jobs-1/internal/repository")

type NotificationService struct{ repo *repository.NotificationRepository }

func NewNotificationService(repo *repository.NotificationRepository) *NotificationService {
	return &NotificationService{repo: repo}
}

func (s *NotificationService) Create(userID int64, title, body, nType string) error {
	n := &model.Notification{UserID: userID, Title: title, Body: body, Type: nType}
	return s.repo.Create(n)
}

func (s *NotificationService) GetForUser(userID int64, limit, offset int) ([]model.Notification, error) {
	return s.repo.FindByUserID(userID, limit, offset)
}

func (s *NotificationService) MarkRead(id, userID int64) error {
	return s.repo.MarkRead(id, userID)
}
