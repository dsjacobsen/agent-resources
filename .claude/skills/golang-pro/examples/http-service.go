// Package main demonstrates a production-ready HTTP service structure
// This is an example file for the golang-pro skill
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// =============================================================================
// Domain Models
// =============================================================================

// User represents a domain user entity
type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// =============================================================================
// Errors
// =============================================================================

var (
	ErrNotFound     = errors.New("not found")
	ErrBadRequest   = errors.New("bad request")
	ErrUnauthorized = errors.New("unauthorized")
)

// =============================================================================
// Repository Layer (Data Access)
// =============================================================================

// UserRepository defines the interface for user data operations
type UserRepository interface {
	GetByID(ctx context.Context, id string) (*User, error)
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id string) error
}

// InMemoryUserRepository is a simple in-memory implementation
type InMemoryUserRepository struct {
	users map[string]*User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]*User),
	}
}

func (r *InMemoryUserRepository) GetByID(ctx context.Context, id string) (*User, error) {
	user, ok := r.users[id]
	if !ok {
		return nil, fmt.Errorf("user %s: %w", id, ErrNotFound)
	}
	return user, nil
}

func (r *InMemoryUserRepository) Create(ctx context.Context, user *User) error {
	r.users[user.ID] = user
	return nil
}

func (r *InMemoryUserRepository) Update(ctx context.Context, user *User) error {
	if _, ok := r.users[user.ID]; !ok {
		return fmt.Errorf("user %s: %w", user.ID, ErrNotFound)
	}
	r.users[user.ID] = user
	return nil
}

func (r *InMemoryUserRepository) Delete(ctx context.Context, id string) error {
	if _, ok := r.users[id]; !ok {
		return fmt.Errorf("user %s: %w", id, ErrNotFound)
	}
	delete(r.users, id)
	return nil
}

// =============================================================================
// Service Layer (Business Logic)
// =============================================================================

// UserService handles user-related business logic
type UserService struct {
	repo   UserRepository
	logger *slog.Logger
}

func NewUserService(repo UserRepository, logger *slog.Logger) *UserService {
	return &UserService{
		repo:   repo,
		logger: logger,
	}
}

func (s *UserService) GetUser(ctx context.Context, id string) (*User, error) {
	s.logger.InfoContext(ctx, "fetching user", slog.String("user_id", id))

	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("getting user: %w", err)
	}

	return user, nil
}

func (s *UserService) CreateUser(ctx context.Context, name, email string) (*User, error) {
	user := &User{
		ID:        fmt.Sprintf("user_%d", time.Now().UnixNano()),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return nil, fmt.Errorf("creating user: %w", err)
	}

	s.logger.InfoContext(ctx, "user created", slog.String("user_id", user.ID))
	return user, nil
}

// =============================================================================
// HTTP Handlers
// =============================================================================

// UserHandler handles HTTP requests for users
type UserHandler struct {
	service *UserService
	logger  *slog.Logger
}

func NewUserHandler(service *UserService, logger *slog.Logger) *UserHandler {
	return &UserHandler{
		service: service,
		logger:  logger,
	}
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id") // Go 1.22+ routing

	user, err := h.service.GetUser(r.Context(), id)
	if err != nil {
		h.handleError(w, r, err)
		return
	}

	h.respondJSON(w, http.StatusOK, user)
}

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.handleError(w, r, fmt.Errorf("%w: invalid JSON", ErrBadRequest))
		return
	}

	user, err := h.service.CreateUser(r.Context(), req.Name, req.Email)
	if err != nil {
		h.handleError(w, r, err)
		return
	}

	h.respondJSON(w, http.StatusCreated, user)
}

func (h *UserHandler) handleError(w http.ResponseWriter, r *http.Request, err error) {
	var status int
	var message string

	switch {
	case errors.Is(err, ErrNotFound):
		status = http.StatusNotFound
		message = "resource not found"
	case errors.Is(err, ErrBadRequest):
		status = http.StatusBadRequest
		message = err.Error()
	case errors.Is(err, ErrUnauthorized):
		status = http.StatusUnauthorized
		message = "unauthorized"
	default:
		status = http.StatusInternalServerError
		message = "internal server error"
		h.logger.ErrorContext(r.Context(), "internal error", slog.Any("error", err))
	}

	h.respondJSON(w, status, map[string]string{"error": message})
}

func (h *UserHandler) respondJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		h.logger.Error("failed to encode response", slog.Any("error", err))
	}
}

// =============================================================================
// Middleware
// =============================================================================

func LoggingMiddleware(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			next.ServeHTTP(w, r)

			logger.Info("request completed",
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.Duration("duration", time.Since(start)),
			)
		})
	}
}

func RecoveryMiddleware(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					logger.Error("panic recovered",
						slog.Any("error", err),
						slog.String("path", r.URL.Path),
					)
					http.Error(w, "internal server error", http.StatusInternalServerError)
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}

// =============================================================================
// Router Setup
// =============================================================================

func NewRouter(userHandler *UserHandler, logger *slog.Logger) http.Handler {
	mux := http.NewServeMux()

	// Health check
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	// User routes (Go 1.22+ routing)
	mux.HandleFunc("GET /users/{id}", userHandler.GetUser)
	mux.HandleFunc("POST /users", userHandler.CreateUser)

	// Apply middleware
	var handler http.Handler = mux
	handler = LoggingMiddleware(logger)(handler)
	handler = RecoveryMiddleware(logger)(handler)

	return handler
}

// =============================================================================
// Server Configuration
// =============================================================================

type ServerConfig struct {
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

func DefaultServerConfig() ServerConfig {
	return ServerConfig{
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
}

func NewHTTPServer(cfg ServerConfig, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:         cfg.Addr,
		Handler:      handler,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}
}

// =============================================================================
// Main Application
// =============================================================================

func main() {
	// Setup structured logging
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)

	// Setup signal handling for graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Initialize dependencies
	repo := NewInMemoryUserRepository()
	userService := NewUserService(repo, logger)
	userHandler := NewUserHandler(userService, logger)

	// Setup router and server
	router := NewRouter(userHandler, logger)
	cfg := DefaultServerConfig()
	srv := NewHTTPServer(cfg, router)

	// Start server in background
	go func() {
		logger.Info("server starting", slog.String("addr", srv.Addr))
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("server error", slog.Any("error", err))
		}
	}()

	// Wait for shutdown signal
	<-ctx.Done()
	logger.Info("shutdown signal received")

	// Graceful shutdown with timeout
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		logger.Error("shutdown error", slog.Any("error", err))
	}

	logger.Info("server stopped gracefully")
}
