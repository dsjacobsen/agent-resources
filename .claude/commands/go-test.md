---
name: go-test
description: Generate comprehensive tests for Go code
---

Generate tests for the Go code in the current context:

1. **Analyze the code** - Understand functions, methods, edge cases
2. **Create table-driven tests** - Group related test cases
3. **Cover edge cases** - Nil inputs, empty values, boundaries
4. **Test error paths** - Ensure errors are properly returned
5. **Add benchmarks** - For performance-critical code

Use the `go-test-generator` agent. Generate tests that:
- Use subtests with `t.Run()`
- Include descriptive test names
- Use `t.Helper()` for helpers
- Use `t.Cleanup()` for teardown
- Are deterministic (no flaky tests)

Place tests in `*_test.go` files alongside the source.

