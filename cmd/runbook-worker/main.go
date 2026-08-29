package main

import (
	"log/slog"
	"os"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"github.com/tempmail-new/operator-runbook-temporal-go/runbook"
)

const defaultTaskQueue = "operator-runbook"

func main() {
	if err := run(); err != nil {
		slog.Error("runbook worker failed", "err", err)
		os.Exit(1)
	}
}

func run() error {
	hostPort := envString("TEMPORAL_ADDRESS", client.DefaultHostPort)
	taskQueue := envString("TEMPORAL_TASK_QUEUE", defaultTaskQueue)

	c, err := client.Dial(client.Options{HostPort: hostPort})
	if err != nil {
		return err
	}
	defer c.Close()

	w := worker.New(c, taskQueue, worker.Options{})
	w.RegisterWorkflow(runbook.OperatorRunbookWorkflow)
	w.RegisterActivity(runbook.RunHealthCheck)
	w.RegisterActivity(runbook.RequestHumanRemediation)

	slog.Info("runbook worker started", "temporal_address", hostPort, "task_queue", taskQueue)
	return w.Run(worker.InterruptCh())
}

func envString(name string, fallback string) string {
	value := os.Getenv(name)
	if value == "" {
		return fallback
	}
	return value
}
