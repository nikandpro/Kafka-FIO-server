package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/nikandpro/kafka-fio-server/pkg/config"
	"github.com/nikandpro/kafka-fio-server/pkg/database/postgres"
	appKafka "github.com/nikandpro/kafka-fio-server/pkg/kafka"
	"github.com/nikandpro/kafka-fio-server/pkg/server"
	"github.com/nikandpro/kafka-fio-server/pkg/service"
)

func init() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}
	db, err := sql.Open("postgres", cfg.DBPath)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Exec("create table users(id serial primary key, name varchar(50), surname varchar(50), patronymic varchar(50), agify int, genderize varchar(50), nationalize varchar(50));")

}

func main() {
	log.Println("start main...")

	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	db, err := postgres.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

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
