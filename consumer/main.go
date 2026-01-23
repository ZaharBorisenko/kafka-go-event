package main

import (
	"github.com/IBM/sarama"
	"github.com/spf13/viper"
	"strings"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigFile("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
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
