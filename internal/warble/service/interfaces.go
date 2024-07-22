package service

import (
	"context"

	"github.com/yelaco/warble/internal/warble/service/entity"
)

type WarbleUsecase interface {
	CreateWarble(ctx context.Context, authorID, body string) (entity.Warble, error)
	GetWarbleByID(ctx context.Context, warbleID string) (entity.Warble, error)
	GetWarblesByAuthorID(ctx context.Context, authorID string) ([]entity.Warble, error)
}

type WarbleRepository interface {
	InsertWarble(ctx context.Context, warble entity.Warble) error
	GetWarbleByID(ctx context.Context, id string) (entity.Warble, error)
	GetWarblesByAuthorID(ctx context.Context, authorID string) ([]entity.Warble, error)
}
