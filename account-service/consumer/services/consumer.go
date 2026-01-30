package services

import "github.com/IBM/sarama"

type consumerHandler struct {
	eventHandler EventHandler
}

func NewConsumerHandler(handler EventHandler) sarama.ConsumerGroupHandler {
	return consumerHandler{eventHandler: handler}
}

func (c consumerHandler) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (c consumerHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (c consumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		c.eventHandler.Handle(msg.Topic, msg.Value)
		session.MarkMessage(msg, "")
	}

	return nil
}
