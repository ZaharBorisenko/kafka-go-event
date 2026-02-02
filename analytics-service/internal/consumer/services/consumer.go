package services

import (
	"fmt"
	"github.com/IBM/sarama"
)

type consumerHandler struct {
	eventHandler EventHandler
}

func NewConsumerHandler(handler EventHandler) sarama.ConsumerGroupHandler {
	return &consumerHandler{eventHandler: handler}
}

func (h *consumerHandler) Setup(session sarama.ConsumerGroupSession) error {
	fmt.Println("Consumer group session started")
	return nil
}

func (h *consumerHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	fmt.Println("Consumer group session ended")
	return nil
}

func (h *consumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		h.eventHandler.Handle(msg.Topic, msg.Value)

		session.MarkMessage(msg, "")
	}
	return nil
}
