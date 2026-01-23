package main

import (
	"fmt"
	"github.com/IBM/sarama"
	"time"
)

func main() {
	addr := []string{"localhost:29092"}

	producer, err := sarama.NewSyncProducer(addr, nil)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	fmt.Println("producer start sending message...")
	msg := sarama.ProducerMessage{
		Topic:     "test-topic",
		Value:     sarama.StringEncoder("Hello consumer with key"),
		Key:       sarama.StringEncoder("user-1"),
		Timestamp: time.Time{},
	}

	p, o, err := producer.SendMessage(&msg)
	if err != nil {
		panic(err)
	}

	fmt.Printf("partition: %v, offset: %v", p, o)
}
