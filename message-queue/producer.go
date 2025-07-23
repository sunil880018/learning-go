func publishStatusesToQueue(ch *amqp.Channel, statuses []Status) {
	for _, status := range statuses {
		body, _ := json.Marshal(status)
		err := ch.Publish(
			"",             // default exchange
			"status_queue", // routing key (queue name)
			false, false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        body,
			})
		if err != nil {
			log.Println("Failed to publish:", err)
		}
	}
}
