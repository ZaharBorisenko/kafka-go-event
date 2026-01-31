package main

import (
	"context"
	"events"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/spf13/viper"
	"strings"
)

type TestHandler struct{}

func (h *TestHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (h *TestHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h *TestHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("Message claimed: value = %s, timestamp = %v, topic = %s\n",
			string(msg.Value), msg.Timestamp, msg.Topic)
		// Здесь можно добавить десериализацию через json.Unmarshal в структуры из events
		sess.MarkMessage(msg, "")
	}
	return nil
}

func main() {
	// Настройка viper (аналогично твоему коду)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	consumer, err := sarama.NewConsumerGroup(
		viper.GetStringSlice("kafka.servers"),
		"test-service-group",
		nil,
	)
	if err != nil {
		panic(err)
	}

	handler := &TestHandler{}
	fmt.Println("Test service consumer started...")

	for {
		err := consumer.Consume(context.Background(), events.Topics, handler)
		if err != nil {
			fmt.Printf("Error from consumer: %v\n", err)
		}
	}
}
