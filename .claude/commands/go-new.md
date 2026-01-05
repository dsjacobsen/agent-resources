---
name: go-new
description: Create a new Go project with proper structure
---

Create a new Go project. Ask for:

1. **Project name** - Module name (e.g., github.com/user/project)
2. **Project type** - API service, CLI tool, or library
3. **Features needed** - Database, auth, config, etc.

Use the `go-scaffold` agent to create:

**For API Service:**
- `cmd/server/main.go` with graceful shutdown
- `internal/` with handler, service, repository layers
- Health check endpoint
- Middleware (logging, recovery)
- Config from environment variables

**For CLI Tool:**
- Cobra-based command structure
- Config file support
- Multiple output formats

**For Library:**
- Clean public API
- Examples directory
- Comprehensive tests

**Always include:**
- `go.mod` with correct module path
- `Makefile` with build, test, lint targets
- `.golangci.yml` linter config
- `.gitignore` for Go projects
- `README.md` with setup instructions
- `.github/workflows/ci.yml` for GitHub Actions

