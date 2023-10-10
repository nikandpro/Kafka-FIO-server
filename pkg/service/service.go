package service

import (
	"context"
	"fmt"
)

type Service struct {
	ctx context.Context

	inputKafka     chan string
	failQueueKafka chan string
}

func NewService(ctx context.Context, in chan string, fail chan string) *Service {
	return &Service{ctx: ctx, inputKafka: in, failQueueKafka: fail}
}

func (s *Service) StartService() error {
	fmt.Println("Start service...")

	for k := range s.inputKafka {
		fmt.Println("kafka message: ", k)
	}

	return nil
}
