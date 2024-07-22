package repository

import (
	"context"

	"github.com/yelaco/warble/internal/warble/service"
	"github.com/yelaco/warble/internal/warble/service/entity"
	"gorm.io/gorm"
)

type warbleRepo struct {
	db *gorm.DB
}

func NewWarbleRepo(db *gorm.DB) service.WarbleRepository {
	return warbleRepo{
		db: db,
	}
}

func (wr warbleRepo) InsertWarble(ctx context.Context, warble entity.Warble) error {
	// insert warble

	return nil
}

func (wr warbleRepo) GetWarbleByID(ctx context.Context, id string) (entity.Warble, error) {
	return entity.Warble{}, nil
}
func (wr warbleRepo) GetWarblesByAuthorID(ctx context.Context, authorID string) ([]entity.Warble, error) {
	return []entity.Warble{}, nil
}
