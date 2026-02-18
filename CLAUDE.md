# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
# Run the server (default port :4000)
go run cmd/web/*

# Run with custom flags
go run cmd/web/* -addr=":9999" -dsn="web:pass@tcp(localhost:3306)/snippetbox?parseTime=true"

# Build
go build -o bin/web cmd/web/*

# Download dependencies
go mod download
```

## Architecture

This is a learning project following a Go web development tutorial. It uses the standard library only (plus a MySQL driver).

**Entry point:** `cmd/web/main.go` — parses CLI flags (`-addr`, `-dsn`), opens the DB connection, wires up dependencies, and starts the server.

**Dependency injection pattern:** All handlers are methods on an `application` struct defined in `main.go`:
```go
type application struct {
    errorLog *log.Logger
    infoLog  *log.Logger
}
```
New shared dependencies (DB models, template cache, session manager, etc.) should be added as fields on this struct.

**Key files in `cmd/web/`:**
- `routes.go` — All route registrations via `app.routes()`
- `handlers.go` — HTTP handler methods on `application`
- `helpers.go` — Centralized error helpers: `serverError()` (500 + stack trace), `clientError()`, `notFound()`

**Templates:** HTML templates live in `ui/html/`. The naming convention is `*.layout.html`, `*.page.html`, and `*.partial.html`. Templates are parsed at request time (not cached yet); a future step will add a template cache.

**Static assets:** Served from `ui/static/` at the `/static/` URL prefix.

## Database

MySQL is required. The default DSN assumes a local MySQL instance with:
- User: `web`, Password: `pass`
- Database: `snippetbox`
- `parseTime=true` is required

The `openDB()` helper in `main.go` opens and pings the DB on startup. Database models will live in `pkg/models/` (not yet implemented).

## Project Status

This is an early-stage tutorial project. Tests, a template cache, session management, and database models have not been implemented yet.
