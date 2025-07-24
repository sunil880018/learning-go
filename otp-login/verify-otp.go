package otp

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type OTPVerifyRequest struct {
	PhoneNumber string `json:"phone_number"`
	OTP         string `json:"otp"`
	CountryCode string `json:"country_code"`
}

type OTPData struct {
	OTP           string `json:"otp"`
	InitialTime   int64  `json:"initial_time"`
	RetryAttempts int    `json:"retry_attempts"`
}

type User struct {
	ID        int
	Phone     string
	IsDeleted bool
}

type Auth struct {
	Phone     string
	DeletedAt *time.Time
}

type Config struct {
	OTP_TTL         int64
	MAX_OTP_RETRIES int
	MemcachedPrefix string
	RedisSMSTopic   string
}

type Service struct {
	DB     *gorm.DB
	Cache  *memcache.Client
	Redis  *redis.Client
	Config Config
	Logger *log.Logger
}

func (s *Service) VerifyOTP(ctx context.Context, req OTPVerifyRequest) (map[string]interface{}, error) {
	s.Logger.Printf("[RequestID: %v] Starting OTP verification for phone: %v", ctx.Value("request_id"), req.PhoneNumber)

	if req.PhoneNumber == "" || req.OTP == "" || req.CountryCode == "" {
		return nil, errors.New("Missing required fields")
	}

	key := fmt.Sprintf("%s%s", s.Config.MemcachedPrefix, req.PhoneNumber)
	item, err := s.Cache.Get(key)
	if err == memcache.ErrCacheMiss {
		return nil, errors.New("OTP session expired or not found")
	} else if err != nil {
		s.Logger.Printf("Memcache error: %v", err)
		return nil, errors.New("Internal Server Error")
	}

	var otpData OTPData
	err = json.Unmarshal(item.Value, &otpData)
	if err != nil {
		s.Logger.Printf("Failed to parse OTP cache: %v", err)
		return nil, errors.New("Internal Server Error")
	}

	if time.Now().Unix() > otpData.InitialTime+s.Config.OTP_TTL {
		return nil, errors.New("OTP expired")
	}

	if strings.TrimSpace(req.OTP) != otpData.OTP {
		otpData.RetryAttempts++
		if otpData.RetryAttempts > s.Config.MAX_OTP_RETRIES {
			return nil, errors.New("OTP retry limit exceeded")
		}
		updated, _ := json.Marshal(otpData)
		s.Cache.Set(&memcache.Item{Key: key, Value: updated})
		return nil, errors.New("Invalid OTP")
	}

	var auth Auth
	if err := s.DB.Where("phone = ? AND deleted_at IS NULL", req.PhoneNumber).First(&auth).Error; err != nil {
		return nil, errors.New("User not found")
	}

	var user User
	if err := s.DB.Where("phone = ?", req.PhoneNumber).First(&user).Error; err != nil {
		return nil, errors.New("User details not found")
	}
	if user.IsDeleted {
		return nil, errors.New("User is not active")
	}

	// Post-verification updates like login time or metadata can be done here
	s.Cache.Delete(key)

	s.Logger.Printf("OTP verified successfully for phone: %v", req.PhoneNumber)

	// Optionally notify via Redis
	msg := fmt.Sprintf("User %v verified OTP successfully", req.PhoneNumber)
	s.Redis.Publish(ctx, s.Config.RedisSMSTopic, msg)

	return map[string]interface{}{
		"status":  200,
		"message": "OTP verified successfully",
		"data": map[string]interface{}{
			"user_id":       user.ID,
			"session_token": "abcd1234efgh5678",
		},
	}, nil
}
