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
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type UserHandler struct {
	hServer *http.Server
}

var _ api.Handler = (*UserHandler)(nil)

func NewUserHandler() *UserHandler {
	return &UserHandler{}
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
	return &api.User{
		ID:    0,
		Name:  "Test",
		Email: "Test@example.com",
	}, nil
}

func (u *UserHandler) GetAllUsers(ctx context.Context) (api.GetAllUsersRes, error) {
	var apiUsers api.GetAllUsersOKApplicationJSON
	for i := 0; i < 2; i++ {
		apiUsers = append(apiUsers, api.User{
			ID:    i,
			Name:  "",
			Email: "",
		})
	}
	return &apiUsers, nil
}

func (u *UserHandler) GetUserById(ctx context.Context, params api.GetUserByIdParams) (api.GetUserByIdRes, error) {
	return &api.UserCategory{
		ID:       0,
		Name:     "Test",
		Category: "Test",
	}, nil
}

func (u *UserHandler) NewError(ctx context.Context, err error) *api.ErrorStatusCode {
	switch {
	case errors.Is(err, ErrUserNotFound):
		return &api.ErrorStatusCode{
			StatusCode: http.StatusNotFound,
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
