---
name: go-endpoint
description: Add a new REST API endpoint with full implementation
---

Add a new API endpoint. Ask for:

1. **Resource name** - What entity (e.g., users, products, orders)
2. **Operations** - CRUD or specific operations needed
3. **Fields** - What data fields to include

Use the `go-api-builder` agent to create:

**All layers:**
- `internal/model/{resource}.go` - Domain model
- `internal/dto/{resource}_dto.go` - Request/response types with validation
- `internal/repository/{resource}_repository.go` - Data access interface
- `internal/service/{resource}_service.go` - Business logic
- `internal/handler/{resource}_handler.go` - HTTP handlers

**For each operation:**
- `GET /{resources}` - List with pagination
- `GET /{resources}/{id}` - Get single
- `POST /{resources}` - Create
- `PUT /{resources}/{id}` - Update
- `DELETE /{resources}/{id}` - Delete

**Include:**
- Input validation
- Proper error handling with status codes
- Structured logging
- Route registration

