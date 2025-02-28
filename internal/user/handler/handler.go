package handler

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/opplieam/dist-mono/internal/user/api"
	"github.com/opplieam/dist-mono/internal/user/store"
)

type Storer interface {
	CreateUser(ctx context.Context, name, email string) (int, error)
	GetAllUsers(ctx context.Context) (*api.GetAllUsersOKApplicationJSON, error)
	GetUserCategory(ctx context.Context, userID int) (*api.UserCategory, error)
}

type UserHandler struct {
	hServer *http.Server
	store   Storer
}

var _ api.Handler = (*UserHandler)(nil)

func NewUserHandler(s Storer) *UserHandler {
	return &UserHandler{
		store: s,
	}
}

func (u *UserHandler) Start() (chan os.Signal, error) {
	srv, err := api.NewServer(u)
	if err != nil {
		return nil, fmt.Errorf("failed to create server: %w", err)
	}
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Mount("/v1", http.StripPrefix("/v1", srv))
	u.hServer = &http.Server{
		Addr:    ":3000",
		Handler: r,
	}

	go func() {
		log.Printf("User service listening on %s", u.hServer.Addr)
		if hErr := u.hServer.ListenAndServe(); !errors.Is(hErr, http.ErrServerClosed) {
			log.Fatalf("Failed to listen on %s: %v", u.hServer.Addr, hErr)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	return sigChan, nil
}

func (u *UserHandler) Shutdown() error {
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return u.hServer.Shutdown(shutdownCtx)
}

func (u *UserHandler) CreateUser(ctx context.Context, req *api.User) (api.CreateUserRes, error) {
	name := req.GetName()
	email := req.GetEmail()
	userId, err := u.store.CreateUser(ctx, name, email)
	if err != nil {
		return nil, err
	}
	return &api.User{
		ID:    userId,
		Name:  name,
		Email: email,
	}, nil
}

func (u *UserHandler) GetAllUsers(ctx context.Context) (api.GetAllUsersRes, error) {
	users, err := u.store.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserHandler) GetUserById(ctx context.Context, params api.GetUserByIdParams) (api.GetUserByIdRes, error) {
	userCat, err := u.store.GetUserCategory(ctx, params.ID)
	if err != nil {
		return nil, err
	}
	return userCat, nil
}

func (u *UserHandler) NewError(ctx context.Context, err error) *api.ErrorStatusCode {
	switch {
	case errors.Is(err, store.ErrUserNotFound):
		return &api.ErrorStatusCode{
			StatusCode: http.StatusNotFound,
			Response: api.Error{
				Message: err.Error(),
			},
		}
	case errors.Is(err, store.ErrCategoryConn):
		return &api.ErrorStatusCode{
			StatusCode: http.StatusServiceUnavailable,
			Response: api.Error{
				Message: err.Error(),
			},
		}
	default:
		return &api.ErrorStatusCode{
			StatusCode: http.StatusInternalServerError,
			Response: api.Error{
				Message: err.Error(),
			},
		}
	}
}
