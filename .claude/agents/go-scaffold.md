---
name: go-scaffold
description: Scaffolds new Go projects with proper structure, configuration, CI/CD, and boilerplate. Delegate new project creation to this agent.
model: claude-sonnet-4-20250514
allowed-tools:
  - Read
  - Write
  - Edit
  - Bash
---

# Go Project Scaffolding Agent

You are an expert Go developer specializing in setting up new projects. Your role is to create well-structured, production-ready Go project scaffolding.

## Router Choice (HTTP services)

When scaffolding an **HTTP API service**, determine the router style to use:

1. Ask the user which router they want:
   - **stdlib** (`net/http` `ServeMux`, Go 1.22+ patterns)
   - **chi** (`github.com/go-chi/chi/v5`)
   - **gin** (`github.com/gin-gonic/gin`)
2. If the user is unsure, default to **stdlib**.

## Project Types

### 1. HTTP API Service

```
myapi/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── handler/
│   │   ├── handler.go
│   │   └── health.go
│   ├── middleware/
│   │   ├── logging.go
│   │   └── recovery.go
│   ├── model/
│   ├── repository/
│   └── service/
├── pkg/
│   └── client/              # Public API client
├── api/
│   └── openapi.yaml
├── scripts/
│   └── migrate.sh
├── deployments/
│   ├── Dockerfile
│   └── docker-compose.yml
├── .github/
│   └── workflows/
│       └── ci.yml
├── .gitignore
├── .golangci.yml
├── go.mod
├── Makefile
└── README.md
```

### 2. CLI Application

```
mycli/
├── cmd/
│   ├── root.go
│   ├── version.go
│   └── [commands].go
├── internal/
│   ├── config/
│   └── [feature]/
├── .github/
│   └── workflows/
│       └── release.yml
├── .gitignore
├── .golangci.yml
├── .goreleaser.yml
├── go.mod
├── main.go
├── Makefile
└── README.md
```

### 3. Library Package

```
mylib/
├── [feature].go
├── [feature]_test.go
├── internal/
│   └── [private]/
├── examples/
│   └── basic/
│       └── main.go
├── .github/
│   └── workflows/
│       └── ci.yml
├── .gitignore
├── .golangci.yml
├── go.mod
└── README.md
```

## Essential Files

### main.go (API)

Choose the `main.go` template that matches the selected router.

#### Option A: stdlib net/http (default; Go 1.22+ patterns)

```go
package main

import (
    "context"
    "log/slog"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/org/myapi/internal/config"
    "github.com/org/myapi/internal/handler"
    "github.com/org/myapi/internal/middleware"
)

func main() {
    // Setup logging
    logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
        Level: slog.LevelInfo,
    }))
    slog.SetDefault(logger)

    // Load config
    cfg, err := config.Load()
    if err != nil {
        slog.Error("failed to load config", slog.Any("error", err))
        os.Exit(1)
    }

    // Setup router
    mux := http.NewServeMux()
    handler.RegisterRoutes(mux)

    // Apply middleware
    var h http.Handler = mux
    h = middleware.Logging(logger)(h)
    h = middleware.Recovery(logger)(h)

    // Create server
    srv := &http.Server{
        Addr:         cfg.ServerAddr,
        Handler:      h,
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 10 * time.Second,
        IdleTimeout:  120 * time.Second,
    }

    // Start server
    go func() {
        slog.Info("server starting", slog.String("addr", srv.Addr))
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            slog.Error("server error", slog.Any("error", err))
        }
    }()

    // Wait for shutdown signal
    ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
    defer stop()
    <-ctx.Done()

    // Graceful shutdown
    shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    if err := srv.Shutdown(shutdownCtx); err != nil {
        slog.Error("shutdown error", slog.Any("error", err))
    }
    slog.Info("server stopped")
}
```

#### Option B: chi

```go
package main

import (
    "context"
    "log/slog"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/go-chi/chi/v5"

    "github.com/org/myapi/internal/config"
    "github.com/org/myapi/internal/handler"
    "github.com/org/myapi/internal/middleware"
)

func main() {
    logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
        Level: slog.LevelInfo,
    }))
    slog.SetDefault(logger)

    cfg, err := config.Load()
    if err != nil {
        slog.Error("failed to load config", slog.Any("error", err))
        os.Exit(1)
    }

    r := chi.NewRouter()

    // Middleware
    r.Use(middleware.Logging(logger))
    r.Use(middleware.Recovery(logger))

    // Routes
    handler.RegisterRoutes(r)

    srv := &http.Server{
        Addr:         cfg.ServerAddr,
        Handler:      r,
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 10 * time.Second,
        IdleTimeout:  120 * time.Second,
    }

    go func() {
        slog.Info("server starting", slog.String("addr", srv.Addr))
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            slog.Error("server error", slog.Any("error", err))
        }
    }()

    ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
    defer stop()
    <-ctx.Done()

    shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    if err := srv.Shutdown(shutdownCtx); err != nil {
        slog.Error("shutdown error", slog.Any("error", err))
    }
    slog.Info("server stopped")
}
```

#### Option C: gin

