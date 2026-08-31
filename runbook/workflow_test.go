package runbook_test

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/mock"
	"go.temporal.io/sdk/testsuite"

	"github.com/tempmail-new/operator-runbook-temporal-go/runbook"
)

func TestOperatorRunbookWorkflow_HealthyTarget(t *testing.T) {
	var suite testsuite.WorkflowTestSuite
	env := suite.NewTestWorkflowEnvironment()
	env.RegisterWorkflow(runbook.OperatorRunbookWorkflow)

	input := runbook.RunbookInput{
		RunbookID: "demo",
		Target:    "checkout",
	}
	check := runbook.HealthCheckResult{
		Target: "checkout",
		Status: runbook.HealthStatusHealthy,
		Detail: "synthetic health check passed",
	}

	env.OnActivity(runbook.RunHealthCheck, mock.Anything, runbook.HealthCheckRequest{
		Target: "checkout",
	}).Return(check, nil).Once()

	env.ExecuteWorkflow(runbook.OperatorRunbookWorkflow, input)

	assertWorkflowCompleted(t, env, input)

	var got runbook.RunbookResult
	if err := env.GetWorkflowResult(&got); err != nil {
		t.Fatalf("OperatorRunbookWorkflow(%+v) result error = %v, want nil", input, err)
	}

	want := runbook.RunbookResult{
		RunbookID: "demo",
		Target:    "checkout",
		Verdict:   runbook.VerdictHealthy,
		Checks:    []runbook.HealthCheckResult{check},
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("OperatorRunbookWorkflow(%+v) mismatch (-want +got):\n%s", input, diff)
	}

	env.AssertExpectations(t)
}

func TestOperatorRunbookWorkflow_UnhealthyTargetRequestsRemediation(t *testing.T) {
	var suite testsuite.WorkflowTestSuite
	env := suite.NewTestWorkflowEnvironment()
	env.RegisterWorkflow(runbook.OperatorRunbookWorkflow)

	input := runbook.RunbookInput{
		RunbookID: "demo",
		Target:    "checkout",
	}
	check := runbook.HealthCheckResult{
		Target: "checkout",
		Status: runbook.HealthStatusUnhealthy,
		Detail: "error budget burn is high",
	}
	remediation := runbook.RemediationResult{
		Action: "page_operator",
		Detail: "operator review requested",
	}

	env.OnActivity(runbook.RunHealthCheck, mock.Anything, runbook.HealthCheckRequest{
		Target: "checkout",
	}).Return(check, nil).Once()
	env.OnActivity(runbook.RequestHumanRemediation, mock.Anything, runbook.RemediationRequest{
		Target: "checkout",
		Reason: "error budget burn is high",
	}).Return(remediation, nil).Once()

	env.ExecuteWorkflow(runbook.OperatorRunbookWorkflow, input)

	assertWorkflowCompleted(t, env, input)

	var got runbook.RunbookResult
	if err := env.GetWorkflowResult(&got); err != nil {
		t.Fatalf("OperatorRunbookWorkflow(%+v) result error = %v, want nil", input, err)
	}

	want := runbook.RunbookResult{
		RunbookID:    "demo",
		Target:       "checkout",
		Verdict:      runbook.VerdictNeedsOperator,
		Checks:       []runbook.HealthCheckResult{check},
		Remediations: []runbook.RemediationResult{remediation},
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("OperatorRunbookWorkflow(%+v) mismatch (-want +got):\n%s", input, diff)
	}

	env.AssertExpectations(t)
}

func TestOperatorRunbookWorkflow_SimulatedUnhealthyTargetRequestsRemediation(t *testing.T) {
	var suite testsuite.WorkflowTestSuite
	env := suite.NewTestWorkflowEnvironment()
	env.RegisterWorkflow(runbook.OperatorRunbookWorkflow)
	env.RegisterActivity(runbook.RunHealthCheck)
	env.RegisterActivity(runbook.RequestHumanRemediation)

	input := runbook.RunbookInput{
		RunbookID:         "demo",
		Target:            "checkout",
		SimulateUnhealthy: true,
	}

	env.ExecuteWorkflow(runbook.OperatorRunbookWorkflow, input)

	assertWorkflowCompleted(t, env, input)

	var got runbook.RunbookResult
	if err := env.GetWorkflowResult(&got); err != nil {
		t.Fatalf("OperatorRunbookWorkflow(%+v) result error = %v, want nil", input, err)
	}

	want := runbook.RunbookResult{
		RunbookID: "demo",
		Target:    "checkout",
		Verdict:   runbook.VerdictNeedsOperator,
		Checks: []runbook.HealthCheckResult{
			{
				Target: "checkout",
				Status: runbook.HealthStatusUnhealthy,
				Detail: "synthetic health check failed by request",
			},
		},
		Remediations: []runbook.RemediationResult{
			{
				Action: "page_operator",
				Detail: "operator review requested for checkout: synthetic health check failed by request",
			},
		},
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("OperatorRunbookWorkflow(%+v) mismatch (-want +got):\n%s", input, diff)
	}
}

func TestOperatorRunbookWorkflow_ActivityFailureReturnsError(t *testing.T) {
	var suite testsuite.WorkflowTestSuite
	env := suite.NewTestWorkflowEnvironment()
	env.RegisterWorkflow(runbook.OperatorRunbookWorkflow)

	input := runbook.RunbookInput{
		RunbookID: "demo",
		Target:    "checkout",
	}
	env.OnActivity(runbook.RunHealthCheck, mock.Anything, runbook.HealthCheckRequest{
		Target: "checkout",
	}).Return(runbook.HealthCheckResult{}, errors.New("probe timeout")).Once()

	env.ExecuteWorkflow(runbook.OperatorRunbookWorkflow, input)

	if !env.IsWorkflowCompleted() {
		t.Fatalf("OperatorRunbookWorkflow(%+v) completed = false, want true", input)
	}
	if err := env.GetWorkflowError(); err == nil {
		t.Fatalf("OperatorRunbookWorkflow(%+v) error = nil, want non-nil", input)
	}

	env.AssertExpectations(t)
}

func assertWorkflowCompleted(t *testing.T, env *testsuite.TestWorkflowEnvironment, input runbook.RunbookInput) {
	t.Helper()

	if !env.IsWorkflowCompleted() {
		t.Fatalf("OperatorRunbookWorkflow(%+v) completed = false, want true", input)
	}
	if err := env.GetWorkflowError(); err != nil {
		t.Fatalf("OperatorRunbookWorkflow(%+v) error = %v, want nil", input, err)
	}
}
