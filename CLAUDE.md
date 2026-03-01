# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Overview

This is a personal Go learning repository. It contains code written while following multiple Go learning resources in parallel. Each top-level directory corresponds to a different source:

| Directory | Source |
|---|---|
| `introduction-to-programming-in-go/` | Go Bootcamp book — chapters are numbered subdirs (`1TheBasics`, `2Types`, `3CollectionTypes`, …) |
| `DocTutorials/ATourOfGo/` | Official [A Tour of Go](https://go.dev/tour/) — has its own `go.mod` |
| `DocTutorials/web-service-gin/` | Official Go + Gin REST API tutorial — has its own `go.mod` |
| `DocTutorials/Accessing_a_relational_database/` | Official Go + database/sql tutorial |
| `Tutorials/OrderBooksSystem/` | Practice project: CRUD with GORM + PostgreSQL |
| `Tutorials/CSVTutorial/` | CSV file handling tutorial |
| `LeetCodeExplore/` | LeetCode problem solutions (`ch1`–`ch8`) |
| `learning-go-book/` | "Learning Go" book chapters |
| `golang101Book/` | golang101 book chapters |
| `w3schools/` | W3Schools Go exercises |

## Code Style Notes

- The `introduction-to-programming-in-go/` files contain multiple `main()` functions and package-level snippets in a single file — this is intentional; the files are notes/examples, not runnable programs.
- Runnable projects (with `go.mod`) are in `DocTutorials/` and `Tutorials/`.

## Running Code

For modules with a `go.mod`, run from that module's directory:

```bash
cd <module-dir>
go run .          # run the module
go build .        # build
go test ./...     # run tests (if any)
```

For standalone snippet files (no `go.mod`), they are reference notes and are not meant to be run directly.

## Notable Projects

### OrderBooksSystem (`Tutorials/OrderBooksSystem/`)
- Uses GORM with PostgreSQL driver
- Models: `Book`, `Customer`, `Order`, `OrderDetail`
- Requires a running PostgreSQL instance; connection string is hardcoded in `main.go`

### ATourOfGo (`DocTutorials/ATourOfGo/`)
- Depends on `golang.org/x/tour`
- Contains both `main.go` (exercises) and `playground.go` (scratch space)
