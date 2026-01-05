# Go Development Toolkit for Claude Code

A comprehensive collection of Claude Code skills, agents, and commands for Go development.

## 🚀 Quick Start

```bash
# Clone the repo
git clone https://github.com/dsjacobsen/agent-resources.git
cd agent-resources

# Copy to global .claude (available in all projects)
cp -r .claude/* ~/.claude/

# Or copy to local project (available in current project only)
cp -r .claude /path/to/your/project/
```

### Install Individual Components

```bash
# Skill
uvx add-skill dsjacobsen/golang-pro

# Agents
uvx add-agent dsjacobsen/go-scaffold
uvx add-agent dsjacobsen/go-feature
uvx add-agent dsjacobsen/go-api-builder
uvx add-agent dsjacobsen/go-cli-builder
uvx add-agent dsjacobsen/go-reviewer
uvx add-agent dsjacobsen/go-test-generator
uvx add-agent dsjacobsen/go-docs-generator
uvx add-agent dsjacobsen/go-refactor
uvx add-agent dsjacobsen/go-db-builder

# Commands
uvx add-command dsjacobsen/go-new
uvx add-command dsjacobsen/go-feature
uvx add-command dsjacobsen/go-endpoint
uvx add-command dsjacobsen/go-test
uvx add-command dsjacobsen/go-review
uvx add-command dsjacobsen/go-check
uvx add-command dsjacobsen/go-doc
uvx add-command dsjacobsen/go-refactor
uvx add-command dsjacobsen/go-debug
uvx add-command dsjacobsen/go-deps
uvx add-command dsjacobsen/go-db
```

## 📁 Structure

```
.claude/
├── skills/
│   └── golang-pro/           # Go expertise & patterns
│
├── agents/                    # Autonomous specialists
│   ├── go-scaffold.md        # New project setup
│   ├── go-feature.md         # Feature implementation
│   ├── go-api-builder.md     # REST API endpoints
│   ├── go-cli-builder.md     # CLI applications
│   ├── go-reviewer.md        # Code review
│   ├── go-test-generator.md  # Test generation
│   ├── go-docs-generator.md  # Documentation
│   ├── go-refactor.md        # Code improvement
│   └── go-db-builder.md      # database/sql + pgx repositories & queries
│
└── commands/                  # Quick shortcuts
    ├── go-new.md             # /go-new
    ├── go-feature.md         # /go-feature
    ├── go-endpoint.md        # /go-endpoint
    ├── go-test.md            # /go-test
    ├── go-review.md          # /go-review
    ├── go-check.md           # /go-check
    ├── go-doc.md             # /go-doc
    ├── go-refactor.md        # /go-refactor
    ├── go-debug.md           # /go-debug
    ├── go-deps.md            # /go-deps
    └── go-db.md              # /go-db
```

## 🧩 Skills vs Agents vs Commands

| Type | Purpose | How to Use |
|------|---------|------------|
| **Skills** | Teach Claude patterns & knowledge | Auto-applies when relevant |
| **Agents** | Do complex tasks autonomously | `@go-reviewer review this code` |
| **Commands** | Quick shortcuts | `/go-check` |

## 📦 What's Included

### Skill: `golang-pro`

Expert Go knowledge covering:
- Idiomatic Go patterns (Effective Go)
- Concurrency (goroutines, channels, sync)
- Error handling best practices
- Testing patterns (table-driven, benchmarks)
- Modern Go 1.21+ features (generics, slog)
- Production patterns (graceful shutdown, middleware)

### Agents

#### Feature Implementation
| Agent | Description |
|-------|-------------|
| `go-scaffold` | Creates new Go projects with proper structure, CI/CD, Docker |
| `go-feature` | Implements features following clean architecture patterns |
| `go-api-builder` | Builds REST endpoints with handlers, services, DTOs |
| `go-cli-builder` | Creates CLI apps with Cobra or stdlib flags |

#### Code Quality
| Agent | Description |
|-------|-------------|
| `go-reviewer` | Reviews code for correctness, security, performance |
| `go-test-generator` | Generates comprehensive test suites |
| `go-docs-generator` | Creates package docs, README, godoc comments |
| `go-refactor` | Improves code structure while preserving behavior |
| `go-db-builder` | Builds `database/sql` + pgx repositories, queries, and transactions |

### Commands

| Command | Description |
|---------|-------------|
| `/go-new` | Create a new Go project |
| `/go-feature` | Implement a new feature |
| `/go-endpoint` | Add a REST API endpoint |
| `/go-test` | Generate tests for code |
| `/go-review` | Review code for issues |
| `/go-check` | Run lint, tests, coverage |
| `/go-doc` | Generate documentation |
| `/go-refactor` | Improve code structure |
| `/go-debug` | Debug an issue |
| `/go-deps` | Analyze dependencies |
| `/go-db` | Implement repositories/queries/transactions (`database/sql` + pgx) |

## 🔄 Example Workflow

```
/go-new                    # Create project
/go-endpoint               # Add API endpoints
/go-feature                # Implement business logic
/go-test                   # Generate tests
/go-check                  # Run quality checks
/go-review                 # Review the code
/go-doc                    # Generate docs
```

## 📚 Resources

- [Claude Code Skills Documentation](https://code.claude.com/docs/en/skills)
- [Claude Code Subagents Documentation](https://code.claude.com/docs/en/subagents)
- [Effective Go](https://go.dev/doc/effective_go)
- [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md)

## 📄 License

MIT

