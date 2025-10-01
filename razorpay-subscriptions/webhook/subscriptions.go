package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// SubscriptionAuthenticatedPayload models the webhook payload
type SubscriptionAuthenticatedPayload struct {
	Entity    string `json:"entity"`
	AccountID string `json:"account_id"`
	Event     string `json:"event"`
	Payload   struct {
		Subscription struct {
			Entity struct {
				ID                  string `json:"id"`
				PlanID              string `json:"plan_id"`
				CustomerID          string `json:"customer_id"`
				Status              string `json:"status"`
				ChargeAt            int64  `json:"charge_at"`
				StartAt             int64  `json:"start_at"`
				EndAt               int64  `json:"end_at"`
				OfferID             string `json:"offer_id"`
				RemainingCount      int    `json:"remaining_count"`
				HasScheduledChanges bool   `json:"has_scheduled_changes"`
				// Add more fields as needed
			} `json:"entity"`
		} `json:"subscription"`
	} `json:"payload"`
	CreatedAt int64 `json:"created_at"`
}

type SubscriptionActivatedPayload struct {
	Entity    string `json:"entity"`
	AccountID string `json:"account_id"`
	Event     string `json:"event"`
	Payload   struct {
		Subscription struct {
			Entity struct {
				ID                  string `json:"id"`
				PlanID              string `json:"plan_id"`
				CustomerID          string `json:"customer_id"`
				Status              string `json:"status"`
				ChargeAt            int64  `json:"charge_at"`
				StartAt             int64  `json:"start_at"`
				EndAt               int64  `json:"end_at"`
				OfferID             string `json:"offer_id"`
				RemainingCount      int    `json:"remaining_count"`
				HasScheduledChanges bool   `json:"has_scheduled_changes"`
				// Add more fields as needed
			} `json:"entity"`
		} `json:"subscription"`
	} `json:"payload"`
	CreatedAt int64 `json:"created_at"`
}

// ValidateRazorpaySignature validates the webhook signature
func ValidateRazorpaySignature(body []byte, signature string, secret string) bool {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(body)
	expected := hex.EncodeToString(mac.Sum(nil))
	return expected == signature
}

// SubscriptionAuthenticatedWebhook handles the subscription.authenticated webhook
func SubscriptionAuthenticatedWebhook(c *gin.Context) {
	// Read raw body
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not read body"})
		return
	}

	// Get webhook signature from header
	signature := c.GetHeader("X-Razorpay-Signature")
	secret := os.Getenv("RAZORPAY_WEBHOOK_SECRET")
	if secret == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Webhook secret missing in env"})
		return
	}

	if !ValidateRazorpaySignature(body, signature, secret) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid signature"})
		return
	}

	// Parse payload
	// 	{
	//   "entity": "event",
	//   "account_id": "acc_F5Motm2sJ5Fomd",
	//   "event": "subscription.authenticated",
	//   "contains": [
	//     "subscription"
	//   ],
	//   "payload": {
	//     "subscription": {
	//       "entity": {
	//         "id": "sub_F5aa7VaVXtXh80",
	//         "entity": "subscription",
	//         "plan_id": "plan_F5Zu0nrXVhHV2m",
	//         "customer_id": "cust_F5ZuzTm0cqYpzp",
	//         "status": "authenticated",
	//         "current_start": null,
	//         "current_end": null,
	//         "ended_at": null,
	//         "quantity": 1,
	//         "notes": [],
	//         "charge_at": 1593109800,
	//         "start_at": 1593109800,
	//         "end_at": 1598380200,
	//         "auth_attempts": 0,
	//         "total_count": 3,
	//         "paid_count": 0,
	//         "customer_notify": true,
	//         "created_at": 1592811228,
	//         "expire_by": null,
	//         "short_url": null,
	//         "has_scheduled_changes": false,
	//         "change_scheduled_at": null,
	//         "source": "api",
	//         "offer_id":"offer_JHD834hjbxzhd38d",
	//         "remaining_count": 3
	//       }
	//     }
	//   },
	//   "created_at": 1592811255
	// }
	var payload SubscriptionAuthenticatedPayload
	if err := json.Unmarshal(body, &payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	// Only handle subscription.authenticated event
	if payload.Event != "subscription.authenticated" {
		c.JSON(http.StatusOK, gin.H{"message": "event ignored"})
		return
	}

	// update the subscription id to customer id

	// TODO: Implement business logic
	// e.g., update subscription status in DB, notify user, trigger email, etc.
	// Example:
	// fmt.Printf("Subscription %s authenticated for customer %s\n", sub.ID, sub.CustomerID)

	c.JSON(http.StatusOK, gin.H{"status": "authenticated"})
}

func SubscriptionActivatedWebhook(c *gin.Context) {
	// Read raw body
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not read body"})
		return
	}

	// Get webhook signature from header
	signature := c.GetHeader("X-Razorpay-Signature")
	secret := os.Getenv("RAZORPAY_WEBHOOK_SECRET")
	if secret == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Webhook secret missing in env"})
		return
	}

	if !ValidateRazorpaySignature(body, signature, secret) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid signature"})
		return
	}

	// Parse payload
	var payload SubscriptionActivatedPayload
	if err := json.Unmarshal(body, &payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	// Only handle subscription.activated event
	if payload.Event != "subscription.activated" {
		c.JSON(http.StatusOK, gin.H{"message": "event ignored"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "activated"})
}
