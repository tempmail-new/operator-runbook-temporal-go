# Operator Journey Index

Use this index when you know the operator question you need to answer and want
the shortest proof path without reading every guide first.

## Prove The Healthy Run

Start with the [local Temporal quickstart](local-temporal-quickstart.md) when
you need to prove the default workflow path from a clean checkout. It covers the
Temporal development server, worker command, starter command, default
`TEMPORAL_ADDRESS`, default `TEMPORAL_TASK_QUEUE`, and the expected `healthy`
verdict.

Run the validation suite first when you only need code-level confidence:

```sh
go test ./...
```

## Prove The Remediation Path

Use the [remediation path quickstart](remediation-path-quickstart.md) when you
need evidence that the workflow can request operator action. It runs the same
worker and starter flow with `-simulate-unhealthy`, then checks for the
`needs_operator` verdict and `page_operator` remediation action.

This path is intentionally synthetic. It proves the local orchestration branch
without adding real pager, ticketing, metrics, deployment, or external service
integrations.

## Inspect Executions In Temporal Web UI

Use the [Temporal Web UI quickstart](temporal-web-ui-quickstart.md) when you
need native Temporal evidence for the healthy and remediation paths. It shows
how to open `http://localhost:8233`, choose the `default` namespace, find the
latest `operator-runbook-checkout-` execution, and inspect the workflow type,
task queue, event history, input, result, `healthy` verdict, `needs_operator`
verdict, and `page_operator` remediation action.

## Recover A Failed Local Proof

Use the [local Temporal troubleshooting guide](local-temporal-troubleshooting.md)
when Temporal startup, worker registration, starter execution, or verdict checks
do not match the expected proof.

Fastest checks:

- Confirm `temporal server start-dev` is still running.
- Confirm the worker prints `runbook worker started`.
- Confirm the worker and starter share the same `TEMPORAL_TASK_QUEUE`.
- Rerun the healthy starter without `-simulate-unhealthy`, or rerun the
  remediation starter with `-simulate-unhealthy`, depending on the proof path.

## Review The Workflow Contract

Read the [workflow contract](workflow-contract.md) when you need to understand
the supported input, result, verdict, activity boundary, retry, and error
expectations before changing code or wiring an external caller.

The current contract has one workflow, `runbook.OperatorRunbookWorkflow`, and two
verdicts: `healthy` and `needs_operator`.

## Prepare Release Hygiene

Use the [release checklist](release-checklist.md) when preparing a tagged
release from validated `main`. It keeps versioning, local validation, release
notes, hosted `validate` evidence, and annotated tag flow in one place.

Update `CHANGELOG.md` before tagging. Do not claim production deployment manifests
or real pager, ticketing, metrics, external service integrations, or additional
workflows until those surfaces exist in the repository.
