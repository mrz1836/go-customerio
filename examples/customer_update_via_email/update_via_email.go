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

	// Update the customer
	err = client.UpdateCustomer("bob@example.com", map[string]interface{}{
		"created_at": time.Now().Unix(),
		"email":      "bob@example.com",
		"first_name": "Bob",
		"plan":       "basic",
	})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Customer Updated Successfully!")
}
