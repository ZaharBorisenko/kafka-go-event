package main

import (
	"consumer/repositories"
	"consumer/services"
	"context"
	"events"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

func initDatabase() *gorm.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.database"),
	)

	dial := mysql.Open(dsn)

	db, err := gorm.Open(dial, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}

	return db
}

func main() {
	consumer, err := sarama.NewConsumerGroup(
		viper.GetStringSlice("kafka.servers"),
		viper.GetString("kafka.group"),
		nil,
	)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	db := initDatabase()

	accountRepo := repositories.NewAccountRepository(db)
	accountEventHandler := services.NewAccountEventHandler(accountRepo)
	accountConsumeHandler := services.NewConsumerHandler(accountEventHandler)

	fmt.Println("account consumer started...")
	for {
		consumer.Consume(context.Background(), events.Topics, accountConsumeHandler)
	}

}

//func main() {
//	addr := []string{"localhost:29092"}
//
//	consumer, err := sarama.NewConsumer(addr, nil)
//	if err != nil {
//		panic(err)
//	}
//	defer consumer.Close()
//
//	partitionConsumer, err := consumer.ConsumePartition("test-topic", 0, sarama.OffsetNewest)
//	if err != nil {
//		panic(err)
//	}
//	defer partitionConsumer.Close()
//
//	fmt.Println("consumer start.")
//	for {
//		select {
//		case err := <-partitionConsumer.Errors():
//			fmt.Println(err)
//		case msg := <-partitionConsumer.Messages():
//			fmt.Println(msg.Topic, ":", string(msg.Value))
//
//		}
//	}
//}
