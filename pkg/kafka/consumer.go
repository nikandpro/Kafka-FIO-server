package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
)

func Consume(ctx context.Context, dataKafka chan []byte) error {

	conf := kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "FIO",
		MaxBytes: 10,
	}

	reader := kafka.NewReader(conf)

	for {
		m, err := reader.ReadMessage(ctx)
		if err != nil {
			return err
		}
		dataKafka <- m.Value
	}
}
