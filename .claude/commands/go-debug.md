---
name: go-debug
description: Debug a Go issue - analyze errors, traces, and behavior
---

Debug the issue described. Gather information:

1. **Error message** - Exact error text
2. **Stack trace** - If available
3. **Expected vs actual** - What should happen vs what happens
4. **Reproduction steps** - How to trigger the issue

**Debugging approach:**

```bash
# Run with race detector
go test -race ./...

# Run with verbose output
go test -v ./path/to/package

# Check for common issues
go vet ./...

# Build with debug info
go build -gcflags="all=-N -l" ./...
```

**Analysis:**

1. **Read the error** - What does it actually say?
2. **Find the source** - Where in code does it originate?
3. **Trace the flow** - How did we get there?
4. **Identify the cause** - Why is it happening?
5. **Propose fix** - How to resolve it?

**Common Go issues to check:**
- Nil pointer dereference
- Slice out of bounds
- Map access on nil map
- Channel deadlock
- Race conditions
- Context cancellation
- Error not checked

Provide the fix with explanation of root cause.

