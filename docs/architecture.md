# Architecture

`operator-runbook-temporal-go` is a small Temporal application for durable
operator runbooks. The first slice proves the orchestration shape without adding
HTTP APIs, user interfaces, or additional workflows.

## Components

- `runbook.OperatorRunbookWorkflow` is the only workflow. It validates the
  requested runbook, schedules a health-check activity, and returns either a
  healthy verdict or a remediation-request verdict. The public input, result,
  verdict, activity, and retry/error expectations are documented in the
  [workflow contract](workflow-contract.md).
- `runbook.RunHealthCheck` is the first activity boundary. The default
  implementation returns a synthetic healthy result so local development is
  self-contained.
- `runbook.RequestHumanRemediation` is the second activity boundary. The
  default implementation records the operator action that would be routed to a
  human-owned incident or ticketing system in a later slice.
- `cmd/runbook-worker` registers the workflow and activities with Temporal.
- `cmd/start-runbook` starts one workflow execution against a local Temporal
  server and prints the result as JSON.

## Determinism

Workflow code only uses deterministic Temporal workflow APIs. Time, I/O, logs,
and external system calls remain outside the workflow body. Activity behavior is
exercised through Temporal's test environment so workflow tests can mock the
boundary and verify both healthy and remediation paths without a live Temporal
cluster.

## Scope Boundaries

This bootstrap intentionally avoids:

- HTTP APIs.
- UI surfaces.
- More than one workflow.
- Real pager, ticketing, metrics, or deployment integrations.

Those integrations belong behind activities once the core runbook contract is
clear and the first workflow has merged.
