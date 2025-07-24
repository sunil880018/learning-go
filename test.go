package main

import (
	"fmt"
	"time"
)

func main() {

	x := fmt.Sprintf("%06d", time.Now().UnixNano()%1000000)
	fmt.Println(x)
	y := map[string]interface{}{
		"status":  200,
		"message": "OTP verified successfully",
		"data": map[string]interface{}{
			"user_id":       12,
			"session_token": "abcd1234efgh5678",
		},
	}
	fmt.Println(y)
}
