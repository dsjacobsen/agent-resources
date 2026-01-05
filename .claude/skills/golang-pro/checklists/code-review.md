# Go Code Review Checklist

Use this checklist when reviewing Go code to ensure quality and idiomatic patterns.

## Correctness

- [ ] Code compiles without errors
- [ ] All tests pass (`go test ./...`)
- [ ] Race detector passes (`go test -race ./...`)
- [ ] Code handles edge cases appropriately
- [ ] Error conditions are properly handled

## Style & Formatting

- [ ] Code is formatted with `gofmt` or `goimports`
- [ ] Code passes `golangci-lint` (or your configured linter)
- [ ] Variable names are clear and follow Go conventions (camelCase)
- [ ] Package names are lowercase, single words
- [ ] Acronyms in names follow Go conventions (URL, HTTP, ID not Url, Http, Id)

## Documentation

- [ ] All exported items have doc comments
- [ ] Package has a package doc comment
- [ ] Doc comments start with the name of the thing being documented
- [ ] Complex logic has inline comments explaining "why"
- [ ] README updated if public API changed

## Error Handling

- [ ] Errors are checked, not ignored
- [ ] Errors are wrapped with context (`fmt.Errorf("doing X: %w", err)`)
- [ ] Sentinel errors are used appropriately
- [ ] Custom error types implement the `error` interface correctly
- [ ] `errors.Is()` and `errors.As()` are used for error checking

## Concurrency

- [ ] Goroutines have clear lifecycle management
- [ ] Contexts are propagated through the call chain
- [ ] Channels are closed by the sender
- [ ] Select statements have default or context.Done cases where appropriate
- [ ] Shared state is protected with appropriate synchronization
- [ ] No potential deadlocks or data races

## Performance

- [ ] Slices are pre-allocated where capacity is known
- [ ] Strings are built with `strings.Builder` for concatenation
- [ ] Heavy allocations are avoided in hot paths
- [ ] `sync.Pool` is used for frequently allocated objects
- [ ] Benchmarks exist for performance-critical code

## Testing

- [ ] Table-driven tests with subtests are used
- [ ] Test cases cover happy path and error paths
- [ ] Tests are deterministic (no flaky tests)
- [ ] Mocks/fakes are used instead of real dependencies
- [ ] Test helper functions are marked with `t.Helper()`
- [ ] `t.Cleanup()` is used for resource cleanup

## Security

- [ ] User input is validated and sanitized
- [ ] SQL queries use parameterized queries (no string interpolation)
- [ ] Sensitive data is not logged
- [ ] File paths are validated to prevent traversal attacks
- [ ] HTTP handlers have appropriate timeouts
- [ ] HTTPS/TLS is used for external connections

## API Design

- [ ] Functions accept interfaces, return structs
- [ ] Interfaces are small and focused (1-3 methods)
- [ ] Functional options pattern is used for optional configuration
- [ ] Breaking changes are avoided or clearly communicated
- [ ] Public API is minimal and necessary

## Dependencies

- [ ] No unnecessary dependencies added
- [ ] Dependencies are at appropriate versions
- [ ] `go.sum` is committed
- [ ] Vendor directory is updated if using vendoring

