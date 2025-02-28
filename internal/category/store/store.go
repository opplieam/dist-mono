package store

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	db "github.com/opplieam/dist-mono/db/sqlc"
)

var (
	ErrCategoryNotFound = errors.New("category not found")
)

type Store struct {
	db *db.Queries
}

func NewStore(q *db.Queries) *Store {
	return &Store{
		db: q,
	}
}

type CategoryResult struct {
	ID   int
	Name string
}

func (s *Store) GetCategoryByID(ctx context.Context, userID int) (*CategoryResult, error) {
	res, err := s.db.GetCategoryByID(ctx, int32(userID))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrCategoryNotFound
		}
	}
	return &CategoryResult{
		ID:   int(res.ID),
		Name: res.Name,
	}, nil
}
