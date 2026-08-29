# Security Policy

## Supported Versions

This repository is pre-release. Security fixes are applied to the default
branch until a tagged release policy exists.

## Reporting A Vulnerability

Report suspected vulnerabilities through GitHub's private vulnerability
reporting flow when it is available for this repository. If private reporting is
not available, open a minimal public issue that says a private security report
is needed, without including exploit details, credentials, logs, customer data,
or incident names.

Please include:

- The affected command, workflow, activity, or documentation path.
- A short impact summary.
- Reproduction steps using synthetic targets and fake credentials.
- Any relevant version, commit, or environment details.

## Scope

In scope:

- Workflow or activity behavior that could expose secrets or private incident
  data.
- Unsafe handling of Temporal connection settings or task queues.
- Documentation that encourages insecure credential handling.

Out of scope:

- Findings that require access to private infrastructure not described in this
  repository.
- Reports containing real credentials, private incidents, or customer data.

The project aims to acknowledge security reports within five business days.
