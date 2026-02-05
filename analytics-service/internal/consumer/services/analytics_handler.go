package services

import (
	"analytics-service/internal/repository"
	"context"
	"encoding/json"
	"events"
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
	ctx := context.Background()
	a.repo.LogEvent(context.Background(), topic, string(data))

	switch topic {
	case "OpenAccountEvent":
		ev := events.OpenAccountEvent{}
		if err := json.Unmarshal(data, &ev); err != nil {
			return
		}
		a.repo.SaveDeposit(ctx, ev.ID, ev.OpeningBalance, "opening")

	case "DepositFundEvent":
		ev := events.DepositFundEvent{}
		if err := json.Unmarshal(data, &ev); err != nil {
			return
		}
		a.repo.SaveDeposit(ctx, ev.ID, ev.Amount, "deposit")

	case "WithdrawFundEvent":
		ev := events.WithdrawFundEvent{}
		if err := json.Unmarshal(data, &ev); err != nil {
			return
		}
		finalAmount := -ev.Amount
		a.repo.SaveDeposit(ctx, ev.ID, finalAmount, "withdraw")
	}

}
