# Contributing

Thanks for helping improve `operator-runbook-temporal-go`.

## Development Setup

Install Go 1.22 or newer, then verify the repository from a clean checkout:

```sh
go mod download
gofmt -w .
go vet ./...
go test ./...
```

The workflow tests use Temporal's in-memory test environment, so they do not
require a running Temporal server. Use `docs/local-temporal-quickstart.md` only
when you want to run the worker and starter commands against a local Temporal
development server.

## Contribution Scope

Keep pull requests small and focused. The current repository proves one durable
operator runbook workflow, so new behavior should preserve the workflow boundary:

- Keep workflow code deterministic.
- Put I/O, time, network calls, and human-system integrations behind
  activities.
- Add or update tests for behavior changes.
- Update README or `docs/` links when a workflow, command, or operating path
  changes.

## Pull Request Checklist

Before opening a pull request, run:

```sh
gofmt -w .
go vet ./...
go test ./...
```

Document any validation you could not run and why. Do not include secrets,
Temporal credentials, incident details, or private target names in issues, pull
requests, logs, screenshots, or test fixtures.
