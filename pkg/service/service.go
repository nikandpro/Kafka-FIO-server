package service

import (
	"context"
	"fmt"
	"log"

	"github.com/nikandpro/kafka-fio-server/pkg/database"
)

type Service struct {
	ctx context.Context

	dataKafka     chan []byte
	failDataKafka chan string

	db database.Database
}

func NewService(ctx context.Context, data chan []byte, failData chan string, db database.Database) *Service {
	return &Service{ctx: ctx, dataKafka: data, failDataKafka: failData, db: db}
}

func (s *Service) StartService() error {
	fmt.Println("Start service...")

	for k := range s.dataKafka {

		user, err := IsCorrect(k)
		if err != nil {
			log.Fatal("Not Correct user", err)
			return err
		}

		s.enrichment(user)
		// fmt.Println("kafka message: ", string(k))
	}

	return nil
}
