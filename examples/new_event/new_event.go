package main

import (
	"log"
	"os"
	"time"

	"github.com/mrz1836/go-customerio"
)

func main() {

	// Load the client (with Tracking API enabled)
	client, err := customerio.NewClient(
		customerio.WithTrackingKey(os.Getenv("TRACKING_SITE_ID"), os.Getenv("TRACKING_API_KEY")),
	)
	if err != nil {
		log.Fatalln(err)
	}

	// New event
	err = client.NewEvent(
		"123", "order_completed", time.Now().UTC(),
		map[string]interface{}{
			"order_id": "1234567",
			"amount":   "99.99",
		})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Event Created Successfully!")
}
