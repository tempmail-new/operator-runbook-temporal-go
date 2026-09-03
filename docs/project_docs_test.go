package docs_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRepositoryTrustSurfaces(t *testing.T) {
	root := filepath.Join("..")

	checks := []struct {
		name     string
		path     string
		contains []string
	}{
		{
			name: "README links trust surfaces",
			path: "README.md",
			contains: []string{
				"[Release checklist](docs/release-checklist.md)",
				"[Workflow contract](docs/workflow-contract.md)",
				"[Remediation path quickstart](docs/remediation-path-quickstart.md)",
				"[Local Temporal troubleshooting](docs/local-temporal-troubleshooting.md)",
				"[Changelog](CHANGELOG.md)",
				"[Contributing](CONTRIBUTING.md)",
				"[Bug reports](https://github.com/tempmail-new/operator-runbook-temporal-go/issues/new?template=bug.yml)",
				"[Enhancement proposals](https://github.com/tempmail-new/operator-runbook-temporal-go/issues/new?template=feature.yml)",
				"[License](LICENSE)",
				"[Security](SECURITY.md)",
				"[Support](SUPPORT.md)",
			},
		},
		{
			name: "LICENSE publishes reuse terms",
			path: "LICENSE",
			contains: []string{
				"MIT License",
				"operator-runbook-temporal-go contributors",
				"Permission is hereby granted, free of charge",
				"THE SOFTWARE IS PROVIDED \"AS IS\"",
			},
		},
		{
			name: "CONTRIBUTING keeps validation explicit",
			path: "CONTRIBUTING.md",
			contains: []string{
				"gofmt -w .",
				"go vet ./...",
				"go test ./...",
				"workflow code deterministic",
				"docs/release-checklist.md",
			},
		},
		{
			name: "release checklist defines publish hygiene",
			path: "docs/release-checklist.md",
			contains: []string{
				"semantic versions",
				"go build ./...",
				"hosted",
				"`CHANGELOG.md`",
				"production deployment manifests",
				"annotated tag",
			},
		},
		{
			name: "changelog summarizes release history",
			path: "CHANGELOG.md",
			contains: []string{
				"Release history for `operator-runbook-temporal-go`",
				"## Unreleased",
				"enhancement proposal",
				"## 0.1.0 - Initial Public Baseline",
				"Go and Temporal runbook workflow",
				"Temporal's development server",
				"workflow contract",
				"bug-report",
				"production deployment manifests",
			},
		},
		{
			name: "workflow contract defines runtime boundary",
			path: "docs/workflow-contract.md",
			contains: []string{
				"runbook.OperatorRunbookWorkflow",
				"runbook.RunbookInput",
				"runbook.RunbookResult",
				"`simulate_unhealthy`",
				"`healthy`",
				"`needs_operator`",
				"`RunHealthCheck`",
				"`RequestHumanRemediation`",
				"`MaximumAttempts`: `2`",
				"`run health check`",
				"`request remediation`",
			},
		},
		{
			name: "remediation quickstart proves unhealthy path",
			path: "docs/remediation-path-quickstart.md",
			contains: []string{
				"go run ./cmd/start-runbook -runbook-id demo-unhealthy -target checkout -simulate-unhealthy",
				"`simulate_unhealthy=true`",
				"`needs_operator`",
				"`TEMPORAL_TASK_QUEUE`",
				"synthetic health check failed by request",
				"real pager, ticketing, metrics, and deployment",
			},
		},
		{
			name: "local troubleshooting maps recovery checks",
			path: "docs/local-temporal-troubleshooting.md",
			contains: []string{
				"temporal server start-dev",
				"runbook worker started",
				"`TEMPORAL_ADDRESS`",
				"`TEMPORAL_TASK_QUEUE`",
				"`-simulate-unhealthy`",
				"\"needs_operator\"",
				"start runbook failed",
				"operator-runbook-checkout-",
				"without adding real pager, ticketing",
			},
		},
		{
			name: "local quickstart routes remediation proof",
			path: "docs/local-temporal-quickstart.md",
			contains: []string{
				"[remediation path quickstart](remediation-path-quickstart.md)",
				"[local Temporal troubleshooting guide](local-temporal-troubleshooting.md)",
				"`-simulate-unhealthy`",
			},
		},
		{
			name: "remediation quickstart routes troubleshooting",
			path: "docs/remediation-path-quickstart.md",
			contains: []string{
				"[local Temporal troubleshooting guide](local-temporal-troubleshooting.md)",
				"`needs_operator`",
				"`TEMPORAL_TASK_QUEUE`",
			},
		},
		{
			name: "architecture links workflow contract",
			path: "docs/architecture.md",
			contains: []string{
				"[workflow contract](workflow-contract.md)",
				"input, result",
				"retry/error expectations",
			},
		},
		{
			name: "SECURITY routes private reports",
			path: "SECURITY.md",
			contains: []string{
				"private vulnerability",
				"without including exploit details",
				"fake credentials",
			},
		},
		{
			name: "SUPPORT separates usage bugs and security",
			path: "SUPPORT.md",
			contains: []string{
				"GitHub blank issues are disabled",
				"Usage Questions",
				"If the question reveals a reproducible failure",
				"Bug Reports",
				"GitHub bug report issue form",
				"Temporal server startup, worker registration, starter CLI",
				"Temporal CLI version",
				"workflow input",
				"minimal public reproduction",
				"Enhancement Proposals",
				"GitHub enhancement proposal issue form",
				"runbook, workflow, or integration ideas",
				"real pager",
				"private incident systems",
				"production-readiness risks",
				"Security Reports",
				"does not yet provide production deployment manifests",
			},
		},
		{
			name: "GitHub bug form captures local runbook failures",
			path: ".github/ISSUE_TEMPLATE/bug.yml",
			contains: []string{
				"Report a reproducible operator-runbook-temporal-go bug.",
				"Temporal server startup",
				"Worker registration or task queue",
				"Starter CLI",
				"Workflow verdict",
				"Remediation path",
				"Repository commit or release",
				"Go version",
				"Temporal CLI version",
				"TEMPORAL_ADDRESS=localhost:7233",
				"TEMPORAL_TASK_QUEUE=operator-runbook",
				"go run ./cmd/runbook-worker",
				"go run ./cmd/start-runbook -runbook-id demo -target checkout",
				"simulate_unhealthy=true",
				"Minimal public reproduction",
				"Follow SECURITY.md instead",
			},
		},
		{
			name: "GitHub enhancement form captures scoped runbook ideas",
			path: ".github/ISSUE_TEMPLATE/feature.yml",
			contains: []string{
				"Propose a scoped runbook, workflow, or integration enhancement.",
				"future runbook, workflow, or integration ideas",
				"Operator problem",
				"New runbook step",
				"Existing workflow behavior",
				"Activity boundary",
				"Pager or ticketing integration",
				"Deployment or secret handling",
				"Metrics or observability integration",
				"Proposed scope",
				"Deterministic proof",
				"without real pager, ticketing, metrics, deployment",
				"Operator evidence",
				"Scope and safety risks",
				"private incident detail",
				"Documentation impact",
			},
		},
		{
			name: "GitHub issue chooser routes public and private requests",
			path: ".github/ISSUE_TEMPLATE/config.yml",
			contains: []string{
				"blank_issues_enabled: false",
				"Support and usage questions",
				"https://github.com/tempmail-new/operator-runbook-temporal-go/blob/main/SUPPORT.md",
				"route reproducible failures to the bug form",
				"Security reports",
				"https://github.com/tempmail-new/operator-runbook-temporal-go/blob/main/SECURITY.md",
				"private disclosure path",
			},
		},
	}

	for _, tt := range checks {
		t.Run(tt.name, func(t *testing.T) {
			body := mustReadProjectFile(t, root, tt.path)
			for _, want := range tt.contains {
				if !strings.Contains(body, want) {
					t.Errorf("project file %s contains %q = false, want true", tt.path, want)
				}
			}
		})
	}
}

func mustReadProjectFile(t *testing.T, root string, path string) string {
	t.Helper()

	body, err := os.ReadFile(filepath.Join(root, path))
	if err != nil {
		t.Fatalf("os.ReadFile(%q) error = %v, want nil", path, err)
	}
	return string(body)
}
