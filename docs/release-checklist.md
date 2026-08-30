# Release Checklist

Use this checklist when publishing a tagged release from a validated `main`
branch. The repository is still a narrow Temporal workflow proof, so releases
should describe the operator value that exists now without implying production
deployment coverage or external integrations.

## Versioning

- Use semantic versions such as `v0.1.0`.
- Increment the minor version for new workflow, command, or documentation
  surfaces that change how operators adopt the repository.
- Increment the patch version for fixes that do not change the public workflow
  contract.
- Keep prerelease labels for experiments that are not ready to document as the
  supported path.

## Pre-Release Validation

Run the same local checks that contributors run before opening a pull request:

```sh
gofmt -w .
go vet ./...
go test ./...
go build ./...
```

Confirm `README.md`, `CONTRIBUTING.md`, `SECURITY.md`, `SUPPORT.md`, and
`LICENSE` still match the release scope. If a release changes the local Temporal
run path, rerun `docs/local-temporal-quickstart.md` manually and record the
result in the release notes.

## Release Notes

Each release note should include:

- The workflow, command, or documentation change that operators can use.
- Any validation evidence, including the local commands above and the hosted
  `validate` workflow result.
- Any known scope boundaries, especially the absence of production deployment
  manifests, real pager or ticketing integrations, and additional workflows.
- Upgrade notes when a command, Go version, Temporal dependency, or workflow
  input/output shape changes.

Do not describe production deployment manifests as shipped until they exist in
the repository.

## Tag Flow

1. Start from the latest `origin/main`.
2. Verify the public pull-request queue is clean or that the release only
   includes already-merged changes.
3. Run the pre-release validation commands.
4. Create an annotated tag, for example `git tag -a v0.1.0 -m "Release v0.1.0"`.
5. Push the tag after validation succeeds.
6. Publish release notes that link the tag, summarize operator-visible changes,
   and call out any manual local Temporal smoke result.
