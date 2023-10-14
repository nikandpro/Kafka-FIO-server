package service

import (
	"context"
	"fmt"

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

		err := s.db.Create(k)
		if err != nil {
			fmt.Println("StartService error", err)
			return err
		}
		// fmt.Println("kafka message: ", string(k))
	}

	return nil
}

// func (s *Service) serializationJSON(str string) (Person, error) {
// 	person := Person{}

// 	err := json.Unmarshal([]byte(str), &person)
// 	if err != nil {
// 		fmt.Println("serializ error", err)

// 	}
// 	return person, nil
// }
