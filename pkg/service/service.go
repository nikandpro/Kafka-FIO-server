package service

import (
	"context"
	"encoding/json"
	"fmt"
)

type Service struct {
	ctx context.Context

	dataKafka     chan []byte
	failDataKafka chan string
}

type Person struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

func NewService(ctx context.Context, data chan []byte, failData chan string) *Service {
	return &Service{ctx: ctx, dataKafka: data, failDataKafka: failData}
}

func (s *Service) StartService() error {
	fmt.Println("Start service...")

	for k := range s.dataKafka {

		person, err := s.serializationJSON(string(k))
		if err != nil {
			return err
		}
		fmt.Println("kafka message: ", string(k), "person = ", person)
	}

	return nil
}

func (s *Service) serializationJSON(str string) (Person, error) {
	person := Person{}

	err := json.Unmarshal([]byte(str), &person)
	if err != nil {
		fmt.Println("serializ error", err)

	}
	return person, nil
}
