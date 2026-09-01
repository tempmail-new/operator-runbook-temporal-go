# operator-runbook-temporal-go

Durable operator runbook orchestration proof in Go and Temporal.

This repository demonstrates one narrow runbook workflow: check an operator
target, record the verdict, and request a human remediation step when the
synthetic health check reports an unhealthy target. The first slice is
intentionally small so the workflow boundary, local development path, and
deterministic tests stay easy to inspect.

## What is included

- One Temporal workflow in `runbook.OperatorRunbookWorkflow`.
- Two activity boundaries: `RunHealthCheck` and `RequestHumanRemediation`.
- Deterministic workflow tests that mock the activity boundary.
- A local worker command and a starter command for Temporal's development
  server.
- GitHub Actions validation for formatting, vetting, and tests.

## Quickstart

Run the validation suite:

```sh
go test ./...
```

Run the local Temporal proof:

```sh
temporal server start-dev
```

In another terminal:

```sh
go run ./cmd/runbook-worker
```

In a third terminal:

```sh
go run ./cmd/start-runbook -runbook-id demo -target checkout
```

The starter prints the workflow result as JSON. The default activities return a
healthy synthetic check.

Run the local remediation proof:

```sh
go run ./cmd/start-runbook -runbook-id demo-unhealthy -target checkout -simulate-unhealthy
```

The starter prints a `needs_operator` verdict with a synthetic remediation
request. See the remediation quickstart for the full local flow.

## Documentation

- [Architecture](docs/architecture.md)
- [Local Temporal quickstart](docs/local-temporal-quickstart.md)
- [Local Temporal troubleshooting](docs/local-temporal-troubleshooting.md)
- [Remediation path quickstart](docs/remediation-path-quickstart.md)
- [Release checklist](docs/release-checklist.md)
- [Workflow contract](docs/workflow-contract.md)

## Project

- [Contributing](CONTRIBUTING.md)
- [Bug reports](https://github.com/tempmail-new/operator-runbook-temporal-go/issues/new?template=bug.yml)
- [License](LICENSE)
- [Security](SECURITY.md)
- [Support](SUPPORT.md)
