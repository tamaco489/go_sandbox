# Project Overview: go_sandbox

## Purpose

Personal sandbox repository for verifying the behavior of Go standard library and third-party libraries. Each subdirectory is an independent Go module.

## Tech Stack

- Language: Go
- Go versions: 1.23.5 (lo, sync), 1.24.2 (slog)
- External libraries: github.com/samber/lo (lo/), github.com/google/uuid (slog/)

## Modules

### lo/

- Demonstrates usage of the samber/lo utility library
- Key features: GroupBy, Map, Reduce, SliceToMap, UniqMap, FilterMap
- Models: User, Player, Pokemon, Achievement, Item

### sync/

- Demonstrates Go's sync package (stdlib only)
- Key features: Mutex, RWMutex, Once

### slog/

- HTTP server demonstrating Go 1.21+ structured logging (log/slog)
- Uses pure stdlib net/http (no frameworks)
- Routes: /api/v1/health, /api/v1/products/{id}, /api/v1/users/me, /api/v1/users/profile/me
- Middleware: auth, logging
- Logging levels by status: 5xx→error, 4xx→warn, 2xx→info
