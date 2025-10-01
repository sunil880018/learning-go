package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// subscription webhooks
	webhookRoute := gin.Default()
	webhookRoute.POST("/webhook/subscription/authenticated", SubscriptionAuthenticatedWebhook)
	webhookRoute.POST("/webhook/subscription/activated", SubscriptionActivatedWebhook)

	// Run webhook server
	if err := webhookRoute.Run(":8081"); err != nil {
		log.Fatalf("Webhook server failed: %v", err)
	}
	fmt.Println("Webhook server running on http://localhost:8081")
}
