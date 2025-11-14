package workflows

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/workflow"
)

func FoodDeliveryWorkflow(ctx workflow.Context, orderID string) (string, error) {
	logger := workflow.GetLogger(ctx)
	logger.Info("Starting food delivery workflow", "OrderID", orderID)

	options := workflow.ActivityOptions{
		ScheduleToCloseTimeout: time.Second * 30,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	var payment, restaurant, driver, delivery string

	// Run activities in order
	if err := workflow.ExecuteActivity(ctx, ChargeCustomer, orderID).Get(ctx, &payment); err != nil {
		return "", fmt.Errorf("payment failed: %w", err)
	}
	if err := workflow.ExecuteActivity(ctx, SendToRestaurant, orderID).Get(ctx, &restaurant); err != nil {
		return "", fmt.Errorf("restaurant failed: %w", err)
	}
	if err := workflow.ExecuteActivity(ctx, AssignDriver, orderID).Get(ctx, &driver); err != nil {
		return "", fmt.Errorf("driver assignment failed: %w", err)
	}
	if err := workflow.ExecuteActivity(ctx, TrackDelivery, orderID).Get(ctx, &delivery); err != nil {
		return "", fmt.Errorf("delivery tracking failed: %w", err)
	}

	return fmt.Sprintf("Order %s -> %s, %s, %s, %s", orderID, payment, restaurant, driver, delivery), nil
}
