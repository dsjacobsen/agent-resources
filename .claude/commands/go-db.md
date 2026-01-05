---
name: go-db
description: Design or implement database/sql + pgx repositories, queries, and transactions
---

Work on database access in this Go codebase. Ask for:

1. **Schema context** - table(s), columns, constraints (or migration snippets)
2. **Operations needed** - CRUD/list/search, filters, ordering, pagination
3. **Transaction needs** - what must be atomic
4. **Error expectations** - not found vs conflict vs validation

Use the `go-db-builder` agent to implement:
- Repository interfaces + implementations
- Parameterized SQL queries
- Transactional flows (`BeginTx`)
- Error mapping (e.g., unique violation → conflict)
- Tests (sqlmock or integration, matching existing style)


