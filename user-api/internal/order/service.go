package order

import (
	"context"
	"encoding/json"
	"user-api/internal/kafka"
)

type Service struct {
	repo     *Repository
	producer *kafka.Producer
}

func NewService(repo *Repository, producer *kafka.Producer) *Service {
	return &Service{repo: repo, producer: producer}
}

func (s *Service) CreateOrder(ctx context.Context, o Order) error {
	if err := s.repo.Create(o); err != nil {
		return err
	}

	// Publish event
	event, _ := json.Marshal(o)
	return s.producer.Publish(ctx, "order", string(event))
}
