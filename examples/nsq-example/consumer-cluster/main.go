package main

import (
	"github.com/nsqio/go-nsq"
	"log"
	"time"
)

func main() {
	cfg := nsq.NewConfig()

	consumer, err := nsq.NewConsumer("topic", "channel", cfg)
	if err != nil {
		log.Fatal(err)
	}

	consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Printf("message: %v", string(message.Body))
		return nil
	}))

	if err := consumer.ConnectToNSQLookupd("localhost:4161"); err != nil {
		log.Fatal(err)
	}

	time.Sleep(10 * time.Second)
	consumer.Stop()
}
