---
name: go-refactor
description: Refactor Go code to improve structure and readability
---

Refactor the Go code in context to improve quality:

Use the `go-refactor` agent. Before refactoring:

1. **Verify tests exist** - Don't refactor without tests
2. **Run tests** - Ensure passing baseline

**Common refactorings:**

- **Extract function** - Break up long functions
- **Extract interface** - Create abstractions for testing
- **Simplify conditionals** - Use early returns
- **Remove duplication** - DRY principle
- **Rename for clarity** - Better names
- **Reorganize packages** - Better structure

**Process:**
1. Make one small change
2. Run tests
3. Commit
4. Repeat

**Rules:**
- Preserve existing behavior
- Keep changes focused
- Run `go test -race ./...` after each change
- Run `golangci-lint run` to verify style

