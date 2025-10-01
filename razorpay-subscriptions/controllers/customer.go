package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type CreateCustomerRequest struct {
	Name         string            `json:"name" binding:"required"`
	Email        string            `json:"email" binding:"required,email"`
	Contact      string            `json:"contact" binding:"required"`
	FailExisting string            `json:"fail_existing,omitempty"`
	Gstin        string            `json:"gstin,omitempty"`
	Notes        map[string]string `json:"notes,omitempty"`
}

// CreateCustomer creates a new customer in Razorpay
func CreateCustomer(c *gin.Context) {
	var customerData CreateCustomerRequest
	if err := c.ShouldBindJSON(&customerData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	keyID := os.Getenv("RAZORPAY_ACCESS_KEY")
	secretKey := os.Getenv("RAZORPAY_SECRET_KEY")
	if keyID == "" || secretKey == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Razorpay keys missing in env"})
		return
	}

	data := map[string]interface{}{
		"name":          customerData.Name,
		"email":         customerData.Email,
		"contact":       customerData.Contact,
		"fail_existing": customerData.FailExisting,
	}
	if customerData.Gstin != "" {
		data["gstin"] = customerData.Gstin
	}
	if customerData.Notes != nil {
		data["notes"] = customerData.Notes
	}

	// Marshal data to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	req, err := http.NewRequest("POST", "https://api.razorpay.com/v1/customers", bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(keyID, secretKey)

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
