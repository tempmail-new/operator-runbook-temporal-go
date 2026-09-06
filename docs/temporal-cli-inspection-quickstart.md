# Temporal CLI Inspection Quickstart

This guide shows the shortest headless path to inspect healthy and remediation
workflow executions with the Temporal CLI.

## Prerequisites

- Go 1.22 or newer.
- Temporal CLI.
- A shell that can keep the Temporal server, worker, and starter commands in
  separate terminals.

Confirm the Temporal CLI is available:

```sh
temporal version
```

## Start The Local Proof

Start Temporal and leave it running:

```sh
temporal server start-dev
```

The development server exposes the frontend API at `localhost:7233` in the
`default` namespace.

Start the worker in a second terminal:

```sh
go run ./cmd/runbook-worker
```

Wait for the worker to print `runbook worker started`.

## Inspect A Healthy Execution

Start the default healthy workflow in a third terminal:

```sh
go run ./cmd/start-runbook -runbook-id demo -target checkout
```

List the newest local runbook executions from the CLI:

```sh
temporal workflow list \
  --address localhost:7233 \
  --namespace default \
  --query 'WorkflowId STARTS_WITH "operator-runbook-checkout-"' \
  --limit 5
```

Copy the newest workflow ID that begins with `operator-runbook-checkout-`, then
inspect its metadata, history, and result:

```sh
WORKFLOW_ID=operator-runbook-checkout-YYYYMMDDTHHMMSSZ

temporal workflow describe \
  --address localhost:7233 \
  --namespace default \
  --workflow-id "$WORKFLOW_ID"

temporal workflow show \
  --address localhost:7233 \
  --namespace default \
  --workflow-id "$WORKFLOW_ID"

temporal workflow result \
  --address localhost:7233 \
  --namespace default \
  --workflow-id "$WORKFLOW_ID"
```

Check these CLI signals:

- `describe` reports workflow type `OperatorRunbookWorkflow`.
- `describe` reports task queue `operator-runbook` unless
  `TEMPORAL_TASK_QUEUE` was overridden.
- `describe` reports a completed execution status.
- `show` includes a `RunHealthCheck` activity event.
- `result` includes the `healthy` verdict.

## Inspect A Remediation Execution

Keep the same Temporal server and worker running, then start the synthetic
unhealthy path:

```sh
go run ./cmd/start-runbook -runbook-id demo-unhealthy -target checkout -simulate-unhealthy
```

List the latest matching executions again and copy the newest workflow ID:

```sh
temporal workflow list \
  --address localhost:7233 \
  --namespace default \
  --query 'WorkflowId STARTS_WITH "operator-runbook-checkout-"' \
  --limit 5
```

Inspect the remediation run:

```sh
WORKFLOW_ID=operator-runbook-checkout-YYYYMMDDTHHMMSSZ

temporal workflow describe \
  --address localhost:7233 \
  --namespace default \
  --workflow-id "$WORKFLOW_ID"

temporal workflow show \
  --address localhost:7233 \
  --namespace default \
  --workflow-id "$WORKFLOW_ID"

temporal workflow result \
  --address localhost:7233 \
  --namespace default \
  --workflow-id "$WORKFLOW_ID"
```

Check these CLI signals:

- `describe` still reports workflow type `OperatorRunbookWorkflow`.
- `show` includes both `RunHealthCheck` and `RequestHumanRemediation` activity
  events.
- `show` includes the `simulate_unhealthy` workflow input.
- `result` includes the `needs_operator` verdict.
- `result` includes the `page_operator` remediation action.

## First Failure Checks

- If `workflow list` returns no executions, rerun the starter and confirm the
  query uses the `operator-runbook-checkout-` workflow ID prefix.
- If `workflow describe` cannot find the ID, copy the full workflow ID from the
  latest `workflow list` output and keep the `default` namespace.
- If the execution stays running or queued, confirm the worker is still running
  with the same `TEMPORAL_TASK_QUEUE` value as the starter.
- If the remediation result is healthy, confirm the starter command includes
  `-simulate-unhealthy`.

Use the [local Temporal troubleshooting guide](local-temporal-troubleshooting.md)
when the server, worker, starter, or verdict does not match the expected local
proof.
