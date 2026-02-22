# Code Style and Conventions

## Language

- Go standard conventions (gofmt)
- Comments written in Japanese

## Structure

- Each module: cmd/main.go as entrypoint
- slog/ follows internal/ package pattern with handler, middleware, controller layers
- utils/ for shared utilities (logger, configuration)

## Naming

- Snake_case for file names (e.g. rw_mutex.go, group_by.go)
- Standard Go camelCase/PascalCase for symbols

## No linting/formatting config found

- No .golangci.yml or similar config files
- Assumed: standard go fmt and go vet
