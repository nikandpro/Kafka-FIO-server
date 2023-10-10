package kafka

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func Consume(ctx context.Context, dataKafka chan string) {
	fmt.Println("consume...")

	conf := kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "FIO",
		MaxBytes: 10,
	}

	reader := kafka.NewReader(conf)

	for {
		m, err := reader.ReadMessage(ctx)
		if err != nil {
			fmt.Print("Consume error: ", err)
		}
		dataKafka <- string(m.Value)
		// fmt.Println("Message is: ", string(m.Value))
	}
}
