package kafka

import (
	"context"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
)

func Produce(ctx context.Context) error {
	i := 0
	jso := `
	{
		"name": "John",
		"surname": "Doe",
		"patronymic": "Jonson"
	}
	`
	conf := kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "FIO",
	}

	writer := kafka.NewWriter(conf)

	for {
		err := writer.WriteMessages(ctx, kafka.Message{
			Key:   []byte(strconv.Itoa(i)),
			Value: []byte(jso),
		})
		if err != nil {
			return err
		}
		i++
		time.Sleep(time.Second * 5)
	}
}
