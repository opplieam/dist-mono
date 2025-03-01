package store

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	db "github.com/opplieam/dist-mono/db/sqlc"
	catApi "github.com/opplieam/dist-mono/internal/category/api"
	"github.com/opplieam/dist-mono/internal/user/api"
)

var (
	ErrUserNotFound    = errors.New("user not found")
	ErrCategoryConn    = errors.New("category service down")
	ErrNoCategoryFound = errors.New("no category found for this user")
)

type Store struct {
	db        *db.Queries
	catClient *catApi.Client
}

func NewStore(q *db.Queries, c *catApi.Client) *Store {
	s := &Store{
		db:        q,
		catClient: c,
	}
	return s
}

func (s *Store) CreateUser(ctx context.Context, name, email string) (int, error) {
	userId, err := s.db.CreateUser(ctx, db.CreateUserParams{
		Name:  name,
		Email: email,
	})
	if err != nil {
		return 0, err
	}
	return int(userId), nil
}

func (s *Store) GetAllUsers(ctx context.Context) (*api.GetAllUsersOKApplicationJSON, error) {
	users, err := s.db.GetAllUsers(ctx)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUserNotFound
		}
	}
	var usersApi api.GetAllUsersOKApplicationJSON
	for _, user := range users {
		usersApi = append(usersApi, api.User{
			ID:    int(user.ID),
			Name:  user.Name,
			Email: user.Email,
		})
	}
	return &usersApi, nil
}

func (s *Store) GetUserCategory(ctx context.Context, userID int) (*api.UserCategory, error) {
	user, err := s.db.GetUserByID(ctx, int32(userID))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUserNotFound
		}
	}
	catRes, err := s.catClient.GetCategoryById(ctx, catApi.GetCategoryByIdParams{ID: userID})

	if err != nil {
		var apiErr *catApi.ErrorStatusCode
		if errors.As(err, &apiErr) {
			return nil, ErrNoCategoryFound
		}
		return nil, ErrCategoryConn
	}

	switch res := catRes.(type) {
	case *catApi.Category:
		return &api.UserCategory{
			ID:       int(user.ID),
			Name:     user.Name,
			Category: res.GetName(),
		}, nil
	default:
		return nil, fmt.Errorf("unexpected response from category service")
	}
}
