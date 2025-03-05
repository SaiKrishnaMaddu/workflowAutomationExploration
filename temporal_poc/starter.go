package main

import (
	"context"
	"log"

	"go.temporal.io/sdk/client"

	"temporalPoc/utils"
)

func main() {
	// Connect to Temporal
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer c.Close()

	// Start a workflow
	workflowOptions := client.StartWorkflowOptions{
		ID:        "hello-workflow",
		TaskQueue: "temporal-task-queue",
	}

	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, utils.SampleWorkflow, "Temporal User")
	if err != nil {
		log.Fatalf("Failed to start workflow: %v", err)
	}

	log.Printf("Started workflow: WorkflowID: %s, RunID: %s", we.GetID(), we.GetRunID())

	// Get workflow result
	var result string
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalf("Failed to get workflow result: %v", err)
	}

	log.Printf("Workflow result: %s", result)
}
