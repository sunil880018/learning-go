package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type CreatePlanInput struct {
	Period   string            `json:"period" binding:"required"`
	Interval int               `json:"interval" binding:"required"`
	Item     PlanItem          `json:"item" binding:"required"`
	Notes    map[string]string `json:"notes"`
}

type PlanItem struct {
	Name        string `json:"name" binding:"required"`
	Amount      int64  `json:"amount" binding:"required"`
	Currency    string `json:"currency" binding:"required"`
	Description string `json:"description"`
}

func CreatePlan(c *gin.Context) {

	keyID := os.Getenv("RAZORPAY_ACCESS_KEY")
	keySecret := os.Getenv("RAZORPAY_SECRET_KEY")

	if keyID == "" || keySecret == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Razorpay credentials missing in env"})
		return
	}

	var input CreatePlanInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Marshal request body
	body, _ := json.Marshal(input)

	// Create request to Razorpay
	req, err := http.NewRequest("POST", "https://api.razorpay.com/v1/plans", bytes.NewBuffer(body))
	fmt.Println(req)
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

	// Return Razorpay response
	c.Data(resp.StatusCode, "application/json", respBody)
}

func GetAllPlans(c *gin.Context) {

	keyID := os.Getenv("RAZORPAY_ACCESS_KEY")
	keySecret := os.Getenv("RAZORPAY_SECRET_KEY")

	if keyID == "" || keySecret == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Razorpay credentials missing in env"})
		return
	}

	req, err := http.NewRequest("GET", "https://api.razorpay.com/v1/plans", nil)
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

	c.Data(resp.StatusCode, "application/json", respBody)
}

func GetPlanById(c *gin.Context) {

	keyID := os.Getenv("RAZORPAY_ACCESS_KEY")
	keySecret := os.Getenv("RAZORPAY_SECRET_KEY")

	if keyID == "" || keySecret == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Razorpay credentials missing in env"})
		return
	}

	planID := c.Param("id")
	if planID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Plan ID is required"})
		return
	}

	req, err := http.NewRequest("GET", "https://api.razorpay.com/v1/plans/"+planID, nil)
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

	c.Data(resp.StatusCode, "application/json", respBody)
}
