package main

import (
	"context"
	"log"
	"os"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

func main() {
	go Producer()
	go Consumer()
	select {}
}

func Producer() {
	// Create a Producer instance
	producer, err := rocketmq.NewProducer(
		producer.WithNameServer([]string{"127.0.0.1:9876"}),
		// producer.WithNsResolver(primitive.NewPassthroughResolver([]string{"127.0.0.1:9876"})),
		producer.WithRetry(2),
		producer.WithCredentials(primitive.Credentials{
			AccessKey: "normal",
			SecretKey: "normal",
		}),
		producer.WithRetry(2),
		producer.WithGroupName("GID_test"),
	)
	if err != nil {
		log.Printf("producer init fail: %v", err)
		os.Exit(1)
	}
	// Start the producer
	err = producer.Start()
	if err != nil {
		log.Printf("producer start fail: %v", err)
		os.Exit(1)
	}
	// Send message with sync
	result, err := producer.SendSync(context.Background(), &primitive.Message{
		Topic: "test",
		Body:  []byte("Hello RocketMQ Go Client!"),
	})
	if err != nil {
		log.Printf("producer sendSync message fail: %v", err)
	} else {
		// Do something with result
		log.Printf("producer send message success: %v", result)
	}
	// Shutdown the producer
	err = producer.Shutdown()
	if err != nil {
		log.Printf("producer shutdown fail: %v", err)
	}

}

func Consumer() {
	// Create a PushConsumer instance
	c, err := rocketmq.NewPushConsumer(
		consumer.WithNameServer([]string{"127.0.0.1:9876"}),
		consumer.WithConsumerModel(consumer.Clustering),
		consumer.WithGroupName("GID_test"),
	)
	if err != nil {
		log.Printf("create consumer fail: %v", err)
		os.Exit(1)
	}
	// Subscribe a topic(only support one topic now), and define your consuming function
	err = c.Subscribe("test", consumer.MessageSelector{}, func(ctx context.Context,
		msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		// Just prints to standard output.
		for _, msg := range msgs {
			message := string(msg.Body)
			log.Printf("consumer recieve: %v\n", message)
		}
		return consumer.ConsumeSuccess, nil
	})
	if err != nil {
		log.Printf("consumer subscribe %s fail: %v", "test", err)
		os.Exit(1)
	}
	// Start the consumer(NOTE: MUST after subscribe)
	err = c.Start()
	if err != nil {
		log.Printf("consumer start fail: %v", err)
	}
}
