# Support

Use the smallest channel that fits the problem.

GitHub blank issues are disabled so setup questions, reproducible bugs,
enhancement ideas, and security reports start from the routed paths below.

## Usage Questions

For setup or usage questions, start here and gather:

- The command you ran.
- Your Go version.
- Whether Temporal CLI is installed.
- The exact error output with secrets removed.

If the question reveals a reproducible failure, open the bug report issue form
with those details.

## Bug Reports

For reproducible bugs, open the
[GitHub bug report issue form](https://github.com/tempmail-new/operator-runbook-temporal-go/issues/new?template=bug.yml).
It asks for:

- The repository commit or release.
- The affected area: Temporal server startup, worker registration, starter CLI,
  workflow verdict, remediation path, validation, or documentation.
- The Go version, Temporal CLI version, operating system, Temporal address, and
  task queue.
- The exact Temporal server, worker, starter, or validation command that
  reproduces the issue.
- The workflow input, using synthetic targets.
- The expected result.
- The actual result.
- A minimal public reproduction and relevant logs with secrets removed.

## Enhancement Proposals

For future runbook, workflow, or integration ideas, open the
[GitHub enhancement proposal issue form](https://github.com/tempmail-new/operator-runbook-temporal-go/issues/new?template=feature.yml).
It asks for:

- The operator pain, recovery gap, or adoption problem.
- The narrow enhancement area, such as a runbook step, workflow behavior,
  activity boundary, local guide, pager or ticketing integration, deployment or
  secret handling, metrics or observability integration, or documentation.
- The smallest proposed scope that would solve the problem.
- The deterministic test or docs proof that does not require real pager,
  ticketing, metrics, deployment, customer, or private incident systems.
- The command output, workflow result, or documentation path a reviewer should
  inspect.
- The nondeterminism, secret-handling, private-incident, external-dependency,
  or production-readiness risks.

## Security Reports

Do not open public issues for vulnerabilities or secret exposure. Follow
`SECURITY.md` instead.

## Project Scope

This repository currently supports one local Go + Temporal runbook proof. It
does not yet provide production deployment manifests, pager integrations,
ticketing integrations, or multiple workflows.
