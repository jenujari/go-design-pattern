# Project Context

## Purpose
The `go-design-pattern` repository demonstrates classic software design patterns implemented in Go. Each pattern is placed in its own package directory (e.g., `Adapter/`, `Observer/`, `Singleton/`), with a self‑contained `main.go` that exercises the pattern and prints observable behavior. The goal is educational: a concise, runnable example that can be used as a teaching aid or a quick reference.

## Tech Stack
- **Go 1.19**
- Standard library only – no third‑party dependencies.
- `go.mod` defines the module path `github.com/openspec/go-design-pattern`.
- Build commands: `go build ./...` and `go run ./path/to/main.go`.

## Project Conventions

### Code Style
- Follow the official `gofmt` and `go vet` defaults.
- Use `//` for single‑line comments and block comments for file headers.
- Naming: exported identifiers use UpperCamel; local variables use lowerCamel.
- Structs and interfaces use UpperCamel.
- Constants and variables following usual Go conventions.
- Exports: only the `main` function and any public types required to interact with the pattern (e.g., `Observer` and `Subject` in `Observer/main.go`).
- Files are named `main.go` inside each pattern directory to keep the example minimal.

### Architecture Patterns
- Each folder implements a single Go design pattern.
- No package level state; all code lives in `main.go` for simplicity.
- If a pattern introduces multiple types, they reside in the same file for readability.

### Testing Strategy
- No unit tests are included yet.
- The repository relies on manual `go run` execution as a quick sanity check.
- Future iterations could add small tests in `*_test.go` files using the `testing` package.

### Git Workflow
- Branch naming: feature/… or bugfix/….
- Commit messages start with a capitalized, short summary (e.g., "Add Observer example").
- Pull requests require the editor to review and run `go run` to confirm build.

## Domain Context
Domain issues are minimal – the code exercises design patterns; domain knowledge is limited to Go language idioms.

## Important Constraints
- Run on Go 1.19 or newer due to module syntax.
- No network dependencies; all logic is local.

## External Dependencies
None – the repository uses only the Go standard library.
