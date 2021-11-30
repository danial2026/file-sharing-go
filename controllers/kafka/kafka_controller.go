package kafka

// appKafka "github.com/danial2026/file-sharing-go/controllers/kafka"

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

const (
	topic          = "message-log"
	broker1Address = "localhost:9092"
	groupID        = "GroupID"
)

func Consume(ctx context.Context) {
	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker1Address},
		Topic:   topic,
		GroupID: groupID,
	})
	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			// panic("could not read message " + err.Error())
		} else {
			//try to map msg to a const with same value from parent-module
			fmt.Println("received: ", string(msg.Value))
		}
	}
}
