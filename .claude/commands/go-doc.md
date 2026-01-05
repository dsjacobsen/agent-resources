---
name: go-doc
description: Generate documentation for Go code
---

Generate documentation for the Go code in context:

Use the `go-docs-generator` agent to create:

**For packages without doc.go:**
- Create `doc.go` with package overview
- Include usage examples
- Document key types and functions

**For undocumented exports:**
- Add doc comments to all exported items
- Start comments with the item name
- Include examples where helpful

**For README:**
- Project overview
- Installation instructions
- Quick start guide
- Configuration options
- API documentation (if applicable)

**Documentation standards:**
- Use complete sentences
- Start with the name being documented
- Include code examples
- Document parameters, return values, and errors
- Add `Example*` test functions for godoc

