package main

import (
	"async-queue/producer"
	"context"
	"log"
	"time"

	"github.com/hibiken/asynq"
	"go.uber.org/ratelimit"
)

// --- Middleware for task handlers ---
func rateLimitedHandler(rl ratelimit.Limiter, handler func(context.Context, *asynq.Task) error) func(context.Context, *asynq.Task) error {
	return func(ctx context.Context, t *asynq.Task) error {
		rl.Take() // Block until allowed by limiter
		return handler(ctx, t)
	}
}

func main() {
	// --- PRODUCER: enqueue a task ---
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: "127.0.0.1:6379"})
	defer client.Close()

	// Create a task with type "email:deliver" and some payload
	task, err := producer.NewEmailDeliveryTask(42, "some:template:id")
	if err != nil {
		log.Fatalf("could not create task: %v", err)
	}

	// Put this task into the "email_high_priority" queue
	info, err := client.Enqueue(task,
		asynq.Queue("email_high_priority"),
		asynq.MaxRetry(10),
		asynq.Retention(24*time.Hour)) // retry 3 times before moving to DLQ and keep the task for 24 hours

	if err != nil {
		log.Fatalf("could not enqueue task: %v", err)
	}
	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)

	// --- CONSUMER: process tasks ---
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: "127.0.0.1:6379"},
		asynq.Config{
			Concurrency: 2, // number of concurrent workers
			// Specify how many concurrent workers to use for each queue.
			// In this case, we have two queues "email_high_priority" and "email_bulk"
			// and each of them will be served by one worker.
			Queues: map[string]int{
				"email_high_priority": 1, // OTP, password reset (high importance)
				"email_bulk":          1, // newsletters, promotions
			},
		},
	)

	// Router for task handlers
	mux := asynq.NewServeMux()

	// Global limiter → e.g., max 50 tasks/sec (adjust per downstream SLA)
	globalLimiter := ratelimit.New(50)

	// “Whenever a task of type email:deliver arrives, call this handler.”

	mux.HandleFunc(
		producer.TypeEmailDelivery,
		rateLimitedHandler(globalLimiter, producer.HandleEmailDeliveryTask),
	)

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
