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

	// Updating the device
	err = client.UpdateDevice("123", &customerio.Device{
		ID:       "abcdefghijklmnopqrstuvwxyz",
		LastUsed: time.Now().UTC().Unix(),
		Platform: customerio.PlatformAndroid,
	})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Device Updated Successfully!")
}
