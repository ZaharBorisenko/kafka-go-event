package services

import (
	"analytics-service/internal/repository"
	"context"
)

type EventHandler interface {
	Handle(topic string, data []byte)
}

type analyticsEventHandler struct {
	repo repository.AnalyticsRepository
}

func NewAnalyticsEventHandler(repo repository.AnalyticsRepository) EventHandler {
	return &analyticsEventHandler{repo: repo}
}

func (a analyticsEventHandler) Handle(topic string, data []byte) {
	a.repo.LogEvent(context.Background(), topic, string(data))
}
