---
name: go-check
description: Run Go linting, testing, and quality checks
---

Run comprehensive quality checks on the Go codebase:

```bash
# Format code
echo "=== Formatting ==="
gofmt -w .
goimports -w .

# Run linter
echo "=== Linting ==="
golangci-lint run ./...

# Run tests
echo "=== Tests ==="
go test -race -coverprofile=coverage.out ./...

# Show coverage
echo "=== Coverage ==="
go tool cover -func=coverage.out | tail -1

# Check for vulnerabilities
echo "=== Vulnerabilities ==="
govulncheck ./...
```

Report any issues found and suggest fixes for:
- Linter errors/warnings
- Test failures
- Low coverage areas
- Security vulnerabilities

If `golangci-lint` or `govulncheck` aren't installed, note this and skip those checks.

