package main

import (
	"context"
	"fmt"

	"github.com/nikandpro/kafka-fio-server/pkg/database/postgres"
	appKafka "github.com/nikandpro/kafka-fio-server/pkg/kafka"
	"github.com/nikandpro/kafka-fio-server/pkg/service"
)

func main() {
	fmt.Println("start main...")
	db, err := postgres.New()
	if err != nil {
		panic(err)
	}
	// db.Get()

	ctx := context.Background()
	dataKafka := make(chan []byte, 10)
	failDataKafka := make(chan string, 10)

	go appKafka.Produce(ctx)

	service := service.NewService(ctx, dataKafka, failDataKafka, db)
	go func() {
		err := service.StartService()
		if err != nil {
			fmt.Println("service error ", err)
		}
	}()

	appKafka.Consume(ctx, dataKafka)

}
