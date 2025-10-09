```sh
[CRON JOB] --> [Fetch 1K from DB] --loop-->
                   [Worker Pool (e.g., 50 goroutines)]
                           ↓
                      [API Call]
                           ↓
                    [DB Status Update]

| Metric                    | Value                    |
| ------------------------- | ------------------------ |
| Total loop iterations     | 1000 (for 1000 statuses) |
| Total goroutines spawned  | 1000                     |
| Max concurrent goroutines | **50 at a time**         |
| Waits for completion?     | ✅ Yes, via `wg.Wait()`   |

Why use this pattern?
Safe on memory/CPU: launching 1000 goroutines is OK in Go, but not if they all make network calls at once — hence the semaphore.

Scalable: Can easily tune concurrency (e.g. 50, 100) based on API limits or system load.

Actual Execution Behavior:
1000 goroutines are created in total.

But only 50 run concurrently at a time, due to:

sem <- struct{}{} // Acquire slot
...
<-sem             // Release slot

This pattern blocks each goroutine until a "slot" is available in the
semaphore (channel of size 50).

-------------- Message Queue ------------------------
✅ Yes — using a message queue here is not only feasible but highly recommended for
production-scale systems handling 1 lakh+ records and third-party API calls.

✅ When to Use a Message Queue (MQ)?
You should use a message queue when:

✅ You have heavy workloads (like 100,000+ records).

✅ You want to decouple your data fetching from processing.

✅ You need retry, dead-letter queues, and reliable delivery.

✅ You need better observability, scalability, and failure recovery.

           ┌────────────────────┐
           │    Cron Job        │
           │(every hour fetch)  │
           └────────┬───────────┘
                    │
        [ Fetch 1lakh from DB ]
                    │
                    ▼
            ┌─────────────────┐
            │ Message Queue   │ ◄────── Dead Letter Queue (for retries)
            │ (Kafka/NATS/etc)│
            └────────┬────────┘
                     │
        ┌────────────┴────────────┐
        │                         │
        ▼                         ▼
  Worker 1                   Worker N
(call API + update DB)  (call API + update DB)


```