package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Config/env
var (
	RESEND_INTERVAL = 60 * time.Second
	MAX_RETRIES     = 5
	OTP_TTL         = 5 * time.Minute
)

// Globals
var (
	db       *gorm.DB
	mc       *memcache.Client
	rdb      *redis.Client
	redisCtx = context.Background()
)

func init() {
	var err error
	// DB
	db, err = gorm.Open(postgres.Open(os.Getenv("DATABASE_DSN")), &gorm.Config{})
	if err != nil {
		log.Fatalf("DB init: %v", err)
	}
	// Memcached
	mc = memcache.New(os.Getenv("MEMCACHED_ADDR"))
	// Redis
	rdb = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
	})
}

// Models
type Country struct{ Code string }
type Auth struct {
	ID        uint
	Phone     string
	IsDeleted bool
}
type User struct {
	ID        uint
	FirstName string
	LastName  string
}
type Metadata struct{ UserID uint }
type DeletionRequest struct{ UserID uint }

type OtpCache struct {
	Otp        string    `json:"otp"`
	SentAt     time.Time `json:"sent_at"`
	RetryCount int       `json:"retry_count"`
	LastSentAt time.Time `json:"last_sent_at"`
}

type SendOtpReq struct {
	CountryCode string `json:"country_code"`
	Phone       string `json:"phone"`
	RequestID   string `json:"request_id"`
}

type ApiResponse struct {
	Code    int    `json:"-"`
	Message string `json:"message"`
}

func SendPhoneOtpHandler(w http.ResponseWriter, r *http.Request) {
	var req SendOtpReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondJSON(w, ApiResponse{http.StatusBadRequest, "Invalid request payload"})
		return
	}
	log.Printf("[OTP] request=%s data=%+v", req.RequestID, req)
	err := SendPhoneOtp(r.Context(), req)
	if err != nil {
		if resp, ok := err.(ApiResponse); ok {
			respondJSON(w, resp)
			return
		}
		log.Printf("[OTP] error: %v", err)
		respondJSON(w, ApiResponse{http.StatusInternalServerError, "Internal Server Error"})
		return
	}
	respondJSON(w, ApiResponse{http.StatusOK, "OTP sent successfully"})
}

// If the DB call takes more than 2 seconds, the context cancels the operation if your are using time
// If ctx has a timeout or deadline (e.g., context.WithTimeout()), the query will be aborted automatically if it exceeds that time.
// This prevents slow DB calls from hanging indefinitely and consuming resources.

// helpful using context
// ✅ 1. Request Timeout / Cancellation
// ✅ 2. Graceful Shutdown
// ✅ 3. Per-request Logging / Tracing
// ✅ 4. Consistency Across Layers

// | Purpose             | Benefit                                      |
// | ------------------- | -------------------------------------------- |
// | Timeout             | Prevent long or stuck DB operations          |
// | Cancellation        | Clean shutdown of pending queries            |
// | Logging / Tracing   | Propagate request metadata across layers     |
// | Context propagation | Uniform control across HTTP, Redis, DB, etc. |

func SendPhoneOtp(ctx context.Context, req SendOtpReq) error {
	// 2. Validate country
	var country Country
	if err := db.WithContext(ctx).First(&country, "code = ?", req.CountryCode).Error; err != nil {
		return ApiResponse{http.StatusBadRequest, "Invalid country!"}
	}

	key := fmt.Sprintf("otp:%s", req.Phone)
	var cached OtpCache
	item, err := mc.Get(key)
	if err == nil {
		json.Unmarshal(item.Value, &cached)
	}

	// 4. Validate user
	var auth Auth
	if err := db.WithContext(ctx).
		Where("phone = ?", req.Phone).
		Where("is_deleted = false").
		First(&auth).Error; err != nil {
		return ApiResponse{http.StatusBadRequest, "User not found"}
	}
	var user User
	db.WithContext(ctx).First(&user, "id = ?", auth.ID)
	var md Metadata
	db.WithContext(ctx).First(&md, "user_id = ?", user.ID)
	var delReq DeletionRequest
	db.WithContext(ctx).First(&delReq, "user_id = ?", user.ID)

	// 5. Resend interval
	if !cached.LastSentAt.IsZero() && time.Since(cached.LastSentAt) < RESEND_INTERVAL {
		return ApiResponse{http.StatusOK, fmt.Sprintf("Please wait %d seconds before retrying", int(RESEND_INTERVAL.Seconds()))}
	}

	// 6. Retry limit
	if cached.RetryCount >= MAX_RETRIES {
		return ApiResponse{http.StatusBadRequest, "OTP_LIMIT_EXCEEDED: retry after cooldown"}
	}

	// 7/8: resend or new
	if !cached.LastSentAt.IsZero() {
		cached.RetryCount++
		cached.LastSentAt = time.Now()
	} else {
		cached = OtpCache{Otp: GenerateOTP(), SentAt: time.Now(), LastSentAt: time.Now(), RetryCount: 0}
	}
	// store cache
	b, _ := json.Marshal(cached)
	mc.Set(&memcache.Item{Key: key, Value: b, Expiration: int32(OTP_TTL.Seconds())})

	// publish to redis/sms
	pub := map[string]interface{}{
		"phone": req.Phone,
		"otp":   cached.Otp,
	}
	jm, _ := json.Marshal(pub)
	rdb.Publish(redisCtx, "sms", jm)

	log.Printf("[OTP] OTP delivered %s -> %s", req.Phone, cached.Otp)
	return nil
}

func respondJSON(w http.ResponseWriter, resp ApiResponse) {
	w.WriteHeader(resp.Code)
	json.NewEncoder(w).Encode(map[string]string{"message": resp.Message})
}

func GenerateOTP() string {
	return fmt.Sprintf("%06d", time.Now().UnixNano()%1000000)
}

func main() {
	http.HandleFunc("/send-otp", SendPhoneOtpHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
