import asyncio
import signal
from datetime import datetime
from temporalio.client import Client
from temporalio.worker import Worker
from workflows import (
    FoodDeliveryWorkflow,
    charge_customer,
    send_to_restaurant,
    assign_driver,
    track_delivery,
)

TASK_QUEUE = "food-delivery-task-queue"
TEMPORAL_HOST = "localhost:7233"


async def run_worker(client: Client):
    """Start the Temporal worker to process workflows"""
    worker = Worker(
        client,
        task_queue=TASK_QUEUE,
        workflows=[FoodDeliveryWorkflow],
        activities=[
            charge_customer,
            send_to_restaurant,
            assign_driver,
            track_delivery,
        ],
    )
    print(f"🚀 Worker started. Listening on queue: {TASK_QUEUE}")
    await worker.run()

async def execute_workflow(client: Client,orderId: str):
    # Generate unique workflow ID with timestamp
    workflow_id = f"food-delivery-{datetime.now().strftime('%Y%m%d-%H%M%S')}"

    print(f"\n📤 Starting workflow: {workflow_id}")

    # Start workflow execution
    handle = await client.start_workflow(
        FoodDeliveryWorkflow.run,
        orderId,  
        id=workflow_id,
        task_queue=TASK_QUEUE,
    )

    print(f"✅ Workflow started | ID: {handle.id} | RunID: {handle.result_run_id}")

    # Wait for workflow result
    result = await handle.result()
    print(f"📦 Result: {result}\n")




async def main():
    client = await Client.connect(TEMPORAL_HOST)
    print(f"Connected to Temporal at {TEMPORAL_HOST}")

    # Start worker in background
    worker_task = asyncio.create_task(run_worker(client))

    # Wait for worker to initialize
    await asyncio.sleep(2)

    # Execute a test workflow
    orderId ="ORDER-123"
    await execute_workflow(client,orderId)

    # Keep worker running
    print("👀 Worker still running. Press Ctrl+C to exit...")

    # Handle graceful shutdown
    # def signal_handler():
    #     print("\n👋 Shutting down...")
    #     worker_task.cancel()

    # loop = asyncio.get_running_loop()
    # for sig in (signal.SIGINT, signal.SIGTERM):
    #     loop.add_signal_handler(sig, signal_handler)

    # try:
    #     await worker_task
    # except asyncio.CancelledError:
    #     pass

    try:
        await worker_task
    except KeyboardInterrupt:
        print("\n👋 Shutting down...")
        worker_task.cancel()
        try:
            await worker_task
        except asyncio.CancelledError:
            pass
if __name__ == "__main__":
    asyncio.run(main())
