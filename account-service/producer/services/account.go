package services

import (
	"errors"
	"events"
	"github.com/google/uuid"
	"log"
	"producer/command"
)

type AccountService interface {
	OpenAccount(command command.OpenAccountCommand) (id string, err error)
	DepositFund(command command.DepositFundCommand) error
	WithDrawFund(command command.WithDrawFundCommand) error
	CloseAccount(command command.CloseAccountCommand) error
}

type accountService struct {
	eventProducer EventProducer
}

func NewAccountService(eventProducer EventProducer) AccountService {
	return accountService{eventProducer: eventProducer}
}

func (a accountService) OpenAccount(command command.OpenAccountCommand) (id string, err error) {
	if command.AccountHolder == "" || command.AccountType == 0 || command.OpeningBalance == 0 {
		return "", errors.New("bad request")
	}

	event := events.OpenAccountEvent{
		ID:             uuid.NewString(),
		AccountHolder:  command.AccountHolder,
		AccountType:    command.AccountType,
		OpeningBalance: command.OpeningBalance,
	}
	log.Printf("%#v", event)
	return event.ID, a.eventProducer.Produce(event)
}

func (a accountService) DepositFund(command command.DepositFundCommand) error {
	if command.ID == "" || command.Amount == 0 {
		return errors.New("bad request")
	}
	event := events.DepositFundEvent{
		ID:     command.ID,
		Amount: command.Amount,
	}
	log.Printf("%#v", event)
	return a.eventProducer.Produce(event)
}

func (a accountService) WithDrawFund(command command.WithDrawFundCommand) error {
	if command.ID == "" || command.Amount == 0 {
		return errors.New("bad request")
	}
	event := events.WithdrawFundEvent{
		ID:     command.ID,
		Amount: command.Amount,
	}
	log.Printf("%#v", event)
	return a.eventProducer.Produce(event)
}

func (a accountService) CloseAccount(command command.CloseAccountCommand) error {
	if command.ID == "" {
		return errors.New("bad request")
	}
	event := events.CloseAccountEvent{ID: command.ID}
	log.Printf("%#v", event)
	return a.eventProducer.Produce(event)
}
