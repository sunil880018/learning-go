package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type SubscriptionAddon struct {
	Item struct {
		Name     string `json:"name" binding:"required"`
		Amount   int64  `json:"amount" binding:"required"`
		Currency string `json:"currency" binding:"required"`
	} `json:"item" binding:"required"`
}
type CreateSubscriptionInput struct {
	PlanId         string              `json:"plan_id" binding:"required"`
	TotalCount     int                 `json:"total_count" binding:"required"`
	Quantity       int                 `json:"quantity" binding:"required"`
	StartAt        int64               `json:"start_at,omitempty"`
	ExpireBy       int64               `json:"expire_by,omitempty"`
	CustomerNotify int                 `json:"customer_notify,omitempty"`
	Addons         []SubscriptionAddon `json:"addons,omitempty"`
	Notes          map[string]string   `json:"notes,omitempty"`
}

func CreateSubscription(c *gin.Context) {
	keyID := os.Getenv("RAZORPAY_ACCESS_KEY")
	keySecret := os.Getenv("RAZORPAY_SECRET_KEY")

	if keyID == "" || keySecret == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Razorpay credentials missing in env"})
		return
	}
	var input CreateSubscriptionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	body, _ := json.Marshal(input)

	req, err := http.NewRequest("POST", "https://api.razorpay.com/v1/subscriptions", bytes.NewBuffer(body))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(keyID, keySecret)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		c.JSON(resp.StatusCode, gin.H{"error": string(respBody)})
		return
	}

	c.Data(resp.StatusCode, "application/json", respBody)
}

func GetAllSubscriptions(c *gin.Context) {
	keyID := os.Getenv("RAZORPAY_ACCESS_KEY")
	keySecret := os.Getenv("RAZORPAY_SECRET_KEY")

	if keyID == "" || keySecret == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Razorpay credentials missing in env"})
		return
	}

	req, err := http.NewRequest("GET", "https://api.razorpay.com/v1/subscriptions", nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	req.SetBasicAuth(keyID, keySecret)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	c.Data(resp.StatusCode, "application/json", respBody)
}

func GetSubscriptionById(c *gin.Context) {
	keyID := os.Getenv("RAZORPAY_ACCESS_KEY")
	keySecret := os.Getenv("RAZORPAY_SECRET_KEY")

	if keyID == "" || keySecret == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Razorpay credentials missing in env"})
		return
	}

	subscriptionID := c.Param("id")
	if subscriptionID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Subscription ID is required"})
		return
	}

	req, err := http.NewRequest("GET", "https://api.razorpay.com/v1/subscriptions/"+subscriptionID, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	req.SetBasicAuth(keyID, keySecret)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	c.Data(resp.StatusCode, "application/json", respBody)
}

func CancelSubscription(c *gin.Context) {
	// Implementation for canceling a subscription
	keyID := os.Getenv("RAZORPAY_ACCESS_KEY")
	keySecret := os.Getenv("RAZORPAY_SECRET_KEY")

	if keyID == "" || keySecret == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Razorpay credentials missing in env"})
		return
	}

	subscriptionID := c.Param("id")
	if subscriptionID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Subscription ID is required"})
		return
	}

	req, err := http.NewRequest("POST", "https://api.razorpay.com/v1/subscriptions/"+subscriptionID+"/cancel", nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	req.SetBasicAuth(keyID, keySecret)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	c.Data(resp.StatusCode, "application/json", respBody)
}
func PauseSubscription(c *gin.Context) {
	// Implementation for pausing a subscription
	keyID := os.Getenv("RAZORPAY_ACCESS_KEY")
	keySecret := os.Getenv("RAZORPAY_SECRET_KEY")

	if keyID == "" || keySecret == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Razorpay credentials missing in env"})
		return
	}

	subscriptionID := c.Param("id")
	if subscriptionID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Subscription ID is required"})
		return
	}

	req, err := http.NewRequest("POST", "https://api.razorpay.com/v1/subscriptions/"+subscriptionID+"/pause", nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	req.SetBasicAuth(keyID, keySecret)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	c.Data(resp.StatusCode, "application/json", respBody)
}

func ResumeSubscription(c *gin.Context) {
	// Implementation for resuming a subscription
	keyID := os.Getenv("RAZORPAY_ACCESS_KEY")
	keySecret := os.Getenv("RAZORPAY_SECRET_KEY")

	if keyID == "" || keySecret == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Razorpay credentials missing in env"})
		return
	}
	subscriptionID := c.Param("id")
	if subscriptionID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Subscription ID is required"})
		return
	}

	req, err := http.NewRequest("POST", "https://api.razorpay.com/v1/subscriptions/"+subscriptionID+"/resume", nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	req.SetBasicAuth(keyID, keySecret)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	c.Data(resp.StatusCode, "application/json", respBody)
}
