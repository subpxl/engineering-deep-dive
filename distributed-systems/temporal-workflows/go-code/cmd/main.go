package main

import (
	"context"
	"goproject/app/workflows"
	"log"
	"os"
	"os/signal"
	"syscall"

	"time"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

const taskQueue = "food-delivery-task-queue"

func main() {
	// Connect to Temporal
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("❌ Unable to create Temporal client:", err)
	}
	defer c.Close()

	// Start worker in background
	w := worker.New(c, taskQueue, worker.Options{})
	w.RegisterWorkflow(workflows.FoodDeliveryWorkflow)
	w.RegisterActivity(workflows.ChargeCustomer)
	w.RegisterActivity(workflows.SendToRestaurant)
	w.RegisterActivity(workflows.AssignDriver)
	w.RegisterActivity(workflows.TrackDelivery)

	go func() {
		log.Println("🚀 Worker started. Listening on queue:", taskQueue)
		if err := w.Run(worker.InterruptCh()); err != nil {
			log.Fatalln("❌ Worker failed:", err)
		}
	}()

	// Wait for worker to start
	time.Sleep(1 * time.Second)

	// Execute workflow as client
	workflowID := "food-delivery-" + time.Now().Format("20060102-150405")
	workflowOptions := client.StartWorkflowOptions{
		ID:        workflowID,
		TaskQueue: taskQueue,
	}

	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, workflows.FoodDeliveryWorkflow, "ORDER-123")
	if err != nil {
		log.Fatalln("❌ Unable to execute workflow:", err)
	}

	log.Printf("✅ Workflow started | ID: %s | RunID: %s\n", we.GetID(), we.GetRunID())

	// Wait for result
	var result string
	if err := we.Get(context.Background(), &result); err != nil {
		log.Fatalln("❌ Workflow failed:", err)
	}

	log.Println("📦 Result:", result)

	// Keep worker running until Ctrl+C
	log.Println("👀 Worker still running. Press Ctrl+C to exit...")
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	log.Println("👋 Shutting down...")
}
