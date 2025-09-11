// [CRON JOB] --> [Fetch 1K from DB] --loop-->
//                    [Worker Pool (e.g., 50 goroutines)]
//                            ↓
//                       [API Call]
//                            ↓
//                     [DB Status Update]

// | Metric                    | Value                    |
// | ------------------------- | ------------------------ |
// | Total loop iterations     | 1000 (for 1000 statuses) |
// | Total goroutines spawned  | 1000                     |
// | Max concurrent goroutines | **50 at a time**         |
// | Waits for completion?     | ✅ Yes, via `wg.Wait()`   |

// Why use this pattern?
// Safe on memory/CPU: launching 1000 goroutines is OK in Go, but not if they all make network calls at once — hence the semaphore.

// Scalable: Can easily tune concurrency (e.g. 50, 100) based on API limits or system load.

// Actual Execution Behavior:
// 1000 goroutines are created in total.

// But only 50 run concurrently at a time, due to:

// sem <- struct{}{} // Acquire slot
// ...
// <-sem             // Release slot

// This pattern blocks each goroutine until a "slot" is available in the
// semaphore (channel of size 50).

// -------------- Message Queue ------------------------
// ✅ Yes — using a message queue here is not only feasible but highly recommended for
// production-scale systems handling 1 lakh+ records and third-party API calls.

// ✅ When to Use a Message Queue (MQ)?
// You should use a message queue when:

// ✅ You have heavy workloads (like 100,000+ records).

// ✅ You want to decouple your data fetching from processing.

// ✅ You need retry, dead-letter queues, and reliable delivery.

// ✅ You need better observability, scalability, and failure recovery.

//            ┌────────────────────┐
//            │    Cron Job        │
//            │(every hour fetch)  │
//            └────────┬───────────┘
//                     │
//         [ Fetch 1lakh from DB ]
//                     │
//                     ▼
//             ┌─────────────────┐
//             │ Message Queue   │ ◄────── Dead Letter Queue (for retries)
//             │ (Kafka/NATS/etc)│
//             └────────┬────────┘
//                      │
//         ┌────────────┴────────────┐
//         │                         │
//         ▼                         ▼
//   Worker 1                   Worker N
// (call API + update DB)  (call API + update DB)

package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

const (
	batchSize   = 1000
	workerCount = 50
	maxRetries  = 3
	apiURL      = "https://third-party.api/endpoint"
)

type Status struct {
	ID     int
	UserID string
}

func fetchStatuses(db *sql.DB, offset int) ([]Status, error) {
	rows, err := db.Query(`SELECT id, user_id FROM statuses WHERE status = 'pending' LIMIT $1 OFFSET $2`, batchSize, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var statuses []Status
	for rows.Next() {
		var s Status
		if err := rows.Scan(&s.ID, &s.UserID); err != nil {
			return nil, err
		}
		statuses = append(statuses, s)
	}
	return statuses, nil
}

func callAPI(userID string) error {
	// Simulated API call
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s?user_id=%s", apiURL, userID), nil)
	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Handle response if needed
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API failed: %d", resp.StatusCode)
	}
	return nil
}

func updateStatus(db *sql.DB, id int, status string) {
	_, _ = db.Exec(`UPDATE statuses SET status = $1, updated_at = NOW() WHERE id = $2`, status, id)
}

func processRecord(db *sql.DB, s Status, wg *sync.WaitGroup, sem chan struct{}) {
	defer wg.Done()
	sem <- struct{}{} // acquire

	defer func() { <-sem }() // release

	var err error
	for attempt := 1; attempt <= maxRetries; attempt++ {
		err = callAPI(s.UserID)
		if err == nil {
			updateStatus(db, s.ID, "done")
			return
		}
		time.Sleep(2 * time.Second) // backoff
	}
	updateStatus(db, s.ID, "failed")
}

func main() {
	db, err := sql.Open("postgres", "your_postgres_dsn")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	offset := 0
	for {
		statuses, err := fetchStatuses(db, offset)
		if err != nil {
			panic(err)
		}
		if len(statuses) == 0 {
			break
		}

		var wg sync.WaitGroup
		sem := make(chan struct{}, workerCount)

		for _, s := range statuses {
			wg.Add(1)
			go processRecord(db, s, &wg, sem)
		}
		wg.Wait()

		offset += batchSize
	}
}
