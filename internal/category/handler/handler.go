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
	"github.com/opplieam/dist-mono/internal/category/api"
)

var (
	ErrCategoryNotFound = errors.New("category not found")
)

type CategoryHandler struct {
	hServer *http.Server
}

var _ api.Handler = (*CategoryHandler)(nil)

func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{}
}

func (h *CategoryHandler) Start() (chan os.Signal, error) {
	srv, err := api.NewServer(h)
	if err != nil {
		return nil, fmt.Errorf("failed to create server: %w", err)
	}
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Mount("/v1", http.StripPrefix("/v1", srv))
	h.hServer = &http.Server{
		Addr:    ":4000",
		Handler: r,
	}

	go func() {
		log.Printf("Category service listening on %s", h.hServer.Addr)
		if hErr := h.hServer.ListenAndServe(); !errors.Is(hErr, http.ErrServerClosed) {
			log.Fatalf("Failed to listen on %s: %v", h.hServer.Addr, hErr)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	return sigChan, nil
}

func (h *CategoryHandler) Shutdown() error {
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return h.hServer.Shutdown(shutdownCtx)
}

func (h *CategoryHandler) GetCategoryById(ctx context.Context, params api.GetCategoryByIdParams) (api.GetCategoryByIdRes, error) {
	// TODO: Add real database call
	log.Printf("GetCategoryById: %v", params)
	return &api.Category{
		ID:   0,
		Name: "Test",
	}, nil
}

func (h *CategoryHandler) NewError(ctx context.Context, err error) *api.ErrorStatusCode {
	switch {
	case errors.Is(err, ErrCategoryNotFound):
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
