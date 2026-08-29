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
