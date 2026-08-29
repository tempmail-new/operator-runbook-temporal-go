package runbook

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

const (
	HealthStatusHealthy   = "healthy"
	HealthStatusUnhealthy = "unhealthy"

	VerdictHealthy       = "healthy"
	VerdictNeedsOperator = "needs_operator"
)

type RunbookInput struct {
	RunbookID string `json:"runbook_id"`
	Target    string `json:"target"`
}

func (i RunbookInput) Validate() error {
	if i.RunbookID == "" {
		return errors.New("runbook id is required")
	}
	if i.Target == "" {
		return errors.New("target is required")
	}
	return nil
}

type HealthCheckRequest struct {
	Target string `json:"target"`
}

type HealthCheckResult struct {
	Target string `json:"target"`
	Status string `json:"status"`
	Detail string `json:"detail"`
}

type RemediationRequest struct {
	Target string `json:"target"`
	Reason string `json:"reason"`
}

type RemediationResult struct {
	Action string `json:"action"`
	Detail string `json:"detail"`
}

type RunbookResult struct {
	RunbookID    string              `json:"runbook_id"`
	Target       string              `json:"target"`
	Verdict      string              `json:"verdict"`
	Checks       []HealthCheckResult `json:"checks"`
	Remediations []RemediationResult `json:"remediations,omitempty"`
}

func OperatorRunbookWorkflow(ctx workflow.Context, input RunbookInput) (RunbookResult, error) {
	if err := input.Validate(); err != nil {
		return RunbookResult{}, err
	}

	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Second,
			BackoffCoefficient: 1,
			MaximumAttempts:    2,
		},
	})

	var check HealthCheckResult
	checkRequest := HealthCheckRequest{Target: input.Target}
	if err := workflow.ExecuteActivity(ctx, RunHealthCheck, checkRequest).Get(ctx, &check); err != nil {
		return RunbookResult{}, fmt.Errorf("run health check: %w", err)
	}

	result := RunbookResult{
		RunbookID: input.RunbookID,
		Target:    input.Target,
		Verdict:   VerdictHealthy,
		Checks:    []HealthCheckResult{check},
	}
	if check.Status == HealthStatusHealthy {
		return result, nil
	}

	var remediation RemediationResult
	remediationRequest := RemediationRequest{
		Target: input.Target,
		Reason: check.Detail,
	}
	if err := workflow.ExecuteActivity(ctx, RequestHumanRemediation, remediationRequest).Get(ctx, &remediation); err != nil {
		return RunbookResult{}, fmt.Errorf("request remediation: %w", err)
	}

	result.Verdict = VerdictNeedsOperator
	result.Remediations = []RemediationResult{remediation}
	return result, nil
}

func RunHealthCheck(ctx context.Context, request HealthCheckRequest) (HealthCheckResult, error) {
	if request.Target == "" {
		return HealthCheckResult{}, errors.New("target is required")
	}

	return HealthCheckResult{
		Target: request.Target,
		Status: HealthStatusHealthy,
		Detail: "synthetic health check passed",
	}, nil
}

func RequestHumanRemediation(ctx context.Context, request RemediationRequest) (RemediationResult, error) {
	if request.Target == "" {
		return RemediationResult{}, errors.New("target is required")
	}
	if request.Reason == "" {
		return RemediationResult{}, errors.New("reason is required")
	}

	return RemediationResult{
		Action: "page_operator",
		Detail: fmt.Sprintf("operator review requested for %s: %s", request.Target, request.Reason),
	}, nil
}
