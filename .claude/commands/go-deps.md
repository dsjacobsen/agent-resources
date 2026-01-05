---
name: go-deps
description: Analyze and manage Go dependencies
---

Analyze the Go module dependencies:

```bash
# Show all dependencies
go list -m all

# Show why a dependency is needed
go mod why -m <module>

# Find outdated dependencies
go list -u -m all

# Check for vulnerabilities
govulncheck ./...

# Tidy dependencies
go mod tidy

# Verify checksums
go mod verify
```

**Report on:**

1. **Direct dependencies** - What's in go.mod
2. **Outdated packages** - What can be updated
3. **Vulnerabilities** - Security issues found
4. **Unused imports** - Dependencies that can be removed
5. **Major version updates** - Breaking changes available

**Recommend:**
- Safe updates (patch/minor versions)
- Updates requiring code changes
- Dependencies to consider replacing
- Security patches to apply immediately

**For updates:**
```bash
# Update specific package
go get package@version

# Update all to latest minor/patch
go get -u ./...

# Update all to latest (including major)
go get -u=patch ./...
```

