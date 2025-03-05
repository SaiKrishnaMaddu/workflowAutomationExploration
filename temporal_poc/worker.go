package main

import (
	"fmt"
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"temporalPoc/utils"
)

func main() {
	fmt.Println("Temporal Poc")
	c, err := client.NewLazyClient(client.Options{})
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	defer c.Close()

	w := worker.New(c, "teporal-task-queue", worker.Options{})
	w.RegisterWorkflow(utils.SampleWorkflow)
	w.RegisterActivity(utils.SampleActivity)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalf("failed to start worker: %v", err)
	}

}
