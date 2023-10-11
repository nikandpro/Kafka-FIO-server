package service

import (
	"context"
	"encoding/json"
	"fmt"
)

type Service struct {
	ctx context.Context

	inputKafka     chan string
	failQueueKafka chan string
}

type Person struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

func NewService(ctx context.Context, in chan string, fail chan string) *Service {
	return &Service{ctx: ctx, inputKafka: in, failQueueKafka: fail}
}

func (s *Service) StartService() error {
	fmt.Println("Start service...")

	for k := range s.inputKafka {

		person, err := s.serializationJSON(k)
		if err != nil {
			return err
		}
		fmt.Println("kafka message: ", k, "person = ", person)
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
