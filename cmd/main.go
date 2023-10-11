package main

import (
	"context"
	"fmt"

	appKafka "github.com/nikandpro/kafka-fio-server/pkg/kafka"
	"github.com/nikandpro/kafka-fio-server/pkg/service"
)

func main() {
	fmt.Println("start main...")
	ctx := context.Background()
	dataKafka := make(chan string, 10)
	failDataKafka := make(chan string, 10)

	go appKafka.Produce(ctx)

	service := service.NewService(ctx, dataKafka, failDataKafka)
	go func() {
		err := service.StartService()
		if err != nil {
			fmt.Println("service error ", err)
		}
	}()

	appKafka.Consume(ctx, dataKafka)

	// person := mysql.Person{}
	// fmt.Println(person)

}
