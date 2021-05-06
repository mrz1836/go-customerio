package main

import (
	"log"
	"os"

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

	// Updating the device
	err = client.DeleteDevice("123", "abcdefghijklmnopqrstuvwxyz")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Device Deleted Successfully!")
}
