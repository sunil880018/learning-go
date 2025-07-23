func consumeQueue(ch *amqp.Channel) {
	msgs, err := ch.Consume("status_queue", "", false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	for msg := range msgs {
		var status Status
		if err := json.Unmarshal(msg.Body, &status); err != nil {
			log.Println("Error decoding message:", err)
			msg.Nack(false, false)
			continue
		}

		// Call third-party API
		if processStatus(status) {
			msg.Ack(false)
		} else {
			msg.Nack(false, true) // retry
		}
	}
}
