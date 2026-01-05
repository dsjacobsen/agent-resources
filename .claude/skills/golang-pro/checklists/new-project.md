# Go New Project Checklist

Use this checklist when starting a new Go project to ensure proper setup.

## Initial Setup

- [ ] Initialize module: `go mod init github.com/username/projectname`
- [ ] Create standard directory structure (see below)
- [ ] Add `.gitignore` for Go projects
- [ ] Setup linter configuration (`.golangci.yml`)
- [ ] Create initial README.md
- [ ] Add LICENSE file

## Directory Structure

```
project/
├── cmd/
│   └── appname/
│       └── main.go
├── internal/
│   ├── config/
│   ├── handler/
│   ├── service/
│   └── repository/
├── pkg/                    # (optional) public packages
├── api/                    # API specifications (OpenAPI, proto)
├── scripts/
├── deployments/
├── .github/
│   └── workflows/
├── .gitignore
├── .golangci.yml
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

## Essential Files

### .gitignore

```gitignore
# Binaries
*.exe
*.exe~
*.dll
*.so
*.dylib
bin/
dist/

# Test binary
*.test

# Output of go coverage tool
*.out
coverage.html

# Dependency directories
vendor/

# IDE
.idea/
.vscode/
*.swp
*.swo

# OS
.DS_Store
Thumbs.db

# Environment
.env
.env.local
```

### Makefile

```makefile
.PHONY: build test lint clean run

BINARY_NAME=appname
VERSION?=$(shell git describe --tags --always --dirty)

build:
	go build -ldflags "-X main.version=$(VERSION)" -o bin/$(BINARY_NAME) ./cmd/$(BINARY_NAME)

test:
	go test -race -coverprofile=coverage.out ./...

lint:
	golangci-lint run

clean:
	rm -rf bin/ coverage.out

run:
	go run ./cmd/$(BINARY_NAME)
```

### .golangci.yml

```yaml
run:
  timeout: 5m

linters:
  enable:
    - errcheck
    - gosimple
    - govet
    - ineffassign
    - staticcheck
    - typecheck
    - unused
    - gofmt
    - goimports
    - misspell
    - unconvert
    - unparam

linters-settings:
  gofmt:
    simplify: true
  goimports:
    local-prefixes: github.com/username/projectname
```

## CI/CD Setup

### GitHub Actions (.github/workflows/ci.yml)

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
      - name: Test
        run: go test -race -coverprofile=coverage.out ./...
      - name: Upload coverage
        uses: codecov/codecov-action@v4

  lint:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with:
          go-version: '1.22'
      - uses: golangci/golangci-lint-action@v4
```

## Configuration Management

- [ ] Use environment variables for runtime config
- [ ] Support configuration files (YAML/TOML)
- [ ] Validate configuration at startup
- [ ] Document all configuration options
- [ ] Provide sensible defaults

## Logging Setup

- [ ] Use `log/slog` for structured logging
- [ ] Configure log level from environment
- [ ] Include request IDs in logs
- [ ] Don't log sensitive information

## Health Checks

- [ ] Implement `/health` or `/healthz` endpoint
- [ ] Include readiness and liveness probes for K8s
- [ ] Check critical dependencies in health check

## Documentation

- [ ] README with project overview
- [ ] Installation instructions
- [ ] Configuration documentation
- [ ] API documentation (if applicable)
- [ ] Contributing guidelines

