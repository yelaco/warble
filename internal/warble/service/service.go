package service

import (
	"context"

	"github.com/yelaco/warble/internal/warble/service/entity"
)

type WarbleService struct {
	repository WarbleRepository
}

func (wu *WarbleService) CreateWarble(ctx context.Context, authorID, body string) (entity.Warble, error) {
	warble := entity.NewWarble(authorID, body)

	err := wu.repository.InsertWarble(ctx, warble)
	if err != nil {
		return entity.Warble{}, nil
	}

	return warble, nil
}

func (wu *WarbleService) GetWarbleByID(ctx context.Context, warbleID string) (entity.Warble, error) {
	return entity.Warble{}, nil
}

func (wu *WarbleService) GetWarblesByAuthorID(ctx context.Context, authorID string) ([]entity.Warble, error) {
	return []entity.Warble{}, nil
}
