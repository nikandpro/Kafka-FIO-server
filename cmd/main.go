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
	log.Println("init cfg")
	db, err := sql.Open("postgres", cfg.DBPath)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Exec("create table users(id serial primary key, name varchar(50), surname varchar(50), patronymic varchar(50), age int, gender varchar(50));")
	db.Exec("create table country(id VARCHAR(50) UNIQUE, probability float);")
	db.Exec("create table users_country(id serial primary key, user_id int, country_id varchar(50), foreign key (user_id) references users (id), foreign key (country_id) references country (id));")
	log.Println("migration db")
}

//create table users_country(id serial primary key, user_id int, country_id varchar(50), foreign key (user_id) references users (id), foreign key (country_id) references country (id));

func main() {
	log.Println("start...")
	ctx := context.Background()
	dataKafka := make(chan []byte, 10)
	failDataKafka := make(chan string, 10)

	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	db, err := postgres.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		err := appKafka.Produce(ctx)
		if err != nil {
			log.Fatal(err)
		}

	}()

	service := service.NewService(ctx, dataKafka, failDataKafka, db)
	go func() {
		err := service.StartService()
		if err != nil {
			fmt.Println("service error ", err)
		}
	}()

	server := server.NewServer(ctx, db, *cfg)

	go func() {
		err := server.StartServer()
		if err != nil {
			log.Fatal("server error ", err)
		}
	}()

	err = appKafka.Consume(ctx, dataKafka)
	if err != nil {
		log.Fatal(err)
	}

}
