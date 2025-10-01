package main

import (
	"fmt"

	"razorpay-subscriptions/controllers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	r := gin.Default()
	_ = godotenv.Load()

	// use by admin
	r.POST("/create-plan", controllers.CreatePlan)
	r.GET("/all-plans", controllers.GetAllPlans)
	r.GET("/plan/:id", controllers.GetPlanById)

	// use by users
	r.POST("/create-subscription", controllers.CreateSubscription)
	r.GET("/all-subscriptions", controllers.GetAllSubscriptions)
	r.GET("/subscription/:id", controllers.GetSubscriptionById)
	r.POST("/subscription/:id/cancel", controllers.CancelSubscription)
	r.POST("/subscription/:id/pause", controllers.PauseSubscription)
	r.POST("/subscription/:id/resume", controllers.ResumeSubscription)

	// create customer
	r.POST("/create-customer", controllers.CreateCustomer)

	r.Run(":8080")
	fmt.Println("Server running on http://localhost:8080")

	// subscription webhooks

}
