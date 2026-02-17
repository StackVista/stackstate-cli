# AGENTS.md

This file provides guidance to OpenCode agents when working with code in this repository.

## Important Product Naming

**CRITICAL**: Always use "SUSE Observability" when referring to the product. "StackState" is deprecated and should not be used in new code, documentation, or communications.

## Project Overview

SUSE Observability CLI (`sts`) - A command-line tool for managing topology-powered observability within the SUSE Observability platform. Built with Go and Cobra framework.

## Development Environment

This project uses Nix Flakes for development environment management:

```bash
# Enter development shell (requires Nix with flakes enabled)
nix develop

# Build the CLI
nix build

# Or use Go directly
go build -o sts main.go
```

## Common Commands

```bash
# Run CLI locally
go run main.go

# Run all tests
go test -v ./...

# Run tests for a specific package
go test -v ./cmd/monitor/...

# Run a single test
go test -v ./cmd/monitor/... -run TestMonitorApply

# Lint
golangci-lint run

# Regenerate OpenAPI clients (when openapi_version changes)
nix develop -c ./scripts/generate_stackstate_api.sh
```

## Architecture

### Command Structure Pattern

All commands follow noun/verb/flags structure:
```
sts [noun] [verb] [flags]
```
Example: `sts context save --api-url URL --api-token TOKEN`

### Key Components

- **Entry point**: `main.go` - Sets up DI container, logging, error handling
- **Commands**: `cmd/` - Each subdirectory is a command group (monitor/, settings/, etc.)
- **DI Container**: `internal/di/Deps` - Dependency injection for testability
- **API Clients**: `generated/stackstate_api/` and `generated/stackstate_admin_api/` - Auto-generated from OpenAPI
- **Output**: `internal/printer/Printer` - All output must go through this interface
- **Errors**: `internal/common/CLIError` - Custom error type with exit codes and usage hints

### Code Generation

OpenAPI clients are generated from external specs:
- Version pinned in `stackstate_openapi/openapi_version`
- Custom templates in `stackstate_openapi/template/`
- Generated code in `generated/` - do not edit manually

## Command Development Rules

From CMD_DEVELOPMENT.md - critical rules for consistency:

### Output
- Always use `Printer` for stdout/stderr output
- Every command must support both human-readable and JSON output (`-o json`)
- Use logger for debug/verbose output, not Printer

### Errors
- Return `CLIError` from commands, don't print errors directly
- Set `ShowUsage=true` for input/parsing errors
- Set `ShowUsage=false` for API/technical errors

### Help Text Requirements
- `Use:` - Skeleton usage. Show required flags directly, mutually exclusive flags in `{}` with `|`
- `Short:` - One sentence, no period, starts with verb
- `Long:` - Full description with period
- `Example:` - Copy-pasteable examples with `#` comments

### Naming
- Nouns plural if pertaining to a set of entities
- Reuse verbs: `apply`, `list`, `describe`, `export`
- Flag shorthands must be consistent across all commands (e.g., `-f` always means `--file`)

## CLI Help Documentation Specialist

**CRITICAL AGENT**: This agent manages cobra.Command help text according to ADR 001 and ADR 002.

### Policy Overview
Per ADR 001: CLI `--help` output is the **authoritative source** for command documentation. External docs provide overviews only. This means help text must be self-sufficient, consistent, and include practical examples.

### Use Field Rules
- **IMPORTANT**: Do NOT include flags in the `Use` field. Required flags and mutually exclusive flag groups are **auto-generated** by `AddRequiredFlagsToUseString()` in `main.go` based on flag annotations set by `MarkMutexFlags()`. Including flags manually causes duplication in the help output.
- **Noun-level commands**: `Use: "context"` (command name only)
- **Noun-verb commands**: `Use: "delete"` (verb name only, no flags)
- **Positional arguments only**: If the command takes positional arguments, include them: `Use: "describe [NAME]"`

### Short Field Rules
- Start with **imperative verb** (Save, List, Delete, Show, Set, Validate)
- Describe what the command does, not how
- Use articles ("a", "the") appropriately
- Maximum one sentence
- **No period** at the end
- Must add context beyond just command name

Examples:
- ✅ `"Save a connection context to the CLI configuration"`
- ❌ `"Save a context"`
- ✅ `"List all configured connection contexts"`
- ❌ `"List contexts"`

### Long Field Rules
- Start with upper case, end with period
- Can be multiple sentences
- **Must add value** beyond Short - explain what the entity is, when to use, or outcome
- Avoid excessive abbreviations
- Provide context about the domain concept

Examples:
- ✅ `"Save a connection context to the CLI configuration file. A context stores the URL and authentication credentials for a SUSE Observability server, allowing you to switch between different servers or environments."`
- ❌ `"Save a context."`

### Example Field Rules
- Use backtick-quoted multiline strings (not string concatenation)
- Each example starts with `#` comment describing what it does
- Comments: lowercase start, no period, imperative verb
- Provide 2-3 examples covering common scenarios
- Examples must be realistic and runnable
- Cover different flag combinations when applicable

Format:
```go
Example: `# save a context with an API token
sts context save --name production --url https://suse-observability.example.com --api-token xxxx-xxxx

# save a context with a service token for CI/CD
sts context save --name ci-pipeline --url https://suse-observability.example.com --service-token xxxx-xxxx`
```

### Flag Description Rules
- Start with upper case, no period at the end
- Describe what the flag does, not implementation details
- Include default value mention if non-obvious

### Quality Checklist
For every cobra.Command:
- [ ] `Use` contains only the command name and positional arguments (no flags)
- [ ] `Short` starts with verb, no period, adds context beyond command name
- [ ] `Long` adds value beyond `Short`, ends with period, explains domain concepts
- [ ] `Example` has 2-3 realistic, copy-pasteable examples with lowercase comments
- [ ] Flag descriptions are clear and properly capitalized
- [ ] Uses "SUSE Observability" (not "StackState") in descriptions
- [ ] No string concatenation - uses backtick multiline strings

### When to Invoke This Agent
- Creating new commands
- Modifying existing command help text
- Code reviews involving cobra.Command changes
- Ensuring consistency across all CLI help documentation

## Testing Patterns

Commands are tested using mock DI container:

```go
func TestCommandName(t *testing.T) {
    cli := di.NewMockDeps(t)
    // Setup mocks
    // Execute command
    // Assert results
}
```

Test utilities in `internal/di/mock_deps.go` and `internal/di/command_test_util.go`.

## CI/CD

- Linting: golangci-lint v2.3.1 (see `.golangci.yml`)
- Tests run on every PR
- Releases via GoReleaser on version tags (`v*`)
