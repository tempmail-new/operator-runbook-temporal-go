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
				"[Contributing](CONTRIBUTING.md)",
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
				"production deployment manifests",
				"annotated tag",
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
			name: "local quickstart routes remediation proof",
			path: "docs/local-temporal-quickstart.md",
			contains: []string{
				"[remediation path quickstart](remediation-path-quickstart.md)",
				"`-simulate-unhealthy`",
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
				"Usage Questions",
				"Bug Reports",
				"Security Reports",
				"does not yet provide production deployment manifests",
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
