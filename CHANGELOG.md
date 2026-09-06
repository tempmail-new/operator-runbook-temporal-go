# Changelog

Release history for `operator-runbook-temporal-go`. Keep entries
operator-facing and update this file before tagging a release.

## Unreleased

- Add a docs-only Temporal Web UI quickstart for inspecting healthy and
  remediation executions in Temporal's native local UI.
- Add a docs-only operator journey index that routes healthy-run proof,
  remediation proof, local troubleshooting, workflow-contract review, and
  release hygiene from one README-linked entry point.
- Add GitHub issue-template routing that disables blank issues and points setup
  support plus security reports to the existing public guidance.
- Add a GitHub-native enhancement proposal issue form for future runbook,
  workflow, and integration ideas with README/support routing and deterministic
  docs regression coverage.

## 0.1.0 - Initial Public Baseline

- Add the Go and Temporal runbook workflow with deterministic tests over mocked
  activity boundaries.
- Provide local worker and starter commands for running the workflow against
  Temporal's development server.
- Document the architecture, local Temporal quickstart, workflow contract,
  remediation path quickstart, local troubleshooting path, and release checklist.
- Add contributor, security, support, bug-report, licensing, and GitHub
  repository metadata surfaces for public evaluation.
- Keep production deployment manifests, real pager or ticketing integrations,
  metrics integrations, and additional workflows outside the current release
  scope.
