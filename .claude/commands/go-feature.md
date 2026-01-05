---
name: go-feature
description: Implement a new feature in the Go codebase
---

Implement a new feature. Ask for:

1. **Feature description** - What should it do?
2. **User stories** - Who uses it and how?
3. **Acceptance criteria** - How do we know it's done?

Use the `go-feature` agent to:

**Phase 1: Discovery**
- Understand existing codebase patterns
- Identify which packages need changes
- Note existing conventions

**Phase 2: Design**
- List components needed (models, services, handlers)
- Identify dependencies
- Plan the implementation order

**Phase 3: Implementation**
- Implement domain model first
- Add repository layer
- Add service layer with business logic
- Add handler/API layer
- Wire up dependencies

**Phase 4: Testing**
- Unit tests for each layer
- Integration tests if needed
- Manual testing

**Follow existing patterns for:**
- Error handling
- Logging
- Configuration
- Naming conventions

