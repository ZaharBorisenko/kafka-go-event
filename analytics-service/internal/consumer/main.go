package main

import (
	"analytics-service/internal/consumer/services"
	"analytics-service/internal/repository"
	"context"
	"events"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/spf13/viper"
	"strings"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func main() {
	conn, err := repository.ClickHouseConnect()
	if err != nil {
		panic(err)
	}

	analyticsRepo := repository.NewAnalyticsRepository(conn)
	analyticsEventHandler := services.NewAnalyticsEventHandler(analyticsRepo)
	handler := services.NewConsumerHandler(analyticsEventHandler)

	consumer, err := sarama.NewConsumerGroup(
		viper.GetStringSlice("kafka.servers"),
		viper.GetString("kafka.group"),
		nil,
	)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	fmt.Println("Analytics service consumer started...")

	for {
		// Используем наш собранный handler
		err := consumer.Consume(context.Background(), events.Topics, handler)
		if err != nil {
			fmt.Printf("Error from consumer: %v\n", err)
		}
	}

}
