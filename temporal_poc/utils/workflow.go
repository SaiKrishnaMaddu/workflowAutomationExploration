package utils

import (
	"go.temporal.io/sdk/workflow"
)

// SampleWorkflow is a simple workflow function.
func SampleWorkflow(ctx workflow.Context, name string) (string, error) {
	logger := workflow.GetLogger(ctx)
	logger.Info("Workflow started", "name", name)

	var result string
	err := workflow.ExecuteActivity(ctx, SampleActivity, name).Get(ctx, &result)
	if err != nil {
		return "", err
	}

	logger.Info("Workflow completed", "result", result)
	return result, nil
}
