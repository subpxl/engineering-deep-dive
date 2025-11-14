from temporalio import workflow, activity
import asyncio
from datetime import timedelta


# -------------------------
# Activities (real work)
# -------------------------
@activity.defn(name="charge_customer")
async def charge_customer(order_id: str) -> str:
    print(f"💳 Charging customer for order {order_id}")
    return "payment_confirmed"

@activity.defn(name="send_to_restaurant")
async def send_to_restaurant(order_id: str) -> str:
    print(f"🍴 Sending order {order_id} to restaurant")
    return "restaurant_confirmed"

@activity.defn(name="assign_driver")
async def assign_driver(order_id: str) -> str:
    print(f" Assigning driver for order {order_id}")
    await asyncio.sleep(1)  # simulate waiting
    return "driver_assigned"

@activity.defn(name="track_delivery")
async def track_delivery(order_id: str) -> str:
    print(f"📦 Tracking delivery for order {order_id}")
    return "delivered"

# -------------------------
# Workflow (the recipe)
# -------------------------
@workflow.defn
class FoodDeliveryWorkflow:
    @workflow.run
    async def run(self, order_id: str) -> str:
        payment = await workflow.execute_activity(
            charge_customer, order_id, start_to_close_timeout=timedelta(seconds=10)
        )
        restaurant = await workflow.execute_activity(
            send_to_restaurant, order_id, schedule_to_close_timeout=timedelta(seconds=10)
        )
        driver = await workflow.execute_activity(
            assign_driver, order_id, schedule_to_close_timeout=timedelta(seconds=10)
        )
        delivery = await workflow.execute_activity(
            track_delivery, order_id, schedule_to_close_timeout=timedelta(seconds=10)
        )

        return f"Order {order_id} -> {payment}, {restaurant}, {driver}, {delivery}"
