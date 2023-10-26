package main

import (
	"context"
	"fmt"
	"log"

	"github.com/nikandpro/kafka-fio-server/pkg/config"
	"github.com/nikandpro/kafka-fio-server/pkg/database/postgres"
	appKafka "github.com/nikandpro/kafka-fio-server/pkg/kafka"
	"github.com/nikandpro/kafka-fio-server/pkg/server"
	"github.com/nikandpro/kafka-fio-server/pkg/service"
)

func main() {
	fmt.Println("start main...")

	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	db, err := postgres.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	db.Get()

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

	server := server.NewServer(ctx)

	go func() {
		err := server.StartServer()
		if err != nil {
			log.Fatal("server error", err)
		}
	}()

	appKafka.Consume(ctx, dataKafka)

}
