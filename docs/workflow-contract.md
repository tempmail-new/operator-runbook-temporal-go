# Workflow Contract

`runbook.OperatorRunbookWorkflow` is the only workflow in this repository. It
accepts one operator target, runs a health-check activity, and returns a compact
result that callers can store as runbook evidence.

## Input

The workflow accepts `runbook.RunbookInput`:

- `runbook_id`: required caller-owned identifier for the runbook execution.
- `target`: required operator target, such as a service, queue, or checkout
  path.
- `simulate_unhealthy`: optional synthetic test toggle. When true, the default
  health-check activity returns an unhealthy status so the local remediation path
  can be exercised without editing source.

The workflow validates both fields before scheduling activities. Invalid input
fails the workflow without running the health check.

## Result

The workflow returns `runbook.RunbookResult`:

- `runbook_id`: the original runbook identifier.
- `target`: the original operator target.
- `verdict`: the final workflow verdict.
- `checks`: the ordered health-check results observed during the run.
- `remediations`: omitted for healthy runs and populated when operator action is
  requested.

## Verdicts

- `healthy`: the health-check activity returned a healthy status, so no operator
  remediation was requested.
- `needs_operator`: the health-check activity returned a non-healthy status and
  the workflow recorded a human remediation request.

The contract intentionally exposes only these verdicts while the repository has
one workflow and synthetic activities. New verdicts should land with updated
tests and documentation.

## Activity Boundaries

- `RunHealthCheck` receives `HealthCheckRequest{Target, SimulateUnhealthy}` and
  returns `HealthCheckResult{Target, Status, Detail}`. The default
  implementation is synthetic so the local proof stays self-contained. With
  `SimulateUnhealthy` set, it returns `unhealthy` with the detail `synthetic
  health check failed by request`.
- `RequestHumanRemediation` receives `RemediationRequest{Target, Reason}` and
  returns `RemediationResult{Action, Detail}`. The default implementation records
  the operator action that would be handed to a human-owned incident or ticketing
  system in a later slice.

External calls belong behind these activities. Workflow code must stay
deterministic and should not perform direct I/O, wall-clock reads, network calls,
or process-level side effects.

## Retry And Error Expectations

Both activities use the same workflow activity options:

- `StartToCloseTimeout`: one minute.
- `InitialInterval`: one second.
- `BackoffCoefficient`: `1`.
- `MaximumAttempts`: `2`.

Input validation errors are returned directly before activity scheduling.
Health-check activity failures are wrapped as `run health check`. Remediation
activity failures are wrapped as `request remediation`. Callers should treat
returned workflow errors as failed runbook evidence rather than partial success.
