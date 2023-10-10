package kafka

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
)

func Produce(ctx context.Context) {
	fmt.Println("produce")
	i := 0

	conf := kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "FIO",
	}

	writer := kafka.NewWriter(conf)

	for {
		err := writer.WriteMessages(ctx, kafka.Message{
			Key:   []byte(strconv.Itoa(i)),
			Value: []byte("this is Message: " + strconv.Itoa(i)),
		})
		if err != nil {
			fmt.Println("Some produce eroor: ", err)
		}
		i++
		time.Sleep(time.Second * 5)
	}
}
