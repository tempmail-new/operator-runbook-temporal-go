# Remediation Path Quickstart

This guide proves the workflow's `needs_operator` branch against a local
Temporal development server without editing source code.

## Validate The Workflow

The default test suite still runs without a Temporal server:

```sh
go test ./...
```

The tests include a real-activity workflow execution with
`simulate_unhealthy=true` so the local toggle cannot drift from the workflow
contract.

## Run The Healthy Path

Start Temporal:

```sh
temporal server start-dev
```

Start the worker:

```sh
go run ./cmd/runbook-worker
```

Start the default healthy workflow:

```sh
go run ./cmd/start-runbook -runbook-id demo -target checkout
```

The JSON result should include `"verdict": "healthy"` and no remediation entry.

## Run The Remediation Path

Keep the same Temporal server and worker running, then start a simulated
unhealthy workflow:

```sh
go run ./cmd/start-runbook -runbook-id demo-unhealthy -target checkout -simulate-unhealthy
```

The JSON result should include these fields:

```json
{
  "verdict": "needs_operator",
  "checks": [
    {
      "status": "unhealthy",
      "detail": "synthetic health check failed by request"
    }
  ],
  "remediations": [
    {
      "action": "page_operator"
    }
  ]
}
```

The flag is intentionally synthetic. It proves the orchestration and operator
handoff branch locally while real pager, ticketing, metrics, and deployment
integrations remain outside this repository's current scope.

## First Failure Checks

- If the result is still healthy, confirm the starter command includes
  `-simulate-unhealthy`.
- If the workflow stays queued, confirm the worker is running with the same
  `TEMPORAL_TASK_QUEUE` as the starter.
- If the starter cannot connect, confirm `temporal server start-dev` is still
  running and `TEMPORAL_ADDRESS` matches the server address.
