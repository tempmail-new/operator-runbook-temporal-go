# Local Temporal Quickstart

This guide runs the bootstrap workflow against Temporal's local development
server.

## Prerequisites

- Go 1.22 or newer.
- Temporal CLI.

Install the Temporal CLI from Temporal's official installation guide, then
confirm it is available:

```sh
temporal version
```

## Validate Without A Server

The workflow tests use Temporal's in-memory test environment and mocked activity
boundaries, so they do not need a running Temporal server:

```sh
go test ./...
```

## Run Locally

Start Temporal:

```sh
temporal server start-dev
```

Start the worker:

```sh
go run ./cmd/runbook-worker
```

Start one workflow execution:

```sh
go run ./cmd/start-runbook -runbook-id demo -target checkout
```

The starter prints a JSON result with the runbook ID, target, verdict, and
recorded checks.

To prove the operator remediation branch without editing source, run the
[remediation path quickstart](remediation-path-quickstart.md) with the starter's
`-simulate-unhealthy` flag.

## Configuration

Both commands use these environment variables:

| Variable | Default | Purpose |
| --- | --- | --- |
| `TEMPORAL_ADDRESS` | `localhost:7233` | Temporal frontend address. |
| `TEMPORAL_TASK_QUEUE` | `operator-runbook` | Task queue for worker and starter. |

Use the same `TEMPORAL_TASK_QUEUE` value for both commands.

## First Failure Checks

- If the starter cannot connect, confirm `temporal server start-dev` is still
  running and `TEMPORAL_ADDRESS` matches the server address.
- If the workflow stays queued, confirm `cmd/runbook-worker` is running with the
  same `TEMPORAL_TASK_QUEUE`.
- If tests fail before a workflow starts, run `gofmt -w .` and `go test ./...`
  again so formatting and generated dependency metadata are current.

For the full recovery checklist covering Temporal startup, worker registration,
and starter verdicts, see the
[local Temporal troubleshooting guide](local-temporal-troubleshooting.md).
