package services

import (
	repo "consumer/repositories"
	"encoding/json"
	"events"
	"fmt"
	"log"
	"reflect"
)

type EventHandler interface {
	Handle(topic string, eventBytes []byte)
}

type accountEventHandler struct {
	accountRepo repo.AccountRepository
}

func NewAccountEventHandler(accountRepo repo.AccountRepository) EventHandler {
	return accountEventHandler{accountRepo: accountRepo}
}

func (s accountEventHandler) Handle(topic string, eventBytes []byte) {
	switch topic {
	case reflect.TypeOf(events.OpenAccountEvent{}).Name():
		event := events.OpenAccountEvent{}
		err := json.Unmarshal(eventBytes, &event)
		if err != nil {
			log.Println(err)
			return
		}
		account := repo.BankAccount{
			ID:            event.ID,
			AccountHolder: event.AccountHolder,
			AccountType:   event.AccountType,
			Balance:       event.OpeningBalance,
		}
		err = s.accountRepo.Save(account)
		if err != nil {
			log.Println(err)
			return
		}

		log.Printf("[%v] %#v", topic, event)

	case reflect.TypeOf(events.DepositFundEvent{}).Name():
		event := events.DepositFundEvent{}
		err := json.Unmarshal(eventBytes, &event)
		if err != nil {
			log.Println(err)
			return
		}

		account, err := s.accountRepo.FindByID(event.ID)
		if err != nil {
			log.Println(err)
			return
		}
		account.Balance += event.Amount

		err = s.accountRepo.Save(account)
		if err != nil {
			log.Println(err)
			return
		}

		log.Printf("[%v] %#v", topic, event)

	case reflect.TypeOf(events.WithdrawFundEvent{}).Name():
		event := events.WithdrawFundEvent{}
		err := json.Unmarshal(eventBytes, &event)
		if err != nil {
			log.Println(err)
			return
		}

		account, err := s.accountRepo.FindByID(event.ID)
		if err != nil {
			log.Println(err)
			return
		}
		account.Balance -= event.Amount

		err = s.accountRepo.Save(account)
		if err != nil {
			log.Println(err)
			return
		}

		log.Printf("[%v] %#v", topic, event)

	case reflect.TypeOf(events.CloseAccountEvent{}).Name():
		event := events.CloseAccountEvent{}
		err := json.Unmarshal(eventBytes, &event)
		if err != nil {
			fmt.Println(err)
			return
		}

		err = s.accountRepo.Delete(event.ID)
		if err != nil {
			fmt.Println(err)
			return
		}

		log.Printf("[%v] %#v", topic, event)

	default:
		log.Println("no event handler")
	}
}
