# Support

Use the smallest channel that fits the problem.

## Usage Questions

For setup or usage questions, open a GitHub issue with:

- The command you ran.
- Your Go version.
- Whether Temporal CLI is installed.
- The exact error output with secrets removed.

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

## Security Reports

Do not open public issues for vulnerabilities or secret exposure. Follow
`SECURITY.md` instead.

## Project Scope

This repository currently supports one local Go + Temporal runbook proof. It
does not yet provide production deployment manifests, pager integrations,
ticketing integrations, or multiple workflows.
