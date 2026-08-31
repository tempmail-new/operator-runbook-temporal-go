package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"

	"go.temporal.io/sdk/client"

	"github.com/tempmail-new/operator-runbook-temporal-go/runbook"
)

const defaultTaskQueue = "operator-runbook"

func main() {
	if err := run(context.Background()); err != nil {
		slog.Error("start runbook failed", "err", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	runbookID := flag.String("runbook-id", "demo", "logical runbook identifier")
	target := flag.String("target", "checkout", "operator target to check")
	simulateUnhealthy := flag.Bool("simulate-unhealthy", false, "make the synthetic health check return unhealthy")
	flag.Parse()

	input := runbook.RunbookInput{
		RunbookID:         *runbookID,
		Target:            *target,
		SimulateUnhealthy: *simulateUnhealthy,
	}
	if err := input.Validate(); err != nil {
		return err
	}

	hostPort := envString("TEMPORAL_ADDRESS", client.DefaultHostPort)
	taskQueue := envString("TEMPORAL_TASK_QUEUE", defaultTaskQueue)

	c, err := client.Dial(client.Options{HostPort: hostPort})
	if err != nil {
		return err
	}
	defer c.Close()

	workflowID := fmt.Sprintf("operator-runbook-%s-%s", sanitizeID(input.Target), time.Now().UTC().Format("20060102T150405Z"))
	workflowRun, err := c.ExecuteWorkflow(ctx, client.StartWorkflowOptions{
		ID:        workflowID,
		TaskQueue: taskQueue,
	}, runbook.OperatorRunbookWorkflow, input)
	if err != nil {
		return err
	}

	var result runbook.RunbookResult
	if err := workflowRun.Get(ctx, &result); err != nil {
		return err
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(result); err != nil {
		return err
	}
	return nil
}

func envString(name string, fallback string) string {
	value := os.Getenv(name)
	if value == "" {
		return fallback
	}
	return value
}

func sanitizeID(value string) string {
	value = strings.ToLower(value)
	value = strings.ReplaceAll(value, " ", "-")
	value = strings.ReplaceAll(value, "/", "-")
	return value
}
