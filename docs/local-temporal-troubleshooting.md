# Local Temporal Troubleshooting

Use this guide when the local Temporal proof does not start, the workflow stays
queued, or the starter does not return the expected healthy or remediation
verdict.

## Fast Path

1. Confirm the Temporal CLI is installed:

   ```sh
   temporal version
   ```

2. Start the local development server and leave it running:

   ```sh
   temporal server start-dev
   ```

3. Start the worker in a second terminal and wait for this log line:

   ```text
   runbook worker started
   ```

4. Start the workflow in a third terminal:

   ```sh
   go run ./cmd/start-runbook -runbook-id demo -target checkout
   ```

Use the same `TEMPORAL_ADDRESS` and `TEMPORAL_TASK_QUEUE` values for the worker
and starter. The defaults are `localhost:7233` and `operator-runbook`.

## Temporal Server Startup

If the starter prints `start runbook failed`, or if the worker cannot connect:

- Confirm `temporal server start-dev` is still running in the foreground.
- Confirm `TEMPORAL_ADDRESS` points at the dev server address. Unset it for the
  default local server:

  ```sh
  unset TEMPORAL_ADDRESS
  ```

- If another process is using the default port, either stop that process or
  restart the worker and starter with the address printed by the Temporal dev
  server.
- Run `go test ./...` when you only need to validate workflow logic. The tests
  use Temporal's in-memory test environment and do not require a server.

## Worker Registration

If the starter connects but the workflow stays queued or never prints JSON:

- Confirm the worker terminal is still running and shows `runbook worker
  started`.
- Confirm the worker and starter use the same task queue:

  ```sh
  export TEMPORAL_TASK_QUEUE=operator-runbook
  ```

- Restart the worker after changing `TEMPORAL_TASK_QUEUE`; task queue changes do
  not move already-running workers.
- Keep one worker running per local proof. Starting the starter before the worker
  is allowed, but the workflow will wait until a worker polls the matching task
  queue.

## Starter And Verdicts

If the starter exits before starting a workflow:

- Provide non-empty `-runbook-id` and `-target` values. The workflow validates
  both before scheduling activities.
- If you run the starter repeatedly in the same second for the same target, wait
  one second and rerun it. The local workflow ID includes the target and a UTC
  timestamp with second precision, such as `operator-runbook-checkout-...`.

If the healthy path does not return the expected verdict:

- Run the default starter command without `-simulate-unhealthy`.
- The JSON result should include `"verdict": "healthy"` and omit
  `remediations`.

If the remediation proof does not return the expected verdict:

- Include `-simulate-unhealthy` on the starter command:

  ```sh
  go run ./cmd/start-runbook -runbook-id demo-unhealthy -target checkout -simulate-unhealthy
  ```

- The JSON result should include `"verdict": "needs_operator"`, a check with
  `"status": "unhealthy"`, and a remediation action of `page_operator`.

The unhealthy branch is intentionally synthetic. It proves the local
orchestration and operator handoff path without adding real pager, ticketing,
metrics, deployment, or external service integrations.
