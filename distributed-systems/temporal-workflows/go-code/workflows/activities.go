package workflows

import (
	"context"
	"time"

	"go.temporal.io/sdk/activity"
)

// -------------------------
// Activities
// -------------------------

func ChargeCustomer(ctx context.Context, orderID string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("💳 Charging customer", "OrderID", orderID)
	time.Sleep(500 * time.Millisecond)
	return "payment_confirmed", nil
}

func SendToRestaurant(ctx context.Context, orderID string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Sending order %s to restaurant\n", orderID)
	time.Sleep(5 * time.Second)
	return "restaurant_confirmed", nil
}

func AssignDriver(ctx context.Context, orderID string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Assigning driver for order %s\n", orderID)
	time.Sleep(1 * time.Second)
	return "driver_assigned", nil
}

func TrackDelivery(ctx context.Context, orderID string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Tracking delivery for order %s\n", orderID)
	time.Sleep(2 * time.Second)
	return "delivered", nil
}
