---
name: go-review
description: Review Go code for issues, patterns, and best practices
---

Review the Go code in the current context for:

1. **Correctness** - Logic errors, edge cases, error handling
2. **Idiomatic Go** - Following Effective Go guidelines
3. **Concurrency** - Race conditions, proper goroutine management
4. **Performance** - Unnecessary allocations, inefficient patterns
5. **Security** - Input validation, SQL injection, data exposure

Use the `go-reviewer` agent for this task. Provide a structured review with:
- Critical issues (must fix)
- Major issues (should fix)  
- Minor issues (consider fixing)
- What's done well

Focus on actionable feedback with specific file:line references.

