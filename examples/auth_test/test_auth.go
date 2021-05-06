package main

import (
	"log"
	"os"

	"github.com/mrz1836/go-customerio"
)

func main() {

	// Load the client (with both Tracking)
	client, err := customerio.NewClient(
		customerio.WithTrackingKey(os.Getenv("TRACKING_SITE_ID"), os.Getenv("TRACKING_API_KEY")),
	)
	if err != nil {
		log.Fatalln(err)
	}

	// Test authentication
	err = client.TestAuth()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Authentication Successful!")
}
