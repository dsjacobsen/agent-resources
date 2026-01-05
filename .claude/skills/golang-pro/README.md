# Golang Pro Skill

Expert Go developer skill for building efficient, concurrent, and scalable systems with Go 1.21+.

## Installation

```bash
# Using uvx (after pushing to GitHub)
uvx add-skill dsjacobsen/golang-pro

# Or manually copy to your Claude skills directory
cp -r . ~/.claude/skills/golang-pro/
```

## What's Included

```
golang-pro/
├── SKILL.md              # Main skill definition
├── README.md             # This file
├── examples/
│   ├── http-service.go   # Production HTTP server example
│   └── worker-pool.go    # Concurrency patterns example
└── checklists/
    ├── code-review.md    # Go code review checklist
    └── new-project.md    # New project setup checklist
```

## What This Skill Teaches Claude

### Idiomatic Go Patterns
- Effective Go guidelines
- Functional options pattern
- Interface design (accept interfaces, return structs)
- Error handling with wrapping

### Concurrency
- Worker pool implementations
- Fan-out/fan-in pipelines
- Context propagation
- Goroutine lifecycle management

### Production Patterns
- HTTP server with graceful shutdown
- Middleware (logging, recovery)
- Clean architecture layers
- Structured logging with `log/slog`

### Modern Go (1.21+)
- Generics
- Structured logging
- New HTTP routing (Go 1.22+)
- Multi-error handling

### Testing
- Table-driven tests
- Benchmarking
- Test fixtures and helpers

## Usage

The skill auto-activates when you're working with Go code. Just ask:

- "Create a Go HTTP server with graceful shutdown"
- "Implement a worker pool pattern"
- "Help me handle errors idiomatically"
- "Review this Go code"

## Examples

See the `examples/` directory for reference implementations:

- **`http-service.go`** - Complete HTTP service with clean architecture
- **`worker-pool.go`** - Concurrency patterns (worker pool, fan-out/fan-in)

## Checklists

- **`code-review.md`** - Use when reviewing Go code
- **`new-project.md`** - Use when starting a new Go project

## Resources

- [Effective Go](https://go.dev/doc/effective_go)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md)
