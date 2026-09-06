# Temporal Web UI Quickstart

This guide shows the shortest local path to inspect healthy and remediation
workflow executions in Temporal's Web UI.

## Prerequisites

- Go 1.22 or newer.
- Temporal CLI.
- A browser that can reach the local Temporal Web UI.

Confirm the Temporal CLI is available:

```sh
temporal version
```

## Start The Local Proof

Start Temporal and leave it running:

```sh
temporal server start-dev
```

The development server exposes Temporal's Web UI at
`http://localhost:8233` and the frontend API at `localhost:7233`.

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

Open `http://localhost:8233`, choose the `default` namespace, and open the
latest workflow ID that begins with `operator-runbook-checkout-`.

Check these UI signals:

- Workflow type is `OperatorRunbookWorkflow`.
- Task queue is `operator-runbook` unless `TEMPORAL_TASK_QUEUE` was overridden.
- Execution status is completed.
- Event history includes `RunHealthCheck`.
- The workflow input includes `"runbook_id": "demo"` and `"target": "checkout"`.
- The workflow result has the `healthy` verdict.

## Inspect A Remediation Execution

Keep the same Temporal server and worker running, then start the synthetic
unhealthy path:

```sh
go run ./cmd/start-runbook -runbook-id demo-unhealthy -target checkout -simulate-unhealthy
```

Open the latest `operator-runbook-checkout-` execution in the Web UI.

Check these UI signals:

- Workflow type is still `OperatorRunbookWorkflow`.
- Event history includes both `RunHealthCheck` and
  `RequestHumanRemediation`.
- The workflow input includes `"simulate_unhealthy": true`.
- The workflow result has the `needs_operator` verdict.
- The remediation result includes the `page_operator` remediation action.

## First Failure Checks

- If the Web UI does not load, confirm `temporal server start-dev` is still
  running and that no other process is using port `8233`.
- If no workflow appears, rerun the starter and open the latest
  `operator-runbook-checkout-` execution in the `default` namespace.
- If the workflow stays running or queued, confirm the worker is still running
  with the same `TEMPORAL_TASK_QUEUE` value as the starter.
- If the remediation execution is healthy, confirm the starter command includes
  `-simulate-unhealthy`.

Use the [local Temporal troubleshooting guide](local-temporal-troubleshooting.md)
when the server, worker, starter, or verdict does not match the expected local
proof.