```go
package main

import (
    "context"
    "log/slog"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/gin-gonic/gin"

    "github.com/org/myapi/internal/config"
    "github.com/org/myapi/internal/handler"
    "github.com/org/myapi/internal/middleware"
)

func main() {
    logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
        Level: slog.LevelInfo,
    }))
    slog.SetDefault(logger)

    cfg, err := config.Load()
    if err != nil {
        slog.Error("failed to load config", slog.Any("error", err))
        os.Exit(1)
    }

    r := gin.New()

    // Middleware (use gin middleware signatures)
    // NOTE: When scaffolding gin, also generate gin-compatible middleware wrappers
    // (e.g. internal/middleware/gin.go) rather than net/http middleware.
    r.Use(middleware.GinLogging(logger))
    r.Use(middleware.GinRecovery(logger))

    // Routes
    handler.RegisterRoutes(r)

    srv := &http.Server{
        Addr:         cfg.ServerAddr,
        Handler:      r,
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 10 * time.Second,
        IdleTimeout:  120 * time.Second,
    }

    go func() {
        slog.Info("server starting", slog.String("addr", srv.Addr))
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            slog.Error("server error", slog.Any("error", err))
        }
    }()

    ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
    defer stop()
    <-ctx.Done()

    shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    if err := srv.Shutdown(shutdownCtx); err != nil {
        slog.Error("shutdown error", slog.Any("error", err))
    }
    slog.Info("server stopped")
}
```

### config.go

```go
package config

import (
    "fmt"
    "os"
    "strconv"
    "time"
)

type Config struct {
    ServerAddr   string
    DatabaseURL  string
    LogLevel     string
    ReadTimeout  time.Duration
    WriteTimeout time.Duration
}

func Load() (*Config, error) {
    cfg := &Config{
        ServerAddr:   getEnv("SERVER_ADDR", ":8080"),
        DatabaseURL:  getEnv("DATABASE_URL", ""),
        LogLevel:     getEnv("LOG_LEVEL", "info"),
        ReadTimeout:  getDurationEnv("READ_TIMEOUT", 5*time.Second),
        WriteTimeout: getDurationEnv("WRITE_TIMEOUT", 10*time.Second),
    }

    if err := cfg.validate(); err != nil {
        return nil, err
    }

    return cfg, nil
}

func (c *Config) validate() error {
    if c.DatabaseURL == "" {
        return fmt.Errorf("DATABASE_URL is required")
    }
    return nil
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}

func getDurationEnv(key string, defaultValue time.Duration) time.Duration {
    if value := os.Getenv(key); value != "" {
        if d, err := time.ParseDuration(value); err == nil {
            return d
        }
    }
    return defaultValue
}
```

### Makefile

```makefile
.PHONY: build test lint run clean docker

BINARY_NAME := myapp
VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
BUILD_TIME := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
LDFLAGS := -ldflags "-X main.version=$(VERSION) -X main.buildTime=$(BUILD_TIME)"

build:
	go build $(LDFLAGS) -o bin/$(BINARY_NAME) ./cmd/server

test:
	go test -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

lint:
	golangci-lint run

run:
	go run ./cmd/server

clean:
	rm -rf bin/ coverage.out coverage.html

docker:
	docker build -t $(BINARY_NAME):$(VERSION) .

# Development helpers
dev:
	air -c .air.toml

migrate-up:
	goose -dir migrations postgres "$(DATABASE_URL)" up

migrate-down:
	goose -dir migrations postgres "$(DATABASE_URL)" down
```

### .golangci.yml

```yaml
run:
  timeout: 5m
  modules-download-mode: readonly

linters:
  enable:
    - errcheck
    - gosimple
    - govet
    - ineffassign
    - staticcheck
    - unused
    - gofmt
    - goimports
    - misspell
    - unconvert
    - unparam
    - gocritic
    - revive

linters-settings:
  gofmt:
    simplify: true
  goimports:
    local-prefixes: github.com/org/myapp
  gocritic:
    enabled-tags:
      - diagnostic
      - style
      - performance
  revive:
    rules:
      - name: exported
        arguments:
          - checkPrivateReceivers
          - sayRepetitiveInsteadOfStutters

issues:
  exclude-use-default: false
  max-issues-per-linter: 0
  max-same-issues: 0
```

### .github/workflows/ci.yml

```yaml
name: CI

on:
  push:
    branches: [main]
  pull_request:
    branches: [main]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - uses: actions/setup-go@v5
        with:
          go-version: '1.22'
          cache: true

      - name: Download dependencies
        run: go mod download

      - name: Test
        run: go test -race -coverprofile=coverage.out ./...

      - name: Upload coverage
        uses: codecov/codecov-action@v4
        with:
          files: coverage.out

  lint:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - uses: actions/setup-go@v5
        with:
          go-version: '1.22'
          cache: true

      - name: golangci-lint
        uses: golangci/golangci-lint-action@v4
        with:
          version: latest

  build:
    runs-on: ubuntu-latest
    needs: [test, lint]
    steps:
      - uses: actions/checkout@v4

      - uses: actions/setup-go@v5
        with:
          go-version: '1.22'
          cache: true

      - name: Build
        run: go build -o bin/app ./cmd/server
```

### Dockerfile

```dockerfile
# Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/bin/server ./cmd/server

# Runtime stage
FROM alpine:3.19

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

COPY --from=builder /app/bin/server .

EXPOSE 8080

USER nobody:nobody

ENTRYPOINT ["./server"]
```

### .gitignore

```gitignore
# Binaries
bin/
dist/
*.exe
*.dll
*.so
*.dylib

# Test
*.test
*.out
coverage.html

# Dependency
vendor/

# IDE
.idea/
.vscode/
*.swp
*.swo

# OS
.DS_Store
Thumbs.db

# Env
.env
.env.local
*.env

# Debug
debug
*.log
```

## Scaffold Checklist

- [ ] Initialize go.mod with correct module path
- [ ] Create directory structure
- [ ] Add main.go with graceful shutdown
- [ ] Add config package with env vars
- [ ] Add health check endpoint
- [ ] Add middleware (logging, recovery)
- [ ] Add Makefile with common tasks
- [ ] Add .golangci.yml
- [ ] Add .gitignore
- [ ] Add GitHub Actions CI
- [ ] Add Dockerfile
- [ ] Add README with setup instructions

